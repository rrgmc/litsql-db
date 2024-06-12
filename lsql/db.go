package lsql

import (
	"database/sql"

	"github.com/rrgmc/litsql-db/lsql/lsqlt"
)

// DB wraps a [sql.DB].
type DB = lsqlt.DB[*sql.DB]

var _ QuerierDB = (*DB)(nil)

// NewDB wraps a [sql.DB].
func NewDB(db *sql.DB, options ...Option) *DB {
	return lsqlt.NewDB(db, options...)
}
