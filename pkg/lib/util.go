package lib

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"io"
	"net/http"
	"regexp"
	"runtime/debug"
	"strings"

	"tplgo/pkg/logger"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// Recover recover panic for goroutine
func Recover(ctx context.Context) {
	if err := recover(); err != nil {
		logger.Err(ctx, "Goroutine Panic", zap.Any("error", err), zap.ByteString("stack", debug.Stack()))
	}
}

// CtxCopyWithReqID returns a new context with request_id from origin context.
// Often used for goroutine.
func CtxCopyWithReqID(ctx context.Context) context.Context {
	return context.WithValue(context.Background(), middleware.RequestIDKey, middleware.GetReqID(ctx))
}

func Nonce(size uint8) string {
	nonce := make([]byte, size/2)
	io.ReadFull(rand.Reader, nonce)

	return hex.EncodeToString(nonce)
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
