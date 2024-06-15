package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"api/app/ent"
	"api/app/ent/user"
	"api/lib/identity"
	"api/lib/log"
	"api/lib/result"

	"github.com/shenghui0779/yiigo"
	"github.com/shenghui0779/yiigo/xhash"
	"go.uber.org/zap"
)

type ReqLogin struct {
	Account  string `json:"account" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type RespLogin struct {
	AuthToken string `json:"auth_token"`
}

// Login 登录
func Login(ctx context.Context, req *ReqLogin) result.Result {
	record, err := ent.DB.User.Query().Unique(false).Where(user.Account(req.Account)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return result.ErrAuth(result.M("用户不存在"))
		}
		log.Error(ctx, "Error query user", zap.Error(err))
		return result.ErrSystem(result.E(err))
	}

	if xhash.MD5(req.Password+record.Salt) != record.Password {
		return result.ErrAuth(result.M("密码错误"))
	}

	token := xhash.MD5(fmt.Sprintf("auth.%d.%d.%s", record.ID, time.Now().UnixMicro(), yiigo.Nonce(16)))

	authToken, err := identity.New(record.ID, token).AsAuthToken()
	if err != nil {
		log.Error(ctx, "Error auth_token", zap.Error(err))
		return result.ErrAuth(result.E(err))
	}

	_, err = ent.DB.User.Update().Where(user.ID(record.ID)).SetLoginAt(sql.NullTime{Time: time.Now(), Valid: true}).SetLoginToken(token).Save(ctx)
	if err != nil {
		log.Error(ctx, "Error update user", zap.Error(err))
		return result.ErrSystem(result.E(err))
	}
	return result.OK(result.V(&RespLogin{
		AuthToken: authToken,
	}))
}

// Logout 注销
func Logout(ctx context.Context) result.Result {
	id := identity.FromContext(ctx)
	if id.ID() == 0 {
		return result.OK()
	}

	_, err := ent.DB.User.Update().Where(user.ID(id.ID())).SetLoginToken("").Save(ctx)
	if err != nil {
		log.Error(ctx, "Error update user", zap.Error(err))
		return result.ErrSystem(result.E(err))
	}
	return result.OK()
}
