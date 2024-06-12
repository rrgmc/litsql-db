package lpgx

import (
	"github.com/rrgmc/litsql-db/lpgx/lpgxt"
	"github.com/rrgmc/litsql/sq"
)

type Option = lpgxt.Option

func WithQueryHandler(queryHandler sq.Handler) Option {
	return lpgxt.WithQueryHandler(queryHandler)
}
