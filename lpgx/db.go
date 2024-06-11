package lpgx

import (
	"github.com/jackc/pgx/v5"
)

type DB = DBT[*pgx.Conn]

var _ QuerierDB = (*DB)(nil)

func NewDB(db *pgx.Conn, options ...Option) *DB {
	return NewDBT[*pgx.Conn](db, options...)
}