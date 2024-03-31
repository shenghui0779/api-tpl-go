package lib

import (
	"api/lib/log"

	"context"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// Recover recover panic for goroutine
func Recover(ctx context.Context) {
	if err := recover(); err != nil {
		log.Error(ctx, "Goroutine Panic", zap.Any("error", err), zap.ByteString("stack", debug.Stack()))
	}
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
