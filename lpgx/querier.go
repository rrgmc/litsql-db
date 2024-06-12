package lpgx

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QuerierConn = QuerierConnT[*pgx.Conn]

type QuerierPool = QuerierPoolT[*pgxpool.Pool]

type QuerierStmt = QuerierStmtT

type QuerierTx = QuerierTxT[pgx.Tx]
