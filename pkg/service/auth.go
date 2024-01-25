package service

import (
	"api/ent"
	"api/ent/user"
	"api/lib/db"
	"api/lib/log"
	"api/pkg/auth"
	"api/pkg/result"

	"context"
	"fmt"
	"time"

	"github.com/shenghui0779/yiigo"
	"github.com/shenghui0779/yiigo/hash"
	"go.uber.org/zap"
)

type ReqLogin struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type RespLogin struct {
	AuthToken string `json:"auth_token"`
}

// Login 登录
func Login(ctx context.Context, req *ReqLogin) result.Result {
	record, err := db.Client().User.Query().Unique(false).Where(user.Username(req.Username)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return result.ErrAuth(result.M("用户不存在"))
		}

		log.Error(ctx, "error query user", zap.Error(err))
		return result.ErrSystem(result.E(err))
	}

	if hash.MD5(req.Password+record.Salt) != record.Password {
		return result.ErrAuth(result.M("密码错误"))
	}

	token := hash.MD5(fmt.Sprintf("auth.%d.%d.%s", record.ID, time.Now().UnixMicro(), yiigo.Nonce(16)))

	authToken, err := auth.NewIdentity(record.ID, token).AuthToken()
	if err != nil {
		log.Error(ctx, "error auth_token", zap.Error(err))
		return result.ErrAuth(result.E(err))
	}

	_, err = db.Client().User.Update().Where(user.ID(record.ID)).SetLoginAt(time.Now().Unix()).SetLoginToken(token).Save(ctx)
	if err != nil {
		log.Error(ctx, "error update user", zap.Error(err))
		return result.ErrSystem(result.E(err))
	}

	return result.OK(result.V(&RespLogin{
		AuthToken: authToken,
	}))
}

// Logout 注销
func Logout(ctx context.Context) result.Result {
	identity := auth.GetIdentity(ctx)
	if identity.ID() == 0 {
		return result.OK()
	}

	_, err := db.Client().User.Update().Where(user.ID(identity.ID())).
		SetLoginToken("").
		SetUpdatedAt(time.Now().Unix()).
		Save(ctx)
	if err != nil {
		log.Error(ctx, "error update user", zap.Error(err))
		return result.ErrSystem(result.E(err))
	}

	return result.OK()
}
