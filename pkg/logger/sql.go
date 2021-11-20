package logger

import (
	"context"
	"fmt"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

var SQL = new(sqllogger)

type sqllogger struct{}

func (l *sqllogger) Info(ctx context.Context, query string, args ...interface{}) {
	yiigo.Logger().Info(fmt.Sprintf("[%s] [SQL] %s", GetReqID(ctx), query), zap.Any("args", args))
}

func (l *sqllogger) Error(ctx context.Context, err error) {
	yiigo.Logger().Error(fmt.Sprintf("[%s] SQL Builder Error", GetReqID(ctx)), zap.Error(err))
}
