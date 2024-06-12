package lsql

import (
	"context"
	"database/sql"
)

type SQLQuerier interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type SQLQuerierDB interface {
	SQLQuerier
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type SQLQuerierStmt interface {
	QueryContext(ctx context.Context, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, args ...any) *sql.Row
	ExecContext(ctx context.Context, args ...any) (sql.Result, error)
}

type SQLQuerierTx interface {
	SQLQuerier
	Commit() error
	Rollback() error
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
}
