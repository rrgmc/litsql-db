package lsql

import "database/sql"

// Stmt wraps a [sql.Stmt].
type Stmt = StmtT[*sql.Stmt]

var _ QuerierStmt = (*Stmt)(nil)

// NewStmt wraps a [sql.Stmt].
func NewStmt(stmt *sql.Stmt, args []any, options ...Option) *Stmt {
	return NewStmtT[*sql.Stmt](stmt, args, options...)
}
