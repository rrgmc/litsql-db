package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rrgmc/litsql/sq"
)

type Stmt[T PGXQuerier] struct {
	stmt         T
	desc         *pgconn.StatementDescription
	args         []any
	queryHandler sq.Handler
}

// NewStmt wraps a PGXQuerier with a statement description from a Prepare call.
func NewStmt[T PGXQuerier](querier T, desc *pgconn.StatementDescription, args []any, options ...Option) *Stmt[T] {
	var optns dbOptions
	for _, opt := range options {
		opt(&optns)
	}

	if optns.queryHandler == nil {
		optns.queryHandler = sq.NewHandler()
	}

	return &Stmt[T]{
		stmt:         querier,
		desc:         desc,
		args:         args,
		queryHandler: optns.queryHandler,
	}
}

func (d *Stmt[T]) Handler() T {
	return d.stmt
}

func (d *Stmt[T]) Query(ctx context.Context, params any) (pgx.Rows, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	return d.stmt.Query(ctx, d.desc.Name, args...)
}

func (d *Stmt[T]) QueryRow(ctx context.Context, params any) (pgx.Row, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	return d.stmt.QueryRow(ctx, d.desc.Name, args...), nil
}

func (d *Stmt[T]) Exec(ctx context.Context, params any) (pgconn.CommandTag, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	return d.stmt.Exec(ctx, d.desc.Name, args...)
}

func (d *Stmt[T]) buildArgs(params any) ([]any, error) {
	return d.queryHandler.ParseArgs(d.args, params)
}
