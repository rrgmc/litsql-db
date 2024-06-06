package lsql

import (
	"context"
	"database/sql"
)

type DBQuerierDB[T QuerierDB] struct {
	*DBQuerier[T]
}

func NewDBQuerierDB[T QuerierDB](querier T, options ...Option) *DBQuerierDB[T] {
	return &DBQuerierDB[T]{
		DBQuerier: NewDBQuerier[T](querier, options...),
	}
}

func (d *DBQuerierDB[T]) BeginTx(ctx context.Context, opts *sql.TxOptions) (*DBQuerierTx[*sql.Tx], error) {
	tx, err := d.DBQuerier.querier.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &DBQuerierTx[*sql.Tx]{
		DBQuerier: &DBQuerier[*sql.Tx]{
			queryHandler: d.queryHandler,
			querier:      tx,
		},
	}, nil
}
