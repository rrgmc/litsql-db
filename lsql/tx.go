package lsql

import "database/sql"

// Tx wraps a [sql.Tx].
type Tx = TxT[*sql.Tx]

var _ QuerierTx = (*Tx)(nil)

// NewTx wraps a [sql.Tx].
func NewTx(tx *sql.Tx, options ...Option) *Tx {
	return NewTxT[*sql.Tx](tx, options...)
}
