package handlers

import (
	"net/http"
	"strconv"
	"tplgo/internal/result"
	"tplgo/pkg/service"

	"github.com/go-chi/chi/v5"
)

func UserInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		result.ErrParams.Wrap(result.WithMsg(err.Error())).JSON(w, r)

		return
	}

	service.NewUserService().Info(r.Context(), id).JSON(w, r)
}
