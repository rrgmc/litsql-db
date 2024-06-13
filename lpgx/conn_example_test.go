package lpgx_test

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rrgmc/litsql-db/lpgx"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/sq"
)

func ExampleConn() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "test")
	if err != nil {
		panic(err)
	}

	// wrap *pgx.Conn instance
	dconn := lpgx.NewConn(conn)

	query := psql.Select(
		sm.Columns("film_id", "title", "length"),
		sm.From("film"),
		sm.WhereClause("length > ?", sq.NamedArg("length")),
		sm.Limit(10),
	)

	// generate SQL string from litsql and execute it, replacing named parameters.
	rows, err := dconn.Query(ctx, query, sq.MapArgValues{
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
