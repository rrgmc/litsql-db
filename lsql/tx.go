package lsql

import (
	"context"
)

type Tx[T QuerierTx] struct {
	*DBQuerier[T]
}

func NewTx[T QuerierTx](querier T, options ...Option) *Tx[T] {
	return &Tx[T]{
		DBQuerier: NewDBQuerier[T](querier, options...),
	}
}

func (d *Tx[T]) Commit() error {
	return d.DBQuerier.querier.Commit()
}

func (d *Tx[T]) Rollback() error {
	return d.DBQuerier.querier.Rollback()
}

func (d *Tx[T]) Stmt(ctx context.Context, stmt *Stmt) *Stmt {
	return &Stmt{
		stmt:         d.DBQuerier.querier.StmtContext(ctx, stmt.stmt),
		args:         stmt.args,
		queryHandler: stmt.queryHandler,
	}
}
