package controller

import (
	"net/http"

	"api/app/pkg/service/media"
	"api/lib"
	"api/lib/log"
	"api/lib/result"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func MediaList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := new(media.ReqList)
	if err := lib.BindJSON(r, req); err != nil {
		log.Error(ctx, "Error params", zap.Error(err))
		result.ErrParams(result.E(errors.WithMessage(err, "参数错误"))).JSON(w, r)
		return
	}
	media.List(r.Context(), req).JSON(w, r)
}
