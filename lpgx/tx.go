package lpgx

import (
	"github.com/jackc/pgx/v5"
)

type Tx = TxT[pgx.Tx]

var _ QuerierTx = (*Tx)(nil)

func NewTx(tx pgx.Tx, options ...Option) *Tx {
	return NewTxT[pgx.Tx](tx, options...)
}