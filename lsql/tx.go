package lsql

import (
	"context"
	"database/sql"
)

type Tx[T QuerierTx] struct {
	*BaseQuerier[T]
}

func NewTx[T QuerierTx](querier T, options ...Option) *Tx[T] {
	return &Tx[T]{
		BaseQuerier: NewBaseQuerier[T](querier, options...),
	}
}

func (d *Tx[T]) Commit() error {
	return d.BaseQuerier.querier.Commit()
}

func (d *Tx[T]) Rollback() error {
	return d.BaseQuerier.querier.Rollback()
}

func (d *Tx[T]) Stmt(ctx context.Context, stmt *Stmt[*sql.Stmt]) *Stmt[*sql.Stmt] {
	return &Stmt[*sql.Stmt]{
		stmt:         d.BaseQuerier.querier.StmtContext(ctx, stmt.stmt),
		args:         stmt.args,
		queryHandler: stmt.queryHandler,
	}
}
