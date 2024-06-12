package lpgxt

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// PoolT wraps any implementation of [PGXQuerierPool].
type PoolT[T PGXQuerierPool] struct {
	*baseQuerier[T]
}

// NewPoolT wraps any implementation of [PGXQuerierPool].
func NewPoolT[T PGXQuerierPool](querier T, options ...Option) *PoolT[T] {
	return &PoolT[T]{
		baseQuerier: newBaseQuerier[T](querier, options...),
	}
}

func (d *PoolT[T]) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (*TxT[pgx.Tx], error) {
	tx, err := d.baseQuerier.querier.BeginTx(ctx, txOptions)
	if err != nil {
		return nil, err
	}
	return &TxT[pgx.Tx]{
		baseQuerierWithPrepare: &baseQuerierWithPrepare[pgx.Tx]{
			baseQuerier: &baseQuerier[pgx.Tx]{
				queryHandler: d.queryHandler,
				querier:      tx,
			},
		},
	}, nil
}
