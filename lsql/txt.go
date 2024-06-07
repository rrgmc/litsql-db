package lsql

import (
	"context"
	"database/sql"
)

type TxT[T SQLQuerierTx] struct {
	*BaseQuerier[T]
}

func NewTxT[T SQLQuerierTx](querier T, options ...Option) *TxT[T] {
	return &TxT[T]{
		BaseQuerier: NewBaseQuerier[T](querier, options...),
	}
}

func (d *TxT[T]) Commit() error {
	return d.BaseQuerier.querier.Commit()
}

func (d *TxT[T]) Rollback() error {
	return d.BaseQuerier.querier.Rollback()
}

func (d *TxT[T]) Stmt(ctx context.Context, stmt *StmtT[*sql.Stmt]) *StmtT[*sql.Stmt] {
	return &StmtT[*sql.Stmt]{
		stmt:         d.BaseQuerier.querier.StmtContext(ctx, stmt.stmt),
		args:         stmt.args,
		queryHandler: stmt.queryHandler,
	}
}
