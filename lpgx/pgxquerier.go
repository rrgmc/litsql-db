package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PGXQuerier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (commandTag pgconn.CommandTag, err error)
}

type PGXQuerierWithPrepare interface {
	PGXQuerier
	Prepare(ctx context.Context, name, sql string) (sd *pgconn.StatementDescription, err error)
}

type PGXQuerierPool interface {
	PGXQuerier
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type PGXQuerierConn interface {
	PGXQuerierWithPrepare
	PGXQuerierPool
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type PGXQuerierTx interface {
	PGXQuerierWithPrepare
	Begin(ctx context.Context) (pgx.Tx, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
