package logger

import (
	"api/lib/util"
	"context"

	"go.uber.org/zap"
)

var logger = debug()

func Init(cfg *Config) {
	logger = New(cfg)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Info(msg, append(fields, zap.String("req_id", util.GetReqID(ctx)))...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Warn(msg, append(fields, zap.String("req_id", util.GetReqID(ctx)))...)
}

func Err(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Error(msg, append(fields, zap.String("req_id", util.GetReqID(ctx)))...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Panic(msg, append(fields, zap.String("req_id", util.GetReqID(ctx)))...)
}
