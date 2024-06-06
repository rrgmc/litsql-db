package lsql

import (
	"context"
	"database/sql"
)

type DB[T QuerierDB] struct {
	*DBQuerier[T]
}

func NewDB[T QuerierDB](querier T, options ...Option) *DB[T] {
	return &DB[T]{
		DBQuerier: NewDBQuerier[T](querier, options...),
	}
}

func (d *DB[T]) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx[*sql.Tx], error) {
	tx, err := d.DBQuerier.querier.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &Tx[*sql.Tx]{
		DBQuerier: &DBQuerier[*sql.Tx]{
			queryHandler: d.queryHandler,
			querier:      tx,
		},
	}, nil
}
