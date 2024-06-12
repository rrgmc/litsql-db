package lsql

import "database/sql"

type Tx = TxT[*sql.Tx]

var _ QuerierTx = (*Tx)(nil)

// NewTx wraps an [sql.Tx].
func NewTx(tx *sql.Tx, options ...Option) *Tx {
	return NewTxT[*sql.Tx](tx, options...)
}
