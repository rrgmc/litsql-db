package lsql

import (
	"database/sql"

	"github.com/rrgmc/litsql-db/lsql/lsqlt"
)

type QuerierDB = lsqlt.QuerierDB[*sql.DB]

type QuerierStmt = lsqlt.QuerierStmt

type QuerierTx = lsqlt.QuerierTx[*sql.Tx]
