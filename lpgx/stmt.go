package lpgx

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rrgmc/litsql-db/lpgx/lpgxt"
)

type StmtConn = lpgxt.Stmt[*pgx.Conn]

type StmtTx = lpgxt.Stmt[pgx.Tx]

type StmtPoolTx = lpgxt.Stmt[*pgxpool.Tx]

var (
	_ QuerierStmt = (*StmtConn)(nil)
	_ QuerierStmt = (*StmtTx)(nil)
	_ QuerierStmt = (*StmtPoolTx)(nil)
)
