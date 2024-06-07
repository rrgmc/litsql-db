package lsql

import (
	"context"
	"database/sql"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

type BaseQuerier[T Querier] struct {
	queryHandler sq.Handler
	querier      T
}

func NewBaseQuerier[T Querier](querier T, options ...Option) *BaseQuerier[T] {
	var optns dbOptions
	for _, opt := range options {
		opt(&optns)
	}

	if optns.queryHandler == nil {
		optns.queryHandler = sq.NewHandler()
	}

	return &BaseQuerier[T]{
		querier:      querier,
		queryHandler: optns.queryHandler,
	}
}

func (d *BaseQuerier[T]) Handler() T {
	return d.querier
}

func (d *BaseQuerier[T]) Query(ctx context.Context, query litsql.Query, params any) (*sql.Rows, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return nil, err
	}
	return d.querier.QueryContext(ctx, qstr, args...)
}

func (d *BaseQuerier[T]) QueryRow(ctx context.Context, query litsql.Query, params any) (*sql.Row, error) {
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

func (d *BaseQuerier[T]) Exec(ctx context.Context, query litsql.Query, params any) (sql.Result, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return nil, err
	}
	return d.querier.ExecContext(ctx, qstr, args...)
}

func (d *BaseQuerier[T]) Prepare(ctx context.Context, query litsql.Query) (*Stmt[*sql.Stmt], error) {
	qstr, args, err := d.queryHandler.Build(query)
	if err != nil {
		return nil, err
	}
	stmt, err := d.querier.PrepareContext(ctx, qstr)
	if err != nil {
		return nil, err
	}
	return &Stmt[*sql.Stmt]{
		stmt:         stmt,
		args:         args,
		queryHandler: d.queryHandler,
	}, nil
}

func (d *BaseQuerier[T]) buildQuery(query litsql.Query, params any) (string, []any, error) {
	return d.queryHandler.Build(query,
		sq.WithParseArgs(params),
	)
}

type Option func(options *dbOptions)

type dbOptions struct {
	queryHandler sq.Handler
}

func WithQueryHandler(queryHandler sq.Handler) Option {
	return func(options *dbOptions) {
		options.queryHandler = queryHandler
	}
}