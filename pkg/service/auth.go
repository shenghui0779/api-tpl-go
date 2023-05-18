package service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"tplgo/pkg/ent"
	"tplgo/pkg/ent/user"
	"tplgo/pkg/lib"
	"tplgo/pkg/logger"
	"tplgo/pkg/result"
)

// Auth 授权服务
type Auth interface {
	// Login 登录
	Login(w http.ResponseWriter, r *http.Request)
	// Logout 注销
	Logout(w http.ResponseWriter, r *http.Request)
}

func NewAuth() Auth {
	return new(auth)
}

type auth struct{}

type ParamsLogin struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type RespLogin struct {
	AuthToken string `json:"auth_token"`
}

func (a *auth) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := new(ParamsLogin)

	err := lib.BindJSON(r, params)

	if err != nil {
		logger.Err(ctx, "err params", zap.Error(err))
		result.ErrParams().JSON(w, r)

		return
	}

	record, err := ent.DB.User.Query().Unique(false).Where(user.Username(params.Username)).First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			result.ErrAuth(result.Err(errors.New("用户不存在"))).JSON(w, r)

			return
		}

		logger.Err(ctx, "err query user", zap.Error(err))
		result.ErrSystem(result.Err(errors.New("登录失败"))).JSON(w, r)

		return
	}

	if yiigo.MD5(params.Password+record.Salt) != record.Password {
		result.ErrAuth(result.Err(errors.New("密码错误"))).JSON(w, r)

		return
	}

	token := yiigo.MD5(fmt.Sprintf("auth.%d.%d.%s", record.ID, time.Now().UnixMicro(), lib.Nonce(16)))

	authToken, err := lib.NewIdentity(record.ID, token).AuthToken()

	if err != nil {
		logger.Err(ctx, "err auth_token", zap.Error(err))
		result.ErrAuth(result.Err(errors.New("登录失败"))).JSON(w, r)

		return
	}

	_, err = ent.DB.User.Update().Where(user.ID(record.ID)).
		SetLoginToken(token).
		SetLastLoginAt(time.Now().Unix()).
		Save(ctx)

	if err != nil {
		logger.Err(ctx, "err update user", zap.Error(err))
		result.ErrSystem(result.Err(errors.New("登录失败"))).JSON(w, r)

		return
	}

	resp := &RespLogin{
		AuthToken: authToken,
	}

	result.OK(result.Data(resp)).JSON(w, r)
}

func (a *auth) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	identity := lib.GetIdentity(ctx)

	if identity.ID() == 0 {
		result.OK().JSON(w, r)

		return
	}

	_, err := ent.DB.User.Update().Where(user.ID(identity.ID())).
		SetLoginToken("").
		SetUpdatedAt(time.Now().Unix()).
		Save(ctx)

	if err != nil {
		logger.Err(ctx, "err update user", zap.Error(err))
		result.ErrSystem().JSON(w, r)

		return
	}

	result.OK().JSON(w, r)
}
