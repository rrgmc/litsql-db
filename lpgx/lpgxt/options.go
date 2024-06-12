package lpgxt

import "github.com/rrgmc/litsql/sq"

type Option func(options *dbOptions)

type dbOptions struct {
	queryHandler sq.Handler
}

func WithQueryHandler(queryHandler sq.Handler) Option {
	return func(options *dbOptions) {
		options.queryHandler = queryHandler
	}
}
