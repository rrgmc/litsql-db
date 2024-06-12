package lpgx

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QuerierConn = QuerierConnT[*pgx.Conn]

type QuerierPoolConn = QuerierPoolConnT[*pgxpool.Conn]

type QuerierPool = QuerierPoolT[*pgxpool.Pool]

type QuerierStmt = QuerierStmtT

type QuerierTx = QuerierTxT[pgx.Tx]

type QuerierPoolTx = QuerierTxT[*pgxpool.Tx]
