package lsqlt

import (
	"context"
	"database/sql"
)

// Tx wraps any implementation of [SQLQuerierTx].
type Tx[T SQLQuerierTx] struct {
	*baseQuerier[T]
}

// NewTx wraps any implementation of [SQLQuerierTx].
func NewTx[T SQLQuerierTx](querier T, options ...Option) *Tx[T] {
	return &Tx[T]{
		baseQuerier: newBaseQuerier[T](querier, options...),
	}
}

func (d *Tx[T]) Commit() error {
	return d.baseQuerier.querier.Commit()
}

func (d *Tx[T]) Rollback() error {
	return d.baseQuerier.querier.Rollback()
}

func (d *Tx[T]) Stmt(ctx context.Context, stmt *Stmt[*sql.Stmt]) *Stmt[*sql.Stmt] {
	return &Stmt[*sql.Stmt]{
		stmt:         d.baseQuerier.querier.StmtContext(ctx, stmt.stmt),
		args:         stmt.args,
		queryHandler: stmt.queryHandler,
	}
}
