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
}

type QuerierWithPrepareT[T PGXQuerier] interface {
	QuerierT[T]
	Prepare(ctx context.Context, name string, query litsql.Query) (*Stmt[T], error)
}

type QuerierPoolT[T PGXQuerier] interface {
	QuerierT[T]
	BeginTx(ctx context.Context, opts pgx.TxOptions) (*TxT[pgx.Tx], error)
}

type QuerierPoolConnT[T PGXQuerier] interface {
	QuerierT[T]
	QuerierPoolT[T]
	BeginTx(ctx context.Context, opts pgx.TxOptions) (*TxT[pgx.Tx], error)
}

type QuerierConnT[T PGXQuerier] interface {
	QuerierWithPrepareT[T]
	QuerierPoolConnT[T]
}

type QuerierStmtT interface {
	Query(ctx context.Context, params any) (pgx.Rows, error)
	QueryRow(ctx context.Context, params any) (pgx.Row, error)
	Exec(ctx context.Context, params any) (pgconn.CommandTag, error)
}

type QuerierTxT[T PGXQuerier] interface {
	QuerierWithPrepareT[T]
	Begin(ctx context.Context) (*TxT[pgx.Tx], error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
