package lsql

import (
	"context"
	"database/sql"
)

type DBT[T QuerierDB] struct {
	*BaseQuerier[T]
}

func NewDBT[T QuerierDB](querier T, options ...Option) *DBT[T] {
	return &DBT[T]{
		BaseQuerier: NewBaseQuerier[T](querier, options...),
	}
}

func (d *DBT[T]) BeginTx(ctx context.Context, opts *sql.TxOptions) (*TxT[*sql.Tx], error) {
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
