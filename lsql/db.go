package lsql

import "database/sql"

type DB = DBT[*sql.DB]

var _ QuerierDB = (*DB)(nil)

// NewDB wraps an [sql.DB].
func NewDB(db *sql.DB, options ...Option) *DB {
	return NewDBT(db, options...)
}
