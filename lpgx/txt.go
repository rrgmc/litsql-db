package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// TxT wraps any implementation of [PGXQuerierTx].
type TxT[T PGXQuerierTx] struct {
	*baseQuerier[T]
}

// NewTxT wraps any implementation of [PGXQuerierTx].
func NewTxT[T PGXQuerierTx](querier T, options ...Option) *TxT[T] {
	return &TxT[T]{
		baseQuerier: newBaseQuerier[T](querier, options...),
	}
}

func (d *TxT[T]) Begin(ctx context.Context) (*TxT[pgx.Tx], error) {
	tx, err := d.baseQuerier.querier.Begin(ctx)
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

func (d *TxT[T]) Commit(ctx context.Context) error {
	return d.baseQuerier.querier.Commit(ctx)
}

func (d *TxT[T]) Rollback(ctx context.Context) error {
	return d.baseQuerier.querier.Rollback(ctx)
}
