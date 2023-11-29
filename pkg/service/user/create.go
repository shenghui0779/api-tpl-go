package user

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"api/db"
	"api/db/ent/user"
	"api/lib/hash"
	"api/lib/util"
	"api/logger"
	"api/pkg/result"
	"api/pkg/service/internal"
)

type ReqCreate struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := new(ReqCreate)
	if err := internal.BindJSON(r, params); err != nil {
		logger.Err(ctx, "err params", zap.Error(err))
		result.ErrParams(result.E(err)).JSON(w, r)

		return
	}

	records, err := db.Client().User.Query().Unique(false).Select(user.FieldID).Where(user.Username(params.Username)).All(ctx)
	if err != nil {
		logger.Err(ctx, "error query user", zap.Error(err))
		result.ErrSystem(result.E(err)).JSON(w, r)

		return
	}

	if len(records) != 0 {
		result.ErrPerm(result.M("该用户名已被使用")).JSON(w, r)
		return
	}

	now := time.Now().Unix()
	salt := util.Nonce(16)

	_, err = db.Client().User.Create().
		SetUsername(params.Username).
		SetPassword(hash.MD5(params.Password + salt)).
		SetSalt(salt).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		logger.Err(ctx, "error create user", zap.Error(err))
		result.ErrSystem(result.E(err)).JSON(w, r)

		return
	}

	result.OK().JSON(w, r)
}
