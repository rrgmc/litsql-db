package lsql

import (
	"context"
	"database/sql"
)

// SQLQuerier is implemented by [sql.DB], [sql.Tx] and [sql.Conn].
type SQLQuerier interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

// SQLQuerierDB is implemented by [sql.DB] and [sql.Conn].
type SQLQuerierDB interface {
	SQLQuerier
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// SQLQuerierStmt is implemented by [sql.Stmt].
type SQLQuerierStmt interface {
	QueryContext(ctx context.Context, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, args ...any) *sql.Row
	ExecContext(ctx context.Context, args ...any) (sql.Result, error)
}

// SQLQuerierTx is implemented by [sql.Tx].
type SQLQuerierTx interface {
	SQLQuerier
	Commit() error
	Rollback() error
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
}

var (
	_ SQLQuerier = (*sql.DB)(nil)
	_ SQLQuerier = (*sql.Tx)(nil)
	_ SQLQuerier = (*sql.Conn)(nil)

	_ SQLQuerierDB = (*sql.DB)(nil)
	_ SQLQuerierDB = (*sql.Conn)(nil)

	_ SQLQuerierTx = (*sql.Tx)(nil)

	_ SQLQuerierStmt = (*sql.Stmt)(nil)
)
