package lpgx

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool wraps a [pgxpool.Pool].
type Pool = PoolT[*pgxpool.Pool]

var _ QuerierPool = (*Pool)(nil)

// NewPool wraps a [pgxpool.Pool].
func NewPool(conn *pgxpool.Pool, options ...Option) *Pool {
	return NewPoolT(conn, options...)
}
