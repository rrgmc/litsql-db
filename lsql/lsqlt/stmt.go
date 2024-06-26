package lsqlt

import (
	"context"
	"database/sql"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

// Stmt wraps any implementation of [SQLQuerierStmt].
type Stmt[T SQLQuerierStmt] struct {
	stmt         T
	args         []any
	queryHandler sq.Handler
}

// NewStmt wraps any implementation of [SQLQuerierStmt].
func NewStmt[T SQLQuerierStmt](querier T, args []any, options ...Option) *Stmt[T] {
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

func (d *Stmt[T]) Handler() T {
	return d.stmt
}

func (d *Stmt[T]) Query(ctx context.Context, params litsql.ArgValues) (*sql.Rows, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	return d.stmt.QueryContext(ctx, args...)
}

func (d *Stmt[T]) QueryRow(ctx context.Context, params litsql.ArgValues) (*sql.Row, error) {
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

func (d *Stmt[T]) Exec(ctx context.Context, params litsql.ArgValues) (sql.Result, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	return d.stmt.ExecContext(ctx, args...)
}

func (d *Stmt[T]) buildArgs(params litsql.ArgValues) ([]any, error) {
	return d.queryHandler.ParseArgs(d.args, params)
}
