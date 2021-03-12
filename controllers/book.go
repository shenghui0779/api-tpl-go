package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/shenghui0779/demo/helpers"
	"github.com/shenghui0779/demo/service"
)

func BookInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		Err(w, r, helpers.Err(helpers.ErrParams, err.Error()))

		return
	}

	s := service.NewBookService()

	data, err := s.Info(r.Context(), id)

	if err != nil {
		Err(w, r, err)

		return
	}

	OK(w, data)
}
