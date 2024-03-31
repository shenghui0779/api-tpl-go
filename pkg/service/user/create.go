package user

import (
	"context"

	"api/ent/user"
	"api/lib/db"
	"api/lib/log"
	"api/pkg/result"

	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo"
	"github.com/shenghui0779/yiigo/hash"
	"go.uber.org/zap"
)

type ReqCreate struct {
	Phone    string `json:"phone" valid:"required"`
	Nickname string `json:"nickname" valid:"required"`
	Password string `json:"password" valid:"required"`
}

func Create(ctx context.Context, req *ReqCreate) result.Result {
	records, err := db.Client().User.Query().Unique(false).Select(user.FieldID).Where(user.Phone(req.Phone)).All(ctx)
	if err != nil {
		log.Error(ctx, "Error query user", zap.Error(err))
		return result.ErrSystem(result.E(errors.WithMessage(err, "用户查询失败")))
	}
	if len(records) != 0 {
		return result.ErrPerm(result.M("该手机号已被使用"))
	}

	salt := yiigo.Nonce(16)

	_, err = db.Client().User.Create().
		SetPhone(req.Phone).
		SetNickname(req.Nickname).
		SetPassword(hash.MD5(req.Password + salt)).
		SetSalt(salt).
		Save(ctx)
	if err != nil {
		log.Error(ctx, "Error create user", zap.Error(err))
		return result.ErrSystem(result.E(errors.WithMessage(err, "用户创建失败")))
	}

	return result.OK()
}
