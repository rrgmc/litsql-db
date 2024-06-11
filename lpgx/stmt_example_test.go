package lpgx_test

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rrgmc/litsql-db/lpgx"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/sq"
)

func ExampleStmt() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "test")
	if err != nil {
		panic(err)
	}

	// wrap *pgx.Conn instance
	ddb := lpgx.NewDB(conn)

	query := psql.Select(
		sm.Columns("film_id", "title", "length"),
		sm.From("film"),
		sm.WhereClause("length > ?", sq.NamedArg("length")),
		sm.Limit(10),
	)

	queryName := "query1"

	// generate SQL string from litsql and prepare it, storing the named parameters to be replaced later
	dstmt, err := ddb.Prepare(ctx, queryName, query)
	if err != nil {
		panic(err)
	}

	// execute prepared query, replacing named parameters
	rows, err := dstmt.Query(ctx, map[string]any{
		"length": 90,
	})
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var film_id, length int
		var title string
		err = rows.Scan(&film_id, &title, &length)
		if err != nil {
			panic(err)
		}
	}

	if rows.Err() != nil {
		panic(err)
	}
}
