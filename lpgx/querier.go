package lpgx

import (
	"github.com/jackc/pgx/v5"
)

type QuerierDB = QuerierDBT[*pgx.Conn]

type QuerierStmt = QuerierStmtT

type QuerierTx = QuerierTxT[pgx.Tx]
