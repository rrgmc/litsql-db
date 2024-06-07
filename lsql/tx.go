package lsql

import "database/sql"

type Tx = TxT[*sql.Tx]

func NewTx(tx *sql.Tx, options ...Option) *Tx {
	return NewTxT[*sql.Tx](tx, options...)
}
