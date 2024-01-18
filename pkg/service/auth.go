package service

import (
	"api/ent"
	"api/ent/user"
	"api/lib/db"
	"api/lib/log"
	"api/pkg/auth"
	"api/pkg/internal"
	"api/pkg/result"

	"fmt"
	"net/http"
	"time"

	"github.com/shenghui0779/yiigo/hash"
	yiigo_util "github.com/shenghui0779/yiigo/util"
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
func Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := new(ReqLogin)
	if err := internal.BindJSON(r, params); err != nil {
		log.Error(ctx, "error params", zap.Error(err))
		result.ErrParams(result.E(err)).JSON(w, r)

		return
	}

	record, err := db.Client().User.Query().Unique(false).Where(user.Username(params.Username)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			result.ErrAuth(result.M("用户不存在")).JSON(w, r)
		} else {
			log.Error(ctx, "error query user", zap.Error(err))
			result.ErrSystem(result.E(err)).JSON(w, r)
		}

		return
	}

	if hash.MD5(params.Password+record.Salt) != record.Password {
		result.ErrAuth(result.M("密码错误")).JSON(w, r)
		return
	}

	token := hash.MD5(fmt.Sprintf("auth.%d.%d.%s", record.ID, time.Now().UnixMicro(), yiigo_util.Nonce(16)))

	authToken, err := auth.NewIdentity(record.ID, token).AuthToken()
	if err != nil {
		log.Error(ctx, "error auth_token", zap.Error(err))
		result.ErrAuth(result.E(err)).JSON(w, r)

		return
	}

	_, err = db.Client().User.Update().Where(user.ID(record.ID)).SetLoginAt(time.Now().Unix()).SetLoginToken(token).Save(ctx)
	if err != nil {
		log.Error(ctx, "error update user", zap.Error(err))
		result.ErrSystem(result.E(err)).JSON(w, r)

		return
	}

	resp := &RespLogin{
		AuthToken: authToken,
	}

	result.OK(result.V(resp)).JSON(w, r)
}

// Logout 注销
func Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	identity := auth.GetIdentity(ctx)
	if identity.ID() == 0 {
		result.OK().JSON(w, r)

		return
	}

	_, err := db.Client().User.Update().Where(user.ID(identity.ID())).
		SetLoginToken("").
		SetUpdatedAt(time.Now().Unix()).
		Save(ctx)
	if err != nil {
		log.Error(ctx, "error update user", zap.Error(err))
		result.ErrSystem(result.E(err)).JSON(w, r)

		return
	}

	result.OK().JSON(w, r)
}
