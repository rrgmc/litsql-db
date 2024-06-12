package lsql

import "database/sql"

// DB wraps a [sql.DB].
type DB = DBT[*sql.DB]

var _ QuerierDB = (*DB)(nil)

// NewDB wraps a [sql.DB].
func NewDB(db *sql.DB, options ...Option) *DB {
	return NewDBT(db, options...)
}
