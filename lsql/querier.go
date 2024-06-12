package lsql

import (
	"database/sql"

	"github.com/rrgmc/litsql-db/lsql/lsqlt"
)

type QuerierDB = lsqlt.QuerierDBT[*sql.DB]

type QuerierStmt = lsqlt.QuerierStmtT

type QuerierTx = lsqlt.QuerierTxT[*sql.Tx]
