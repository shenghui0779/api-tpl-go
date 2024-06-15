package user

import (
	"context"

	"api/app/ent"
	"api/lib/log"
	"api/lib/result"

	"github.com/shenghui0779/yiigo"
	"github.com/shenghui0779/yiigo/xhash"
	"go.uber.org/zap"
)

type ReqCreate struct {
	Account  string `json:"Account" valid:"required"`
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

func Create(ctx context.Context, req *ReqCreate) result.Result {
	salt := yiigo.Nonce(16)
	_, err := ent.DB.User.Create().
		SetAccount(req.Account).
		SetUsername(req.Username).
		SetPassword(xhash.MD5(req.Password + salt)).
		SetSalt(salt).
		Save(ctx)
	if err != nil {
		log.Error(ctx, "Error create user", zap.Error(err))
		if ent.IsConstraintError(err) {
			return result.ErrSystem(result.M("账号已存在"))
		}
		return result.ErrSystem(result.E(err))
	}
	return result.OK()
}
