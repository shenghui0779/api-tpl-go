package db

import (
	"api/ent"
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

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
