package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rrgmc/litsql/sq"
)

type StmtT[T SQLQuerier] struct {
	stmt         T
	desc         *pgconn.StatementDescription
	args         []any
	queryHandler sq.Handler
}

func NewStmtT[T SQLQuerier](querier T, desc *pgconn.StatementDescription, args []any, options ...Option) *StmtT[T] {
	var optns dbOptions
	for _, opt := range options {
		opt(&optns)
	}

	if optns.queryHandler == nil {
		optns.queryHandler = sq.NewHandler()
	}

	return &StmtT[T]{
		stmt:         querier,
		desc:         desc,
		args:         args,
		queryHandler: optns.queryHandler,
	}
}

func (d *StmtT[T]) Handler() SQLQuerier {
	return d.stmt
}

func (d *StmtT[T]) Query(ctx context.Context, params any) (pgx.Rows, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	return d.stmt.Query(ctx, d.desc.Name, args...)
}

func (d *StmtT[T]) QueryRow(ctx context.Context, params any) (pgx.Row, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return nil, err
	}
	return d.stmt.QueryRow(ctx, d.desc.Name, args...), nil
}

func (d *StmtT[T]) Exec(ctx context.Context, params any) (pgconn.CommandTag, error) {
	args, err := d.buildArgs(params)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	return d.stmt.Exec(ctx, d.desc.Name, args...)
}

func (d *StmtT[T]) buildArgs(params any) ([]any, error) {
	return d.queryHandler.ParseArgs(d.args, params)
}
