package lsqlt

import (
	"context"
	"database/sql"
)

// TxT wraps any implementation of [SQLQuerierTx].
type TxT[T SQLQuerierTx] struct {
	*baseQuerier[T]
}

// NewTxT wraps any implementation of [SQLQuerierTx].
func NewTxT[T SQLQuerierTx](querier T, options ...Option) *TxT[T] {
	return &TxT[T]{
		baseQuerier: newBaseQuerier[T](querier, options...),
	}
}

func (d *TxT[T]) Commit() error {
	return d.baseQuerier.querier.Commit()
}

func (d *TxT[T]) Rollback() error {
	return d.baseQuerier.querier.Rollback()
}

func (d *TxT[T]) Stmt(ctx context.Context, stmt *StmtT[*sql.Stmt]) *StmtT[*sql.Stmt] {
	return &StmtT[*sql.Stmt]{
		stmt:         d.baseQuerier.querier.StmtContext(ctx, stmt.stmt),
		args:         stmt.args,
		queryHandler: stmt.queryHandler,
	}
}
