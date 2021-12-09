package client

import (
	"context"
	"time"
	"tplgo/pkg/logger"

	"go.uber.org/zap"
)

type Logger interface {
	Log(ctx context.Context, data *LogData)
}

type LogData struct {
	URL        string        `json:"url"`
	Method     string        `json:"method"`
	Body       []byte        `json:"body"`
	StatusCode int           `json:"status_code"`
	Response   []byte        `json:"response"`
	Duration   time.Duration `json:"duration"`
	Error      error         `json:"error"`
}

type clientLogger struct{}

func (l *clientLogger) Log(ctx context.Context, data *LogData) {
	fields := make([]zap.Field, 0, 7)

	fields = append(fields,
		zap.String("method", data.Method),
		zap.String("url", data.URL),
		zap.ByteString("body", data.Body),
		zap.ByteString("response", data.Response),
		zap.Int("status", data.StatusCode),
		zap.String("duration", data.Duration.String()),
	)

	if data.Error != nil {
		fields = append(fields, zap.Error(data.Error))

		logger.Err(ctx, "[client] request error", fields...)

		return
	}

	logger.Info(ctx, "[client] request info", fields...)
}

// NewLogger returns default logger
func NewLogger() Logger {
	return new(clientLogger)
}
