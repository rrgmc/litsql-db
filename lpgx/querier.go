package lpgx

import (
	"github.com/jackc/pgx/v5"
)

// type Querier = QuerierT[*sql.Stmt]
// type Querier = QuerierT[*sql.Stmt]

type QuerierDB = QuerierDBT[*pgx.Conn, pgx.Tx]

type QuerierStmt = QuerierStmtT

type QuerierTx = QuerierTxT[pgx.Tx]
