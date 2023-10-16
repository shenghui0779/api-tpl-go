package ent

import (
	"context"
	"database/sql"
	"fmt"
	"runtime/debug"

	cfg "api/config"
	"api/logger"

	"entgo.io/ent/dialect"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

var cli *Client

// Init 初始化DB实例
func Init() {
	cli = NewClient(
		Driver(dialect.DebugWithContext(
			yiigo.EntDriver(),
			func(ctx context.Context, v ...any) {
				if cfg.ENV.Debug {
					logger.Info(ctx, "SQL info", zap.String("SQL", fmt.Sprint(v...)))
				}
			}),
		),
	)
}

// DB 返回DB实例
func DB() *Client {
	return cli
}

// Transaction Executes ent transaction with callback function.
// The provided context is used until the transaction is committed or rolledback.
func Transaction(ctx context.Context, f func(ctx context.Context, tx *Tx) error) error {
	tx, err := cli.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			logger.Err(ctx, "executing transaction panic", zap.Any("error", r), zap.ByteString("stack", debug.Stack()))

			if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
				logger.Err(ctx, "err rolling back transaction when panic", zap.Error(err))
			}
		}
	}()

	if err = f(ctx, tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}

		return err
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}
