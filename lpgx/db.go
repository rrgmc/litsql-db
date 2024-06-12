package lpgx

import (
	"github.com/jackc/pgx/v5"
)

type DB = DBT[*pgx.Conn]

var _ QuerierDB = (*DB)(nil)

// NewDB wraps a [pgx.Conn].
func NewDB(db *pgx.Conn, options ...Option) *DB {
	return NewDBT(db, options...)
}
