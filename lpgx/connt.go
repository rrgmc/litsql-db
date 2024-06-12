package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// ConnT wraps any implementation of [PGXQuerierConn].
type ConnT[T PGXQuerierConn] struct {
	*baseQuerierWithPrepare[T]
}

// NewConnT wraps any implementation of [PGXQuerierConn].
func NewConnT[T PGXQuerierConn](querier T, options ...Option) *ConnT[T] {
	return &ConnT[T]{
		baseQuerierWithPrepare: newBaseQuerierWithPrepare[T](querier, options...),
	}
}

func (d *ConnT[T]) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (*TxT[pgx.Tx], error) {
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

// PoolConnT wraps any implementation of [PGXQuerierPoolConn].
type PoolConnT[T PGXQuerierPoolConn] struct {
	*baseQuerier[T]
}

// NewPoolConnT wraps any implementation of [PGXQuerierPoolConn].
func NewPoolConnT[T PGXQuerierPoolConn](querier T, options ...Option) *PoolConnT[T] {
	return &PoolConnT[T]{
		baseQuerier: newBaseQuerier[T](querier, options...),
	}
}

func (d *PoolConnT[T]) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (*TxT[pgx.Tx], error) {
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
