package lpgx

import (
	"github.com/jackc/pgx/v5"
)

// Tx wraps a [pgx.Tx].
type Tx = TxT[pgx.Tx]

var _ QuerierTx = (*Tx)(nil)

// NewTx wraps a [pgx.Tx].
func NewTx(tx pgx.Tx, options ...Option) *Tx {
	return NewTxT(tx, options...)
}
