package lsql

import "database/sql"

type QuerierDB = QuerierDBT[*sql.DB]

type QuerierStmt = QuerierStmtT

type QuerierTx = QuerierTxT[*sql.Tx]
