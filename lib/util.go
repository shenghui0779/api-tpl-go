package lib

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"io"
	"runtime/debug"

	"api/logger"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Recover recover panic for goroutine
func Recover(ctx context.Context) {
	if err := recover(); err != nil {
		logger.Err(ctx, "Goroutine Panic", zap.Any("error", err), zap.ByteString("stack", debug.Stack()))
	}
}

// CtxNewWithReqID returns a new context with request_id
func CtxNewWithReqID() context.Context {
	return context.WithValue(context.Background(), middleware.RequestIDKey, uuid.NewString())
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
