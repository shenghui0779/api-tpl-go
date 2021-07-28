package helpers

import (
	"context"
	"runtime/debug"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

// Recover recover panic
func Recover(ctx context.Context) {
	if err := recover(); err != nil {
		yiigo.Logger().Error("Whoops! Server Panic",
			zap.String("request_id", middleware.GetReqID(ctx)),
			zap.Any("error", err),
			zap.ByteString("stack", debug.Stack()),
		)
	}
}
