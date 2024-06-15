package ent

import (
	"context"
	"fmt"

	"api/lib/db"
)

var DB *Client

func Init() error {
	driver, err := db.Init()
	if err != nil {
		return err
	}
	DB = NewClient(Driver(driver))
	return nil
}

// Transaction Executes ent transaction with callback function.
// The provided context is used until the transaction is committed or rolledback.
func Transaction(ctx context.Context, fn func(ctx context.Context, tx *Tx) error) error {
	tx, err := DB.Tx(ctx)
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

	if err = fn(ctx, tx); err != nil {
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
