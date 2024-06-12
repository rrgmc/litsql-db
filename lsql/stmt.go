package lsql

import (
	"database/sql"

	"github.com/rrgmc/litsql-db/lsql/lsqlt"
)

// Stmt wraps a [sql.Stmt].
type Stmt = lsqlt.StmtT[*sql.Stmt]

var _ QuerierStmt = (*Stmt)(nil)

// NewStmt wraps a [sql.Stmt].
func NewStmt(stmt *sql.Stmt, args []any, options ...Option) *Stmt {
	return lsqlt.NewStmtT[*sql.Stmt](stmt, args, options...)
}
