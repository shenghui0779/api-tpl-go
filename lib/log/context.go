package log

import (
	"context"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

// GetReqID 获取请求ID
func GetReqID(ctx context.Context) string {
	reqID := middleware.GetReqID(ctx)
	if len(reqID) == 0 {
		reqID = "-"
	}

	return reqID
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
