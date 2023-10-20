package internal

import (
	"api/logger"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

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

func QueryPage(r *http.Request) (offset, limit int) {
	limit = 20

	if v, ok := URLQueryInt(r, "size"); ok && v > 0 {
		limit = int(v)
	}

	if limit > 100 {
		limit = 100
	}

	if v, ok := URLQueryInt(r, "page"); ok && v > 0 {
		offset = (int(v) - 1) * limit
	}

	return
}

func ExcelColumnIndex(name string) int {
	name = strings.ToUpper(name)

	if ok, err := regexp.MatchString(`^[A-Z]{1,2}$`, name); err != nil || !ok {
		return -1
	}

	index := 0

	for i, v := range name {
		if i != 0 {
			index = (index + 1) * 26
		}

		index += int(v - 'A')
	}

	return index
}
