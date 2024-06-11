package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type TxT[T SQLQuerierTx] struct {
	*BaseQuerier[T]
}

func NewTxT[T SQLQuerierTx](querier T, options ...Option) *TxT[T] {
	return &TxT[T]{
		BaseQuerier: NewBaseQuerier[T](querier, options...),
	}
}

func (d *TxT[T]) Begin(ctx context.Context) (*TxT[pgx.Tx], error) {
	tx, err := d.BaseQuerier.querier.Begin(ctx)
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

func (d *TxT[T]) Commit(ctx context.Context) error {
	return d.BaseQuerier.querier.Commit(ctx)
}

func (d *TxT[T]) Rollback(ctx context.Context) error {
	return d.BaseQuerier.querier.Rollback(ctx)
}
