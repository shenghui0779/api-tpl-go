package log

import (
	"api/lib/util"

	"context"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var logger = debug()

// Init 初始化日志实例(如有多个实例，在此方法中初始化)
func Init() {
	logger = New(buildCfg(viper.GetString("log.filename"), viper.GetStringMap("log.options")))
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Info(msg, append(fields, zap.String("req_id", util.GetReqID(ctx)))...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Warn(msg, append(fields, zap.String("req_id", util.GetReqID(ctx)))...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Error(msg, append(fields, zap.String("req_id", util.GetReqID(ctx)))...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Panic(msg, append(fields, zap.String("req_id", util.GetReqID(ctx)))...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Fatal(msg, append(fields, zap.String("req_id", util.GetReqID(ctx)))...)
}
