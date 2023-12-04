package db

import (
	"api/ent"
	"api/lib/log"

	"context"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var cli *ent.Client

// Init 初始化Ent实例(如有多个实例，在此方法中初始化)
func Init() error {
	cfg := buildCfg(viper.GetString("db.driver"), viper.GetString("db.dsn"), viper.GetStringMap("db.options"))

	db, err := New(cfg)
	if err != nil {
		return err
	}

	cli = ent.NewClient(
		ent.Driver(dialect.DebugWithContext(
			entsql.OpenDB(cfg.Driver, db),
			func(ctx context.Context, v ...any) {
				log.Info(ctx, "SQL info", zap.String("SQL", fmt.Sprint(v...)))
			}),
		),
	)

	return nil
}

// Client 返回Ent实例
func Client() *ent.Client {
	return cli
}

// Transaction Executes ent transaction with callback function.
// The provided context is used until the transaction is committed or rolledback.
func Transaction(ctx context.Context, f func(ctx context.Context, tx *ent.Tx) error) error {
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
