package lsql

import (
	"context"
	"database/sql"

	"github.com/rrgmc/litsql/sq"
)

type StmtT[T SQLQuerierStmt] struct {
	stmt         T
	args         []any
	queryHandler sq.Handler
}

// NewStmtT wraps any implementation of [SQLQuerierStmt].
func NewStmtT[T SQLQuerierStmt](querier T, args []any, options ...Option) *StmtT[T] {
	var optns dbOptions
	for _, opt := range options {
		opt(&optns)
	}

	if optns.queryHandler == nil {
		optns.queryHandler = sq.NewHandler()
	}

	return &StmtT[T]{
		stmt:         querier,
		args:         args,
		queryHandler: optns.queryHandler,
	}
}

func (d *StmtT[T]) Handler() T {
	return d.stmt
}

func (d *StmtT[T]) Query(ctx context.Context, params any) (*sql.Rows, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	return d.stmt.QueryContext(ctx, args...)
}

func (d *StmtT[T]) QueryRow(ctx context.Context, params any) (*sql.Row, error) {
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

func (d *StmtT[T]) Exec(ctx context.Context, params any) (sql.Result, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	return d.stmt.ExecContext(ctx, args...)
}

func (d *StmtT[T]) buildArgs(params any) ([]any, error) {
	return d.queryHandler.ParseArgs(d.args, params)
}
