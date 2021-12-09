package ent

import (
	"context"
	"database/sql"
	"runtime/debug"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

// DB ent client.
var DB *Client

func InitDB() {
	DB = NewClient(Driver(yiigo.EntDriver()), Debug())
}

// Transaction Executes ent transaction with callback function.
// The provided context is used until the transaction is committed or rolledback.
func Transaction(ctx context.Context, process func(ctx context.Context, tx *Tx) error) error {
	tx, err := DB.Tx(ctx)

	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			yiigo.Logger().Fatal("ent transaction process panic", zap.Any("error", r), zap.ByteString("stack", debug.Stack()))

			txRollback(tx)
		}
	}()

	if err = process(ctx, tx); err != nil {
		txRollback(tx)

		return err
	}

	if err = tx.Commit(); err != nil {
		txRollback(tx)

		return err
	}

	return nil
}

func txRollback(tx *Tx) {
	if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
		yiigo.Logger().Error("ent transaction rollback error", zap.Error(err))
	}
}
