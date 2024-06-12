package lpgxt

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Pool wraps any implementation of [PGXQuerierPool].
type Pool[T PGXQuerierPool] struct {
	*baseQuerier[T]
}

// NewPool wraps any implementation of [PGXQuerierPool].
func NewPool[T PGXQuerierPool](querier T, options ...Option) *Pool[T] {
	return &Pool[T]{
		baseQuerier: newBaseQuerier[T](querier, options...),
	}
}

func (d *Pool[T]) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (*Tx[pgx.Tx], error) {
	tx, err := d.baseQuerier.querier.BeginTx(ctx, txOptions)
	if err != nil {
		return nil, err
	}
	return &Tx[pgx.Tx]{
		baseQuerierWithPrepare: &baseQuerierWithPrepare[pgx.Tx]{
			baseQuerier: &baseQuerier[pgx.Tx]{
				queryHandler: d.queryHandler,
				querier:      tx,
			},
		},
	}, nil
}
