package lpgx_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/rrgmc/litsql-db/lpgx/lpgxt"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/sq"
	"gotest.tools/v3/assert"
)

func TestNewStmt(t *testing.T) {
	ctx := context.Background()

	dbMock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer dbMock.Close()

	sname := "test1"

	dbMock.ExpectPrepare(sname, `SELECT (.+) FROM film WHERE length > (.+) LIMIT (.+)`)
	dbMock.ExpectQuery(sname).
		WithArgs(90, 10).
		WillReturnRows(dbMock.
			NewRows([]string{"film_id", "title", "length"}).
			AddRow(1, "Test Film", 90))

	dconn := lpgxt.NewConn(dbMock)

	query := psql.Select(
		sm.Columns("film_id", "title", "length"),
		sm.From("film"),
		sm.WhereClause("length > ?", sq.NamedArg("length")),
		sm.Limit(10),
	)

	dstmt, err := dconn.Prepare(ctx, sname, query)
	assert.NilError(t, err)

	rows, err := dstmt.Query(ctx, sq.MapArgValues{
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

func TestNewStmtTx(t *testing.T) {
	ctx := context.Background()

	dbMock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer dbMock.Close()

	sname := "test1"

	dbMock.MatchExpectationsInOrder(false)

	dbMock.ExpectBegin()
	dbMock.ExpectPrepare(sname, `SELECT (.+) FROM film WHERE length > (.+) LIMIT (.+)`)
	dbMock.ExpectQuery(sname).
		WithArgs(90, 10).
		WillReturnRows(pgxmock.
			NewRows([]string{"film_id", "title", "length"}).
			AddRow(1, "Test Film", 90))
	dbMock.ExpectCommit()

	dconn := lpgxt.NewConn(dbMock)

	query := psql.Select(
		sm.Columns("film_id", "title", "length"),
		sm.From("film"),
		sm.WhereClause("length > ?", sq.NamedArg("length")),
		sm.Limit(10),
	)

	dstmt, err := dconn.Prepare(ctx, sname, query)
	assert.NilError(t, err)

	dtx, err := dconn.BeginTx(ctx, pgx.TxOptions{})
	assert.NilError(t, err)

	rows, err := dstmt.Query(ctx, sq.MapArgValues{
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

	err = dtx.Commit(ctx)
	assert.NilError(t, err)
}
