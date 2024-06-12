package lpgx

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rrgmc/litsql-db/lpgx/lpgxt"
)

// Pool wraps a [pgxpool.Pool].
type Pool = lpgxt.Pool[*pgxpool.Pool]

var _ QuerierPool = (*Pool)(nil)

// NewPool wraps a [pgxpool.Pool].
func NewPool(conn *pgxpool.Pool, options ...Option) *Pool {
	return lpgxt.NewPool(conn, options...)
}
