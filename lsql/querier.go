package lsql

import "database/sql"

type Querier = QuerierT[*sql.Stmt]

type QuerierDB = QuerierDBT[*sql.Stmt, *sql.Tx]

type QuerierStmt = QuerierStmtT

type QuerierTx = QuerierTxT[*sql.Stmt]
