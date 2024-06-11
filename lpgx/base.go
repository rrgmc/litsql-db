package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

type BaseQuerier[T PGXQuerier] struct {
	queryHandler sq.Handler
	querier      T
}

func NewBaseQuerier[T PGXQuerier](querier T, options ...Option) *BaseQuerier[T] {
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

func (d *BaseQuerier[T]) Query(ctx context.Context, query litsql.Query, params any) (pgx.Rows, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return nil, err
	}
	return d.querier.Query(ctx, qstr, args...)
}

func (d *BaseQuerier[T]) QueryRow(ctx context.Context, query litsql.Query, params any) (pgx.Row, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return nil, err
	}
	return d.querier.QueryRow(ctx, qstr, args...), nil
}

func (d *BaseQuerier[T]) Exec(ctx context.Context, query litsql.Query, params any) (pgconn.CommandTag, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	return d.querier.Exec(ctx, qstr, args...)
}

func (d *BaseQuerier[T]) Prepare(ctx context.Context, name string, query litsql.Query) (*Stmt[T], error) {
	qstr, args, err := d.queryHandler.Build(query)
	if err != nil {
		return nil, err
	}
	desc, err := d.querier.Prepare(ctx, name, qstr)
	if err != nil {
		return nil, err
	}
	return &Stmt[T]{
		stmt:         d.querier,
		desc:         desc,
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
