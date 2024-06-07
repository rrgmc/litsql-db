package lsql

import (
	"context"
	"database/sql"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

// Querier is something that sqlscan can query and get the *sql.Rows from.
// For example, it can be: *sql.DB, *sql.Conn or *sql.Tx.
type Querier interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type QuerierDB interface {
	Querier
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type QuerierStmt interface {
	QueryContext(ctx context.Context, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, args ...any) *sql.Row
	ExecContext(ctx context.Context, args ...any) (sql.Result, error)
}

type QuerierTx interface {
	Querier
	Commit() error
	Rollback() error
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
}

type DBQuerier[T Querier] struct {
	queryHandler sq.Handler
	querier      T
}

func NewDBQuerier[T Querier](querier T, options ...Option) *DBQuerier[T] {
	var optns dbOptions
	for _, opt := range options {
		opt(&optns)
	}

	if optns.queryHandler == nil {
		optns.queryHandler = sq.NewHandler()
	}

	return &DBQuerier[T]{
		querier:      querier,
		queryHandler: optns.queryHandler,
	}
}

func (d *DBQuerier[T]) Handler() T {
	return d.querier
}

func (d *DBQuerier[T]) Query(ctx context.Context, query litsql.Query, params any) (*sql.Rows, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return nil, err
	}
	return d.querier.QueryContext(ctx, qstr, args...)
}

func (d *DBQuerier[T]) QueryRow(ctx context.Context, query litsql.Query, params any) (*sql.Row, error) {
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

func (d *DBQuerier[T]) Exec(ctx context.Context, query litsql.Query, params any) (sql.Result, error) {
	qstr, args, err := d.buildQuery(query, params)
	if err != nil {
		return nil, err
	}
	return d.querier.ExecContext(ctx, qstr, args...)
}

func (d *DBQuerier[T]) Prepare(ctx context.Context, query litsql.Query) (*Stmt[*sql.Stmt], error) {
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

func (d *DBQuerier[T]) buildQuery(query litsql.Query, params any) (string, []any, error) {
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
