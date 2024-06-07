package lsql

import "database/sql"

type Stmt = StmtT[*sql.Stmt]

var _ QuerierStmt = (*Stmt)(nil)

func NewStmt(stmt *sql.Stmt, args []any, options ...Option) *Stmt {
	return NewStmtT[*sql.Stmt](stmt, args, options...)
}
