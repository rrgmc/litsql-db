package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DBT[T PGXQuerierDB] struct {
	*BaseQuerier[T]
}

func NewDBT[T PGXQuerierDB](querier T, options ...Option) *DBT[T] {
	return &DBT[T]{
		BaseQuerier: NewBaseQuerier[T](querier, options...),
	}
}

func (d *DBT[T]) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (*TxT[pgx.Tx], error) {
	tx, err := d.BaseQuerier.querier.BeginTx(ctx, txOptions)
	if err != nil {
		return nil, err
	}
	return &TxT[pgx.Tx]{
		BaseQuerier: &BaseQuerier[pgx.Tx]{
			queryHandler: d.queryHandler,
			querier:      tx,
		},
	}, nil
}
