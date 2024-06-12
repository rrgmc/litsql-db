package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
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

type PGXQuerierPoolConn interface {
	PGXQuerier
	PGXQuerierPool
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type PGXQuerierConn interface {
	PGXQuerierWithPrepare
	PGXQuerierPoolConn
}

type PGXQuerierTx interface {
	PGXQuerierWithPrepare
	Begin(ctx context.Context) (pgx.Tx, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

var (
	_ PGXQuerier = (*pgx.Conn)(nil)
	_ PGXQuerier = (pgx.Tx)(nil)
	_ PGXQuerier = (*pgxpool.Pool)(nil)
	_ PGXQuerier = (*pgxpool.Conn)(nil)
	_ PGXQuerier = (*pgxpool.Tx)(nil)

	_ PGXQuerierWithPrepare = (*pgx.Conn)(nil)
	_ PGXQuerierWithPrepare = (pgx.Tx)(nil)
	_ PGXQuerierWithPrepare = (*pgxpool.Tx)(nil)

	_ PGXQuerierPool = (*pgxpool.Pool)(nil)

	_ PGXQuerierConn     = (*pgx.Conn)(nil)
	_ PGXQuerierPoolConn = (*pgxpool.Conn)(nil)

	_ PGXQuerierTx = (pgx.Tx)(nil)
	_ PGXQuerierTx = (*pgxpool.Tx)(nil)
)
