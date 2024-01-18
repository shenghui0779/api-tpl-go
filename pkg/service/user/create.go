package user

import (
	"api/ent/user"
	"api/lib/db"
	"api/lib/log"
	"api/pkg/internal"
	"api/pkg/result"

	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo/hash"
	yiigo_util "github.com/shenghui0779/yiigo/util"
	"go.uber.org/zap"
)

type ReqCreate struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := new(ReqCreate)
	if err := internal.BindJSON(r, params); err != nil {
		log.Error(ctx, "err params", zap.Error(err))
		result.ErrParams(result.E(errors.WithMessage(err, "参数错误"))).JSON(w, r)

		return
	}

	records, err := db.Client().User.Query().Unique(false).Select(user.FieldID).Where(user.Username(params.Username)).All(ctx)
	if err != nil {
		log.Error(ctx, "error query user", zap.Error(err))
		result.ErrSystem(result.E(errors.WithMessage(err, "用户查询失败"))).JSON(w, r)

		return
	}
	if len(records) != 0 {
		result.ErrPerm(result.M("该用户名已被使用")).JSON(w, r)
		return
	}

	now := time.Now().Unix()
	salt := yiigo_util.Nonce(16)

	_, err = db.Client().User.Create().
		SetUsername(params.Username).
		SetPassword(hash.MD5(params.Password + salt)).
		SetSalt(salt).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		log.Error(ctx, "error create user", zap.Error(err))
		result.ErrSystem(result.E(errors.WithMessage(err, "用户创建失败"))).JSON(w, r)

		return
	}

	result.OK().JSON(w, r)
}
