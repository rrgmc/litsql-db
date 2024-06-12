package lsqlt

import (
	"context"
	"database/sql"

	"github.com/rrgmc/litsql"
)

type QuerierT[T SQLQuerier] interface {
	Query(ctx context.Context, query litsql.Query, params any) (*sql.Rows, error)
	QueryRow(ctx context.Context, query litsql.Query, params any) (*sql.Row, error)
	Exec(ctx context.Context, query litsql.Query, params any) (sql.Result, error)
	Prepare(ctx context.Context, query litsql.Query) (*StmtT[*sql.Stmt], error)
	Stmt(ctx context.Context, stmt *StmtT[*sql.Stmt]) *StmtT[*sql.Stmt] // allows matching both DB and Tx
}

type QuerierDBT[T SQLQuerier] interface {
	QuerierT[T]
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*TxT[*sql.Tx], error)
}

type QuerierStmtT interface {
	Query(ctx context.Context, params any) (*sql.Rows, error)
	QueryRow(ctx context.Context, params any) (*sql.Row, error)
	Exec(ctx context.Context, params any) (sql.Result, error)
}

type QuerierTxT[T SQLQuerier] interface {
	QuerierT[T]
	Commit() error
	Rollback() error
}
