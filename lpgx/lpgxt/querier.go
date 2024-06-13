package lpgxt

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rrgmc/litsql"
)

type Querier interface {
	Query(ctx context.Context, query litsql.Query, params litsql.ArgValues) (pgx.Rows, error)
	QueryRow(ctx context.Context, query litsql.Query, params litsql.ArgValues) (pgx.Row, error)
	Exec(ctx context.Context, query litsql.Query, params litsql.ArgValues) (pgconn.CommandTag, error)
}

type QuerierWithPrepare[T PGXQuerier] interface {
	Querier
	Prepare(ctx context.Context, name string, query litsql.Query) (*Stmt[T], error)
}

type QuerierPool interface {
	Querier
	BeginTx(ctx context.Context, opts pgx.TxOptions) (*Tx[pgx.Tx], error)
}

type QuerierPoolConn interface {
	Querier
	QuerierPool
	BeginTx(ctx context.Context, opts pgx.TxOptions) (*Tx[pgx.Tx], error)
}

type QuerierConn[T PGXQuerier] interface {
	QuerierWithPrepare[T]
	QuerierPoolConn
}

type QuerierStmt interface {
	Query(ctx context.Context, params litsql.ArgValues) (pgx.Rows, error)
	QueryRow(ctx context.Context, params litsql.ArgValues) (pgx.Row, error)
	Exec(ctx context.Context, params litsql.ArgValues) (pgconn.CommandTag, error)
}

type QuerierTx[T PGXQuerier] interface {
	QuerierWithPrepare[T]
	Begin(ctx context.Context) (*Tx[pgx.Tx], error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
