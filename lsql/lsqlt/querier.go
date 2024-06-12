package lsqlt

import (
	"context"
	"database/sql"

	"github.com/rrgmc/litsql"
)

type Querier interface {
	Query(ctx context.Context, query litsql.Query, params any) (*sql.Rows, error)
	QueryRow(ctx context.Context, query litsql.Query, params any) (*sql.Row, error)
	Exec(ctx context.Context, query litsql.Query, params any) (sql.Result, error)
	Prepare(ctx context.Context, query litsql.Query) (*Stmt[*sql.Stmt], error)
	Stmt(ctx context.Context, stmt *Stmt[*sql.Stmt]) *Stmt[*sql.Stmt] // allows matching both DB and Tx
}

type QuerierDB interface {
	Querier
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx[*sql.Tx], error)
}

type QuerierStmt interface {
	Query(ctx context.Context, params any) (*sql.Rows, error)
	QueryRow(ctx context.Context, params any) (*sql.Row, error)
	Exec(ctx context.Context, params any) (sql.Result, error)
}

type QuerierTx interface {
	Querier
	Commit() error
	Rollback() error
}
