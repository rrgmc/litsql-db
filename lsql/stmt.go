package lsql

import (
	"context"
	"database/sql"

	"github.com/rrgmc/litsql/sq"
)

type Stmt[T QuerierStmt] struct {
	stmt         T
	args         []any
	queryHandler sq.Handler
}

func NewStmt[T QuerierStmt](querier T, args []any, options ...Option) *Stmt[T] {
	var optns dbOptions
	for _, opt := range options {
		opt(&optns)
	}

	if optns.queryHandler == nil {
		optns.queryHandler = sq.NewHandler()
	}

	return &Stmt[T]{
		stmt:         querier,
		args:         args,
		queryHandler: optns.queryHandler,
	}
}

func (d *Stmt[T]) Handler() QuerierStmt {
	return d.stmt
}

func (d *Stmt[T]) Query(ctx context.Context, params any) (*sql.Rows, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	return d.stmt.QueryContext(ctx, args...)
}

func (d *Stmt[T]) QueryRow(ctx context.Context, params any) (*sql.Row, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	row := d.stmt.QueryRowContext(ctx, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func (d *Stmt[T]) Exec(ctx context.Context, params any) (sql.Result, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	return d.stmt.ExecContext(ctx, args...)
}

func (d *Stmt[T]) buildArgs(params any) ([]any, error) {
	return d.queryHandler.ParseArgs(d.args, params)
}
