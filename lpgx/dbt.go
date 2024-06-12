package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// DBT wraps any implementation of [PGXQuerierDB].
type DBT[T PGXQuerierDB] struct {
	*baseQuerier[T]
}

// NewDBT wraps any implementation of [PGXQuerierDB].
func NewDBT[T PGXQuerierDB](querier T, options ...Option) *DBT[T] {
	return &DBT[T]{
		baseQuerier: newBaseQuerier[T](querier, options...),
	}
}

func (d *DBT[T]) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (*TxT[pgx.Tx], error) {
	tx, err := d.baseQuerier.querier.BeginTx(ctx, txOptions)
	if err != nil {
		return nil, err
	}
	return &TxT[pgx.Tx]{
		baseQuerier: &baseQuerier[pgx.Tx]{
			queryHandler: d.queryHandler,
			querier:      tx,
		},
	}, nil
}
