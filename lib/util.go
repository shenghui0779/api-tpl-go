package lib

import (
	"context"
	"runtime/debug"

	"api/lib/log"

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
