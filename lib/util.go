package lib

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"

	"api/lib/log"

	"github.com/go-chi/chi/v5"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

// Safe recover for goroutine when panic
func Safe(ctx context.Context, fn func(ctx context.Context)) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(ctx, "Goroutine Panic", zap.Any("error", err), zap.ByteString("stack", debug.Stack()))
		}
	}()

	fn(ctx)
}

func URLParamInt(r *http.Request, key string) int64 {
	param := chi.URLParam(r, key)
	v, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Error(r.Context(), "Error URLParamInt", zap.Error(err), zap.String("key", key), zap.String("value", param))
		return 0
	}
	return v
}

func URLQuery(query url.Values, key string) (string, bool) {
	if !query.Has(key) {
		return "", false
	}
	return query.Get(key), true
}

func URLQueryInt(ctx context.Context, query url.Values, key string) (int64, bool) {
	v, ok := URLQuery(query, key)
	if !ok || len(v) == 0 {
		return 0, false
	}

	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Error(ctx, "Error URLQueryInt", zap.Error(err), zap.String("key", key), zap.String("value", v))
		return 0, false
	}
	return i, true
}

func QueryPage(ctx context.Context, query url.Values) (offset, limit int) {
	limit = 20
	if v, ok := URLQueryInt(ctx, query, "size"); ok && v > 0 {
		limit = int(v)
	}
	if limit > 100 {
		limit = 100
	}

	if v, ok := URLQueryInt(ctx, query, "page"); ok && v > 0 {
		offset = (int(v) - 1) * limit
	}

	return
}

func PostPage(page, size int) (offset, limit int) {
	limit = 10
	if size > 0 {
		limit = size
	}
	if limit > 1000 {
		limit = 1000
	}
	if page > 0 {
		offset = (page - 1) * limit
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

func CheckFields(fields, columns []string) error {
	if len(fields) == 0 {
		return nil
	}
	for _, v := range fields {
		if !yiigo.SliceIn(columns, v) {
			return fmt.Errorf("invalid field: %s", v)
		}
	}
	return nil
}
