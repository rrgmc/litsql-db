package lsql

import "database/sql"

type DB = DBT[*sql.Stmt, *sql.DB]

var _ QuerierDB = (*DB)(nil)

func NewDB(db *sql.DB, options ...Option) *DB {
	return NewDBT[*sql.Stmt](db, options...)
}
