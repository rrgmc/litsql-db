package lpgx

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Conn wraps a [pgx.Conn].
type Conn = ConnT[*pgx.Conn]

// PoolConn wraps a [pgxpool.Conn].
type PoolConn = PoolConnT[*pgxpool.Conn]

var (
	_ QuerierConn     = (*Conn)(nil)
	_ QuerierPoolConn = (*Conn)(nil)
)

// NewConn wraps a [pgx.Conn].
func NewConn(conn *pgx.Conn, options ...Option) *Conn {
	return NewConnT(conn, options...)
}

// NewPoolConn wraps a [pgxpool.Conn].
func NewPoolConn(conn *pgxpool.Conn, options ...Option) *PoolConn {
	return NewPoolConnT(conn, options...)
}
