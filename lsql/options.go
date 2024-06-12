package lsql

import (
	"github.com/rrgmc/litsql-db/lsql/lsqlt"
	"github.com/rrgmc/litsql/sq"
)

type Option = lsqlt.Option

func WithQueryHandler(queryHandler sq.Handler) Option {
	return lsqlt.WithQueryHandler(queryHandler)
}
