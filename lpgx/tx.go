package lpgx

import (
	"github.com/jackc/pgx/v5"
	"github.com/rrgmc/litsql-db/lpgx/lpgxt"
)

// Tx wraps a [pgx.Tx].
type Tx = lpgxt.TxT[pgx.Tx]

var _ QuerierTx = (*Tx)(nil)

// NewTx wraps a [pgx.Tx].
func NewTx(tx pgx.Tx, options ...Option) *Tx {
	return lpgxt.NewTxT(tx, options...)
}
