package ent

import (
	"context"
	"fmt"

	cfg "api/config"
	"api/logger"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var cli *Client

// Init 初始化DB实例
func Init(db *sqlx.DB) {
	cli = NewClient(
		Driver(dialect.DebugWithContext(
			entsql.OpenDB(db.DriverName(), db.DB),
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
		// if panic, should rollback
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
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
