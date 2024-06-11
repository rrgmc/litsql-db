package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// PGXQuerier is something that lpgx can query and get the pgx.Rows from.
// For example, it can be: pgx.Conn or pgx.Tx.
type PGXQuerier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (commandTag pgconn.CommandTag, err error)
	Prepare(ctx context.Context, name, sql string) (sd *pgconn.StatementDescription, err error)
}

type PGXQuerierDB interface {
	PGXQuerier
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type PGXQuerierTx interface {
	PGXQuerier
	Begin(ctx context.Context) (pgx.Tx, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
