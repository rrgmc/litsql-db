package lpgx

import (
	"github.com/jackc/pgx/v5"
)

// Conn wraps a [pgx.Conn].
type Conn = ConnT[*pgx.Conn]

var _ QuerierConn = (*Conn)(nil)

// NewConn wraps a [pgx.Conn].
func NewConn(conn *pgx.Conn, options ...Option) *Conn {
	return NewConnT(conn, options...)
}
