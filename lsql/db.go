package lsql

import (
	"context"
	"database/sql"
)

type DB[T QuerierDB] struct {
	*BaseQuerier[T]
}

func NewDB[T QuerierDB](querier T, options ...Option) *DB[T] {
	return &DB[T]{
		BaseQuerier: NewBaseQuerier[T](querier, options...),
	}
}

func (d *DB[T]) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx[*sql.Tx], error) {
	tx, err := d.BaseQuerier.querier.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &Tx[*sql.Tx]{
		BaseQuerier: &BaseQuerier[*sql.Tx]{
			queryHandler: d.queryHandler,
			querier:      tx,
		},
	}, nil
}
