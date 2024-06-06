package lsql

import (
	"context"
)

type DBQuerierTx[T QuerierTx] struct {
	*DBQuerier[T]
}

func NewDBQuerierTx[T QuerierTx](querier T, options ...Option) *DBQuerierTx[T] {
	return &DBQuerierTx[T]{
		DBQuerier: NewDBQuerier[T](querier, options...),
	}
}

func (d *DBQuerierTx[T]) Commit() error {
	return d.DBQuerier.querier.Commit()
}

func (d *DBQuerierTx[T]) Rollback() error {
	return d.DBQuerier.querier.Rollback()
}

func (d *DBQuerierTx[T]) Stmt(ctx context.Context, stmt *DBQuerierStmt) *DBQuerierStmt {
	return &DBQuerierStmt{
		stmt:         d.DBQuerier.querier.StmtContext(ctx, stmt.stmt),
		args:         stmt.args,
		queryHandler: stmt.queryHandler,
	}
}
