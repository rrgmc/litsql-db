package lsql

import (
	"database/sql"

	"github.com/rrgmc/litsql-db/lsql/lsqlt"
)

// DB wraps a [sql.DB].
type DB = lsqlt.DBT[*sql.DB]

var _ QuerierDB = (*DB)(nil)

// NewDB wraps a [sql.DB].
func NewDB(db *sql.DB, options ...Option) *DB {
	return lsqlt.NewDBT(db, options...)
}
