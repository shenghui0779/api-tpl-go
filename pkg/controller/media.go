package controller

import (
	"net/http"

	"api/pkg/service/media"
)

func MediaList(w http.ResponseWriter, r *http.Request) {
	media.List(r.Context(), r.URL.Query()).JSON(w, r)
}
