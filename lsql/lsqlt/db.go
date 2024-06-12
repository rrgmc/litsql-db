package lsqlt

import (
	"context"
	"database/sql"
)

// DB wraps any implementation of [SQLQuerierDB].
type DB[T SQLQuerierDB] struct {
	*baseQuerier[T]
}

// NewDB wraps any implementation of [SQLQuerierDB].
func NewDB[T SQLQuerierDB](querier T, options ...Option) *DB[T] {
	return &DB[T]{
		baseQuerier: newBaseQuerier[T](querier, options...),
	}
}

func (d *DB[T]) Stmt(_ context.Context, stmt *Stmt[*sql.Stmt]) *Stmt[*sql.Stmt] {
	// return the same instance, as we are not a transaction.
	return stmt
}

func (d *DB[T]) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx[*sql.Tx], error) {
	tx, err := d.baseQuerier.querier.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &Tx[*sql.Tx]{
		baseQuerier: &baseQuerier[*sql.Tx]{
			queryHandler: d.queryHandler,
			querier:      tx,
		},
	}, nil
}
