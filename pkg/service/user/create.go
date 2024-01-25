package user

import (
	"api/ent/user"
	"api/lib/db"
	"api/lib/log"
	"api/pkg/result"

	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo"
	"github.com/shenghui0779/yiigo/hash"
	"go.uber.org/zap"
)

type ReqCreate struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

func Create(ctx context.Context, req *ReqCreate) result.Result {
	records, err := db.Client().User.Query().Unique(false).Select(user.FieldID).Where(user.Username(req.Username)).All(ctx)
	if err != nil {
		log.Error(ctx, "error query user", zap.Error(err))
		return result.ErrSystem(result.E(errors.WithMessage(err, "用户查询失败")))
	}
	if len(records) != 0 {
		return result.ErrPerm(result.M("该用户名已被使用"))
	}

	now := time.Now().Unix()
	salt := yiigo.Nonce(16)

	_, err = db.Client().User.Create().
		SetUsername(req.Username).
		SetPassword(hash.MD5(req.Password + salt)).
		SetSalt(salt).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		log.Error(ctx, "error create user", zap.Error(err))
		return result.ErrSystem(result.E(errors.WithMessage(err, "用户创建失败")))
	}

	return result.OK()
}
