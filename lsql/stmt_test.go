package lsql

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/sq"
	"gotest.tools/v3/assert"
)

func TestNewDBQuerierStmt(t *testing.T) {
	ctx := context.Background()

	db, dbMock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	dbMock.ExpectPrepare(`SELECT (.+) FROM film WHERE length > (.+) LIMIT (.+)`).
		ExpectQuery().
		WithArgs(90, 10).
		WillReturnRows(sqlmock.
			NewRows([]string{"film_id", "title", "length"}).
			AddRow(1, "Test Film", 90))

	ddb := NewDBQuerierDB(db)

	query := psql.Select(
		sm.Columns("film_id", "title", "length"),
		sm.From("film"),
		sm.WhereC("length > ?", sq.NamedArg("length")),
		sm.Limit(10),
	)

	dstmt, err := ddb.Prepare(ctx, query)
	assert.NilError(t, err)

	rows, err := dstmt.Query(ctx, map[string]any{
		"length": 90,
	})
	assert.NilError(t, err)
	defer rows.Close()

	for rows.Next() {
		var film_id, length int
		var title string
		err = rows.Scan(&film_id, &title, &length)
		assert.NilError(t, err)
	}

	assert.NilError(t, rows.Err())
}
