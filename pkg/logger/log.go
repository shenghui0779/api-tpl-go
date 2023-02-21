package logger

import (
	"context"

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
	yiigo.Logger().Info(msg, append(fields, zap.String("req_id", GetReqID(ctx)))...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	yiigo.Logger().Warn(msg, append(fields, zap.String("req_id", GetReqID(ctx)))...)
}

func Err(ctx context.Context, msg string, fields ...zap.Field) {
	yiigo.Logger().Error(msg, append(fields, zap.String("req_id", GetReqID(ctx)))...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	yiigo.Logger().Panic(msg, append(fields, zap.String("req_id", GetReqID(ctx)))...)
}
