package lsql

import (
	"database/sql"

	"github.com/rrgmc/litsql-db/lsql/lsqlt"
)

// Tx wraps a [sql.Tx].
type Tx = lsqlt.TxT[*sql.Tx]

var _ QuerierTx = (*Tx)(nil)

// NewTx wraps a [sql.Tx].
func NewTx(tx *sql.Tx, options ...Option) *Tx {
	return lsqlt.NewTxT[*sql.Tx](tx, options...)
}
