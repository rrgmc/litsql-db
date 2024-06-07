package lsql

import (
	"context"
	"database/sql"

	"github.com/rrgmc/litsql"
)

type QuerierT[ST SQLQuerierStmt] interface {
	Query(ctx context.Context, query litsql.Query, params any) (*sql.Rows, error)
	QueryRow(ctx context.Context, query litsql.Query, params any) (*sql.Row, error)
	Exec(ctx context.Context, query litsql.Query, params any) (sql.Result, error)
	Prepare(ctx context.Context, query litsql.Query) (*StmtT[ST], error)
	Stmt(ctx context.Context, stmt *StmtT[ST]) *StmtT[ST] // allows matching both DB and Tx
}

type QuerierDBT[ST SQLQuerierStmt, TT SQLQuerierTx] interface {
	QuerierT[ST]
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*TxT[TT], error)
}

type QuerierStmtT interface {
	Query(ctx context.Context, params any) (*sql.Rows, error)
	QueryRow(ctx context.Context, params any) (*sql.Row, error)
	Exec(ctx context.Context, params any) (sql.Result, error)
}

type QuerierTxT[ST SQLQuerierStmt] interface {
	QuerierT[ST]
	Commit() error
	Rollback() error
	Stmt(ctx context.Context, stmt *StmtT[ST]) *StmtT[ST]
}
