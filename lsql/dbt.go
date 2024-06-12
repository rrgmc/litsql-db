package lsql

import (
	"context"
	"database/sql"
)

// DBT wraps any implementation of [SQLQuerierDB].
type DBT[T SQLQuerierDB] struct {
	*baseQuerier[T]
}

// NewDBT wraps any implementation of [SQLQuerierDB].
func NewDBT[T SQLQuerierDB](querier T, options ...Option) *DBT[T] {
	return &DBT[T]{
		baseQuerier: newBaseQuerier[T](querier, options...),
	}
}

func (d *DBT[T]) Stmt(_ context.Context, stmt *StmtT[*sql.Stmt]) *StmtT[*sql.Stmt] {
	// return the same instance, as we are not a transaction.
	return stmt
}

func (d *DBT[T]) BeginTx(ctx context.Context, opts *sql.TxOptions) (*TxT[*sql.Tx], error) {
	tx, err := d.baseQuerier.querier.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &TxT[*sql.Tx]{
		baseQuerier: &baseQuerier[*sql.Tx]{
			queryHandler: d.queryHandler,
			querier:      tx,
		},
	}, nil
}
