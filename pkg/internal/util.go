package internal

import (
	"context"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"api/lib/log"

	"go.uber.org/zap"
)

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
