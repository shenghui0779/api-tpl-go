package logger

import (
	"context"
	"fmt"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

var SQL = new(sqllogger)

type sqllogger struct{}

func (l *sqllogger) Info(ctx context.Context, query string, args ...interface{}) {
	yiigo.Logger().Info(fmt.Sprintf("[%s] [SQL] %s", middleware.GetReqID(ctx), query), zap.Any("args", args))
}

func (l *sqllogger) Error(ctx context.Context, err error) {
	yiigo.Logger().Error(fmt.Sprintf("[%s] SQL Builder Error", middleware.GetReqID(ctx)), zap.Error(err))
}
