package ent

import (
	"context"
	"database/sql"
	"fmt"
	"runtime/debug"

	cfg "tplgo/pkg/config"
	"tplgo/pkg/logger"

	"entgo.io/ent/dialect"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

// DB ent client.
var DB *Client

func InitDB() {
	DB = NewClient(Driver(dialect.DebugWithContext(yiigo.EntDriver(), func(ctx context.Context, v ...any) {
		if cfg.ENV.Debug {
			logger.Info(ctx, "SQL info", zap.String("SQL", fmt.Sprint(v...)))
		}
	})))
}

// Transaction Executes ent transaction with callback function.
// The provided context is used until the transaction is committed or rolledback.
func Transaction(ctx context.Context, f func(ctx context.Context, tx *Tx) error) error {
	tx, err := DB.Tx(ctx)

	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			logger.Err(ctx, "ent tx panic", zap.Any("error", r), zap.ByteString("stack", debug.Stack()))

			rollback(ctx, tx)
		}
	}()

	if err = f(ctx, tx); err != nil {
		rollback(ctx, tx)

		return err
	}

	if err = tx.Commit(); err != nil {
		rollback(ctx, tx)

		return err
	}

	return nil
}

func rollback(ctx context.Context, tx *Tx) {
	if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
		logger.Err(ctx, "err ent tx rollback", zap.Error(err))
	}
}
