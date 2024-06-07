package lsql

import (
	"context"
	"database/sql"
)

type DBT[ST SQLQuerierStmt, T SQLQuerierDB] struct {
	*BaseQuerier[T]
}

func NewDBT[ST SQLQuerierStmt, T SQLQuerierDB](querier T, options ...Option) *DBT[ST, T] {
	return &DBT[ST, T]{
		BaseQuerier: NewBaseQuerier[T](querier, options...),
	}
}

func (d *DBT[ST, T]) Stmt(ctx context.Context, stmt *StmtT[ST]) *StmtT[ST] {
	// return the same instance, as we are not a transaction.
	return stmt
}

func (d *DBT[ST, T]) BeginTx(ctx context.Context, opts *sql.TxOptions) (*TxT[*sql.Tx], error) {
	tx, err := d.BaseQuerier.querier.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &TxT[*sql.Tx]{
		BaseQuerier: &BaseQuerier[*sql.Tx]{
			queryHandler: d.queryHandler,
			querier:      tx,
		},
	}, nil
}
