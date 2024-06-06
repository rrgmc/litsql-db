package lsql

import (
	"context"
	"database/sql"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

type Stmt struct {
	stmt         *sql.Stmt
	args         []any
	queryHandler sq.Handler
}

func NewStmt(querier *sql.Stmt, args []any, options ...Option) *Stmt {
	var optns dbOptions
	for _, opt := range options {
		opt(&optns)
	}

	if optns.queryHandler == nil {
		optns.queryHandler = sq.NewHandler()
	}

	return &Stmt{
		stmt:         querier,
		args:         args,
		queryHandler: optns.queryHandler,
	}
}

func (d *Stmt) Handler() *sql.Stmt {
	return d.stmt
}

func (d *Stmt) Query(ctx context.Context, params any) (*sql.Rows, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	return d.stmt.QueryContext(ctx, args...)
}

func (d *Stmt) QueryRow(ctx context.Context, query litsql.Query, params any) (*sql.Row, error) {
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

func (d *Stmt) Exec(ctx context.Context, query litsql.Query, params any) (sql.Result, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	return d.stmt.ExecContext(ctx, args...)
}

func (d *Stmt) buildArgs(params any) ([]any, error) {
	return d.queryHandler.ParseArgs(d.args, params)
}
