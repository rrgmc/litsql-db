package lpgx

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rrgmc/litsql-db/lpgx/lpgxt"
)

type Querier = lpgxt.Querier

type QuerierPool = lpgxt.QuerierPool

type QuerierPoolConn = lpgxt.QuerierPoolConn

type QuerierConn = lpgxt.QuerierConn[*pgx.Conn]

type QuerierStmt = lpgxt.QuerierStmt

type QuerierTx = lpgxt.QuerierTx[pgx.Tx]

type QuerierPoolTx = lpgxt.QuerierTx[*pgxpool.Tx]
