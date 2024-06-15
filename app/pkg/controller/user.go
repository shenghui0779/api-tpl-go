package controller

import (
	"net/http"

	"api/app/pkg/service/user"
	"api/lib"
	"api/lib/log"
	"api/lib/result"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := new(user.ReqCreate)
	if err := lib.BindJSON(r, req); err != nil {
		log.Error(ctx, "Error params", zap.Error(err))
		result.ErrParams(result.E(errors.WithMessage(err, "参数错误"))).JSON(w, r)
		return
	}
	user.Create(ctx, req).JSON(w, r)
}

func UserList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := new(user.ReqList)
	if err := lib.BindJSON(r, req); err != nil {
		log.Error(ctx, "Error params", zap.Error(err))
		result.ErrParams(result.E(errors.WithMessage(err, "参数错误"))).JSON(w, r)
		return
	}
	user.List(r.Context(), req).JSON(w, r)
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := new(user.ReqInfo)
	if err := lib.BindJSON(r, req); err != nil {
		log.Error(ctx, "Error params", zap.Error(err))
		result.ErrParams(result.E(errors.WithMessage(err, "参数错误"))).JSON(w, r)
		return
	}
	if len(req.IDs) == 0 {
		result.ErrParams(result.M("用户ID不可为空")).JSON(w, r)
		return
	}
	user.Info(r.Context(), req).JSON(w, r)
}
