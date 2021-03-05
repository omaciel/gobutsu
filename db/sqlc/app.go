package db

import (
	"context"
	"database/sql"
	"fmt"
)

// App provides all functions to execute db queries and transactions
type App struct {
	*Queries
	db *sql.DB
}

// NewApp creates a new App
func NewApp(db *sql.DB) *App {
	return &App{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (app *App) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := app.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
