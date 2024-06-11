package lpgx

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rrgmc/litsql"
)

type QuerierT[ST SQLQuerier] interface {
	Query(ctx context.Context, query litsql.Query, params any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query litsql.Query, params any) (pgx.Row, error)
	Exec(ctx context.Context, query litsql.Query, params any) (pgconn.CommandTag, error)
	Prepare(ctx context.Context, name string, query litsql.Query) (*StmtT[ST], error)
	// Stmt(ctx context.Context, stmt *StmtT[ST]) *StmtT[ST] // allows matching both DB and Tx
}

type QuerierDBT[ST SQLQuerier, TT SQLQuerierTx] interface {
	QuerierT[ST]
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*TxT[TT], error)
}

type QuerierStmtT interface {
	Query(ctx context.Context, params any) (pgx.Rows, error)
	QueryRow(ctx context.Context, params any) (pgx.Row, error)
	Exec(ctx context.Context, params any) (pgconn.CommandTag, error)
}

type QuerierTxT[ST SQLQuerierTx] interface {
	QuerierT[ST]
	Begin(ctx context.Context) (*TxT[ST], error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	// Stmt(ctx context.Context, stmt *StmtT[ST]) *StmtT[ST]
}
