package lpgxt

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Conn wraps any implementation of [PGXQuerierConn].
type Conn[T PGXQuerierConn] struct {
	*baseQuerierWithPrepare[T]
}

// NewConn wraps any implementation of [PGXQuerierConn].
func NewConn[T PGXQuerierConn](querier T, options ...Option) *Conn[T] {
	return &Conn[T]{
		baseQuerierWithPrepare: newBaseQuerierWithPrepare[T](querier, options...),
	}
}

func (d *Conn[T]) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (*Tx[pgx.Tx], error) {
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

// PoolConn wraps any implementation of [PGXQuerierPoolConn].
type PoolConn[T PGXQuerierPoolConn] struct {
	*baseQuerier[T]
}

// NewPoolConn wraps any implementation of [PGXQuerierPoolConn].
func NewPoolConn[T PGXQuerierPoolConn](querier T, options ...Option) *PoolConn[T] {
	return &PoolConn[T]{
		baseQuerier: newBaseQuerier[T](querier, options...),
	}
}

func (d *PoolConn[T]) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (*Tx[pgx.Tx], error) {
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
