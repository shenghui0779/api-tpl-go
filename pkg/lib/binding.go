package lib

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"tplgo/pkg/consts"
	"tplgo/pkg/logger"

	"github.com/go-chi/chi/v5"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

var Validator = yiigo.NewValidator()

func BindJSON(r *http.Request, obj interface{}) error {
	if r.Body != nil && r.Body != http.NoBody {
		defer io.Copy(io.Discard, r.Body)

		if err := json.NewDecoder(r.Body).Decode(obj); err != nil {
			return err
		}
	}

	return Validator.ValidateStruct(obj)
}

// BindForm 解析Form表单并校验
func BindForm(r *http.Request, obj interface{}) error {
	switch yiigo.ContentType(r) {
	case consts.MIMEForm:
		if err := r.ParseForm(); err != nil {
			return err
		}
	case consts.MIMEMultipartForm:
		if err := r.ParseMultipartForm(consts.MaxFormMemory); err != nil {
			if err != http.ErrNotMultipart {
				return err
			}
		}
	}

	if err := yiigo.MapForm(obj, r.Form); err != nil {
		return err
	}

	return Validator.ValidateStruct(obj)
}

func URLParamInt(r *http.Request, key string) int64 {
	param := chi.URLParam(r, key)

	v, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		logger.Err(r.Context(), "err url param to int64", zap.Error(err), zap.String("key", key), zap.String("value", param))

		return 0
	}

	return v
}

func URLQuery(r *http.Request, key string) (string, bool) {
	query := r.URL.Query()

	if !query.Has(key) {
		return "", false
	}

	return query.Get(key), true
}

func URLQueryInt(r *http.Request, key string) (int64, bool) {
	query, ok := URLQuery(r, key)

	if !ok || len(query) == 0 {
		return 0, false
	}

	v, err := strconv.ParseInt(query, 10, 64)

	if err != nil {
		logger.Err(r.Context(), "err url query to int64", zap.Error(err), zap.String("key", key), zap.String("value", query))

		return 0, false
	}

	return v, true
}
