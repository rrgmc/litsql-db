package lpgx

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rrgmc/litsql-db/lpgx/lpgxt"
)

type QuerierConn = lpgxt.QuerierConn[*pgx.Conn]

type QuerierPoolConn = lpgxt.QuerierPoolConn[*pgxpool.Conn]

type QuerierPool = lpgxt.QuerierPool[*pgxpool.Pool]

type QuerierStmt = lpgxt.QuerierStmt

type QuerierTx = lpgxt.QuerierTx[pgx.Tx]

type QuerierPoolTx = lpgxt.QuerierTx[*pgxpool.Tx]
