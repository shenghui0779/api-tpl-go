package log

import (
	"context"

	"go.uber.org/zap"
)

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	log.Info(msg, append(fields, zap.String("req_id", GetReqID(ctx)))...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	log.Warn(msg, append(fields, zap.String("req_id", GetReqID(ctx)))...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	log.Error(msg, append(fields, zap.String("req_id", GetReqID(ctx)))...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	log.Panic(msg, append(fields, zap.String("req_id", GetReqID(ctx)))...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	log.Fatal(msg, append(fields, zap.String("req_id", GetReqID(ctx)))...)
}
