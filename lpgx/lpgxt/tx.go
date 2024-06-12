package lpgxt

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Tx wraps any implementation of [PGXQuerierTx].
type Tx[T PGXQuerierTx] struct {
	*baseQuerierWithPrepare[T]
}

// NewTx wraps any implementation of [PGXQuerierTx].
func NewTx[T PGXQuerierTx](querier T, options ...Option) *Tx[T] {
	return &Tx[T]{
		baseQuerierWithPrepare: newBaseQuerierWithPrepare[T](querier, options...),
	}
}

func (d *Tx[T]) Begin(ctx context.Context) (*Tx[pgx.Tx], error) {
	tx, err := d.baseQuerier.querier.Begin(ctx)
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

func (d *Tx[T]) Commit(ctx context.Context) error {
	return d.baseQuerier.querier.Commit(ctx)
}

func (d *Tx[T]) Rollback(ctx context.Context) error {
	return d.baseQuerier.querier.Rollback(ctx)
}
