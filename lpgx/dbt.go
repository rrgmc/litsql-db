package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DBT[TT PGXQuerierTx, T PGXQuerierDB] struct {
	*BaseQuerier[T]
}

func NewDBT[TT PGXQuerierTx, T PGXQuerierDB](querier T, options ...Option) *DBT[TT, T] {
	return &DBT[TT, T]{
		BaseQuerier: NewBaseQuerier[T](querier, options...),
	}
}

func (d *DBT[TT, T]) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (*TxT[pgx.Tx], error) {
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
