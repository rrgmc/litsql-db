# litsql-db lpgx

[![Test Status](https://github.com/rrgmc/litsql-db/actions/workflows/go.yml/badge.svg)](https://github.com/rrgmc/litsql-db/actions/workflows/go.yml) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/rrgmc/litsql-db) [![Go Reference](https://pkg.go.dev/badge/github.com/rrgmc/litsql-db/lpgx.svg)](https://pkg.go.dev/github.com/rrgmc/litsql-db/lpgx) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/rrgmc/litsql-db)

Golang pgx DB wrappers for [litsql](https://github.com/rrgmc/litsql).

## Installation

```shell
go get -u github.com/rrgmc/litsql-db/lpgx
```

## Examples

#### Query

```go
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
    rows, err := dconn.Query(ctx, query, map[string]any{
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
```

#### Prepared statement

```go
func ExampleStmt() {
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

    queryName := "query1"

    // generate SQL string from litsql and prepare it, storing the named parameters to be replaced later
    dstmt, err := dconn.Prepare(ctx, queryName, query)
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
```

## Author

Rangel Reale (rangelreale@gmail.com)
