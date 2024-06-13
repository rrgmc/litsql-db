package lsql_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql-db/lsql"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/sq"
	"gotest.tools/v3/assert"
)

func TestNewDB(t *testing.T) {
	ctx := context.Background()

	db, dbMock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	dbMock.ExpectQuery(`SELECT (.+) FROM film WHERE length > (.+) LIMIT (.+)`).
		WithArgs(90, 10).
		WillReturnRows(sqlmock.
			NewRows([]string{"film_id", "title", "length"}).
			AddRow(1, "Test Film", 90))

	ddb := lsql.NewDB(db)

	query := psql.Select(
		sm.Columns("film_id", "title", "length"),
		sm.From("film"),
		sm.WhereClause("length > ?", sq.NamedArg("length")),
		sm.Limit(10),
	)

	rows, err := ddb.Query(ctx, query, litsql.MapArgValues{
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

func TestNewDBQueryHandler(t *testing.T) {
	ctx := context.Background()

	db, dbMock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	dbMock.ExpectQuery(`SELECT (.+) FROM film WHERE length > (.+) LIMIT (.+)`).
		WithArgs(90, 10).
		WillReturnRows(sqlmock.
			NewRows([]string{"film_id", "title", "length"}).
			AddRow(1, "Test Film", 90))

	ddb := lsql.NewDB(db, lsql.WithQueryHandler(sq.NewHandler(
		sq.WithDefaultBuildOptions(
			sq.WithWriterOptions(sq.WithUseNewLine(false)),
		),
	)))

	query := psql.Select(
		sm.Columns("film_id", "title", "length"),
		sm.From("film"),
		sm.WhereClause("length > ?", sq.NamedArg("length")),
		sm.Limit(10),
	)

	rows, err := ddb.Query(ctx, query, litsql.MapArgValues{
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
