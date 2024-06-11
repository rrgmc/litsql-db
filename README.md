# litsql-db

DB wrappers for [litsql](https://github.com/rrgmc/litsql).

 * [stdlib sql](lsql/)
 * [pgx](lpgx/)

## Example

#### Query

```go
func ExampleDB() {
    ctx := context.Background()

    db, err := sql.Open("test", ":memory:")
    if err != nil {
        panic(err)
    }

    // wrap *sql.DB instance
    ddb := lsql.NewDB(db)

    query := psql.Select(
        sm.Columns("film_id", "title", "length"),
        sm.From("film"),
        sm.WhereClause("length > ?", sq.NamedArg("length")),
        sm.Limit(10),
    )

    // generate SQL string from litsql and execute it, replacing named parameters.
    rows, err := ddb.Query(ctx, query, map[string]any{
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

    db, err := sql.Open("test", ":memory:")
    if err != nil {
        panic(err)
    }

    // wrap *sql.DB instance
    ddb := lsql.NewDB(db)

    query := psql.Select(
        sm.Columns("film_id", "title", "length"),
        sm.From("film"),
        sm.WhereClause("length > ?", sq.NamedArg("length")),
        sm.Limit(10),
    )

    // generate SQL string from litsql and prepare it, storing named parameter to be replaced later
    dstmt, err := ddb.Prepare(ctx, query)
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
