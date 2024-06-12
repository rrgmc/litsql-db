package lpgx

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rrgmc/litsql-db/lpgx/lpgxt"
)

// Conn wraps a [pgx.Conn].
type Conn = lpgxt.Conn[*pgx.Conn]

// PoolConn wraps a [pgxpool.Conn].
type PoolConn = lpgxt.PoolConn[*pgxpool.Conn]

var (
	_ QuerierConn     = (*Conn)(nil)
	_ QuerierPoolConn = (*Conn)(nil)
)

// NewConn wraps a [pgx.Conn].
func NewConn(conn *pgx.Conn, options ...Option) *Conn {
	return lpgxt.NewConn(conn, options...)
}

// NewPoolConn wraps a [pgxpool.Conn].
func NewPoolConn(conn *pgxpool.Conn, options ...Option) *PoolConn {
	return lpgxt.NewPoolConn(conn, options...)
}
