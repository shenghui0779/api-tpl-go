package user

import (
	"errors"
	"net/http"
	"time"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"api/ent"
	"api/ent/user"
	"api/lib"
	"api/logger"
	"api/pkg/result"
	"api/pkg/service/internal"
)

type ParamsCreate struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := new(ParamsCreate)
	if err := internal.BindJSON(r, params); err != nil {
		logger.Err(ctx, "err params", zap.Error(err))
		result.ErrParams().JSON(w, r)

		return
	}

	records, err := ent.DB.User.Query().Unique(false).Select(user.FieldID).Where(user.Username(params.Username)).All(ctx)
	if err != nil {
		logger.Err(ctx, "err query user", zap.Error(err))
		result.ErrParams().JSON(w, r)

		return
	}

	if len(records) != 0 {
		result.ErrParams(result.Err(errors.New("该用户名已被使用"))).JSON(w, r)

		return
	}

	now := time.Now().Unix()
	salt := lib.Nonce(16)

	_, err = ent.DB.User.Create().
		SetUsername(params.Username).
		SetPassword(yiigo.MD5(params.Password + salt)).
		SetSalt(salt).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		logger.Err(ctx, "err create user", zap.Error(err))
		result.ErrSystem().JSON(w, r)

		return
	}

	result.OK().JSON(w, r)
}
