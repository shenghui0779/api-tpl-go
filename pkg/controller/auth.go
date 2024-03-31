package controller

import (
	"api/lib/log"
	"api/pkg/internal"
	"api/pkg/result"
	"api/pkg/service"

	"net/http"

	"go.uber.org/zap"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := new(service.ReqLogin)
	if err := internal.BindJSON(r, req); err != nil {
		log.Error(ctx, "Error params", zap.Error(err))
		result.ErrParams(result.E(err)).JSON(w, r)

		return
	}

	service.Login(ctx, req).JSON(w, r)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	service.Logout(r.Context()).JSON(w, r)
}
