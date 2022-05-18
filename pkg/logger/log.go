package logger

import (
	"context"
	"fmt"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

func GetReqID(ctx context.Context) string {
	reqID := middleware.GetReqID(ctx)

	if len(reqID) == 0 {
		reqID = "-"
	}

	return reqID
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	yiigo.Logger().Info(fmt.Sprintf("[%s] %s", GetReqID(ctx), msg), fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	yiigo.Logger().Warn(fmt.Sprintf("[%s] %s", GetReqID(ctx), msg), fields...)
}

func Err(ctx context.Context, msg string, fields ...zap.Field) {
	yiigo.Logger().Error(fmt.Sprintf("[%s] %s", GetReqID(ctx), msg), fields...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	yiigo.Logger().Panic(fmt.Sprintf("[%s] %s", GetReqID(ctx), msg), fields...)
}
