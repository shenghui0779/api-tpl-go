package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"tplgo/pkg/logger"

	"github.com/go-chi/chi/v5"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

var validator = yiigo.NewValidator()

func BindJSON(r *http.Request, obj interface{}) error {
	defer io.Copy(io.Discard, r.Body)

	if err := json.NewDecoder(r.Body).Decode(obj); err != nil {
		return err
	}

	return validator.ValidateStruct(obj)
}

func URLParamInt(r *http.Request, key string) int64 {
	v, err := strconv.ParseInt(chi.URLParam(r, key), 10, 64)

	if err != nil {
		logger.Err(r.Context(), "err url param to int64", zap.Error(err), zap.String("param", key))

		return 0
	}

	return v
}

func URLQueryInt(r *http.Request, key string) int64 {
	v, err := strconv.ParseInt(r.URL.Query().Get(key), 10, 64)

	if err != nil {
		logger.Err(r.Context(), "err url query to int64", zap.Error(err), zap.String("query", key))

		return 0
	}

	return v
}
