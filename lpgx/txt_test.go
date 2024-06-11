package lpgx

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/sq"
	"gotest.tools/v3/assert"
)

func TestNewTx(t *testing.T) {
	ctx := context.Background()

	dbMock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer dbMock.Close()

	dbMock.ExpectBegin()
	dbMock.ExpectQuery(`SELECT (.+) FROM film WHERE length > (.+) LIMIT (.+)`).
		WithArgs(90, 10).
		WillReturnRows(pgxmock.
			NewRows([]string{"film_id", "title", "length"}).
			AddRow(1, "Test Film", 90))
	dbMock.ExpectCommit()

	ddb := NewDBT[pgx.Tx](dbMock)

	query := psql.Select(
		sm.Columns("film_id", "title", "length"),
		sm.From("film"),
		sm.WhereClause("length > ?", sq.NamedArg("length")),
		sm.Limit(10),
	)

	dtx, err := ddb.BeginTx(ctx, pgx.TxOptions{})
	assert.NilError(t, err)

	rows, err := dtx.Query(ctx, query, map[string]any{
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
