package lib

import (
	"api/logger"

	"context"
	"runtime/debug"

	"go.uber.org/zap"
)

// Recover recover panic for goroutine
func Recover(ctx context.Context) {
	if err := recover(); err != nil {
		logger.Err(ctx, "Goroutine Panic", zap.Any("error", err), zap.ByteString("stack", debug.Stack()))
	}
}
