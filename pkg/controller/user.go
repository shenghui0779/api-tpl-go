package controller

import (
	"api/lib/log"
	"api/pkg/internal"
	"api/pkg/result"
	"api/pkg/service/user"

	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := new(user.ReqCreate)
	if err := internal.BindJSON(r, req); err != nil {
		log.Error(ctx, "err params", zap.Error(err))
		result.ErrParams(result.E(errors.WithMessage(err, "参数错误"))).JSON(w, r)

		return
	}

	user.Create(ctx, req).JSON(w, r)
}

func UserList(w http.ResponseWriter, r *http.Request) {
	user.List(r.Context(), r.URL.Query()).JSON(w, r)
}
