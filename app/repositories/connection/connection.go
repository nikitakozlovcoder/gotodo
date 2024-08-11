package connection

import (
	"context"
	"database/sql"
	"gotodo/app/repositories/transactionctx"
)

type Executor interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type DbConnection struct {
	DB *sql.DB
}

func (c *DbConnection) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return c.getExecutor(ctx).QueryContext(ctx, query, args...)
}

func (c *DbConnection) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return c.getExecutor(ctx).QueryRowContext(ctx, query, args...)
}

func (c *DbConnection) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.getExecutor(ctx).ExecContext(ctx, query, args...)
}

func (c *DbConnection) Close() {
	c.DB.Close()
}

func (c *DbConnection) getExecutor(ctx context.Context) Executor {
	info, err := transactionctx.GetTransactionInfo(ctx)
	if err == nil && info.IsActive {
		return info.Tx
	}

	return c.DB
}
