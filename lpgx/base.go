package lpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

type baseQuerier[T PGXQuerier] struct {
	queryHandler sq.Handler
	querier      T
}

func newBaseQuerier[T PGXQuerier](querier T, options ...Option) *baseQuerier[T] {
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

func (d *baseQuerier[T]) Query(ctx context.Context, query litsql.Query, params any) (pgx.Rows, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return nil, err
	}
	return d.querier.Query(ctx, qstr, args...)
}

func (d *baseQuerier[T]) QueryRow(ctx context.Context, query litsql.Query, params any) (pgx.Row, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return nil, err
	}
	return d.querier.QueryRow(ctx, qstr, args...), nil
}

func (d *baseQuerier[T]) Exec(ctx context.Context, query litsql.Query, params any) (pgconn.CommandTag, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	return d.querier.Exec(ctx, qstr, args...)
}

func (d *baseQuerier[T]) buildQuery(query litsql.Query, params any) (string, []any, error) {
	return d.queryHandler.Build(query,
		sq.WithParseArgs(params),
	)
}

// with prepare

type baseQuerierWithPrepare[T PGXQuerierWithPrepare] struct {
	*baseQuerier[T]
}

func newBaseQuerierWithPrepare[T PGXQuerierWithPrepare](querier T, options ...Option) *baseQuerierWithPrepare[T] {
	return &baseQuerierWithPrepare[T]{
		baseQuerier: newBaseQuerier[T](querier, options...),
	}
}

func (d *baseQuerierWithPrepare[T]) Prepare(ctx context.Context, name string, query litsql.Query) (*Stmt[T], error) {
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
