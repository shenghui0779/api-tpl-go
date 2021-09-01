package v1

import (
	"net/http"

	"tplgo/pkg/helpers"
	"tplgo/pkg/service"
)

func UserInfo(w http.ResponseWriter, r *http.Request) {
	id, er := helpers.URLParamInt(r, "id")

	if er != nil {
		er.JSON(w, r)

		return
	}

	service.NewUserService().Info(r.Context(), id).JSON(w, r)
}
