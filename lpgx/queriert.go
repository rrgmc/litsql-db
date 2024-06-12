package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rrgmc/litsql"
)

type QuerierT[T PGXQuerier] interface {
	Query(ctx context.Context, query litsql.Query, params any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query litsql.Query, params any) (pgx.Row, error)
	Exec(ctx context.Context, query litsql.Query, params any) (pgconn.CommandTag, error)
	Prepare(ctx context.Context, name string, query litsql.Query) (*Stmt[T], error)
}

type QuerierDBT[T PGXQuerier] interface {
	QuerierT[T]
	BeginTx(ctx context.Context, opts pgx.TxOptions) (*TxT[pgx.Tx], error)
}

type QuerierStmtT interface {
	Query(ctx context.Context, params any) (pgx.Rows, error)
	QueryRow(ctx context.Context, params any) (pgx.Row, error)
	Exec(ctx context.Context, params any) (pgconn.CommandTag, error)
}

type QuerierTxT[T PGXQuerier] interface {
	QuerierT[T]
	Begin(ctx context.Context) (*TxT[pgx.Tx], error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
