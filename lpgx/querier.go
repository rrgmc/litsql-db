package lpgx

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rrgmc/litsql-db/lpgx/lpgxt"
)

type QuerierConn = lpgxt.QuerierConnT[*pgx.Conn]

type QuerierPoolConn = lpgxt.QuerierPoolConnT[*pgxpool.Conn]

type QuerierPool = lpgxt.QuerierPoolT[*pgxpool.Pool]

type QuerierStmt = lpgxt.QuerierStmtT

type QuerierTx = lpgxt.QuerierTxT[pgx.Tx]

type QuerierPoolTx = lpgxt.QuerierTxT[*pgxpool.Tx]
