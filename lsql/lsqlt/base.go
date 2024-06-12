package lsqlt

import (
	"context"
	"database/sql"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

type baseQuerier[T SQLQuerier] struct {
	queryHandler sq.Handler
	querier      T
}

func newBaseQuerier[T SQLQuerier](querier T, options ...Option) *baseQuerier[T] {
	var optns dbOptions
	for _, opt := range options {
		opt(&optns)
	}

	if optns.queryHandler == nil {
		optns.queryHandler = sq.NewHandler()
	}

	return &baseQuerier[T]{
		querier:      querier,
		queryHandler: optns.queryHandler,
	}
}

func (d *baseQuerier[T]) Handler() T {
	return d.querier
}

func (d *baseQuerier[T]) Query(ctx context.Context, query litsql.Query, params any) (*sql.Rows, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return nil, err
	}
	return d.querier.QueryContext(ctx, qstr, args...)
}

func (d *baseQuerier[T]) QueryRow(ctx context.Context, query litsql.Query, params any) (*sql.Row, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return nil, err
	}
	row := d.querier.QueryRowContext(ctx, qstr, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func (d *baseQuerier[T]) Exec(ctx context.Context, query litsql.Query, params any) (sql.Result, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return nil, err
	}
	return d.querier.ExecContext(ctx, qstr, args...)
}

func (d *baseQuerier[T]) Prepare(ctx context.Context, query litsql.Query) (*StmtT[*sql.Stmt], error) {
	qstr, args, err := d.queryHandler.Build(query)
	if err != nil {
		return nil, err
	}
	stmt, err := d.querier.PrepareContext(ctx, qstr)
	if err != nil {
		return nil, err
	}
	return &StmtT[*sql.Stmt]{
		stmt:         stmt,
		args:         args,
		queryHandler: d.queryHandler,
	}, nil
}

func (d *baseQuerier[T]) buildQuery(query litsql.Query, params any) (string, []any, error) {
	return d.queryHandler.Build(query,
		sq.WithParseArgs(params),
	)
}

