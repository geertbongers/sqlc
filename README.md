sqlc
----

[![Build Status](https://travis-ci.org/relops/sqlc.png?branch=master)](https://travis-ci.org/relops/sqlc)
[![GoDoc](http://godoc.org/_?status.png)](http://godoc.org/github.com/relops/sqlc)

`sqlc` is a composable, type safe and fluent API to generate nested and complex SQL queries.

Taking heavy inspiration from [JOOQ][], `sqlc` generates SQL queries for you:
	
```go	
var FOO sqlc.TableLike // (optionally) auto-generated by sqlc by introspecting your DDL

var db *db.DB // For integration with database/sql
var d Dialect // Either sqlite, mysql or postgres

row, err := Select(FOO.BAR).From(FOO).Where(FOO.BAZ.Eq("quux")).QueryRow(d, db)
```

If you don't want to use `database/sql`, you don't have to - ultimately `sqlc` is just a string building tool.

`String(Dialect)` is an API call to just produce the SQL string that you use in any way that you want to:

```go
// Renders `SELECT foo.bar FROM foo WHERE foo.baz = ?`
sql := Select(FOO.BAR).From(FOO).Where(FOO.BAZ.Eq("quux")).String(d)
```

Installing
----------

To install the runtime libraries and the `sqlc` command line tool into your `$GOPATH`:

	$ go get github.com/relops/sqlc

Composing Queries
-----------------

You can compose query objects into reusable and individually executable building blocks. For example, you can create a sub query that is in itself executable:

```go
subQuery := Select(
	CALL_RECORDS.REGION,
	CALL_RECORDS.DURATION.Min(),
	CALL_RECORDS.DURATION.Max(),
	CALL_RECORDS.DURATION.Avg()).
	From(CALL_RECORDS).
	GroupBy(CALL_RECORDS.REGION).
	OrderBy(CALL_RECORDS.REGION)

row, err := subQuery.QueryRow(d, db)
```

And then you re-use the subquery as part of a new query:

```go
row, err := SelectCount().From(subQuery).QueryRow(d, db)
```

Type Safety
-----------

`sqlc` provides type safe methods for INSERTs and UPDATEs:

```go
result, err := InsertInto(CALL_RECORDS).
	SetString(CALL_RECORDS.IMSI, "230023741299234").
	SetTime(CALL_RECORDS.TIMESTAMP, time.Now()).
	SetInt(CALL_RECORDS.DURATION, 10).
	SetString(CALL_RECORDS.REGION, "quux").
	SetString(CALL_RECORDS.CALLING_NUMBER, "76581231298").
	SetString(CALL_RECORDS.CALLED_NUMBER, "76754238764").
	Exec(d, db)
```

For example, the following invocation would not compile:

```go
...
SetTime(CALL_RECORDS.TIMESTAMP, "some string"). // Results in a compile time error
...
```

If you use the `sqlc` code generator, you can keep your application in sync with your current DB schema any divergence between your code and the DDL will be flagged by the Go compiler.

INSERTs, UPDATEs, DELETEs
-------------------------

The support for inserting, updating and deleting rows is basic right now:

```go
// Renders `INSERT INTO foo (bar) VALUES (?)` on MySQL
// Renders `INSERT INTO foo (bar) VALUES ($1)` on Postgres
InsertInto(foo).SetString(bar, "quux").String(d)

// Renders `UPDATE foo SET bar = ? WHERE foo.baz = ?"` on MySQL
// Renders `UPDATE foo SET bar = $1 WHERE foo.baz = $2"` on Postgres
Update(foo).SetString(bar, "quux").Where(baz.Eq("gorp")).String(d)

// Renders `DELETE FROM foo WHERE foo.baz = ?`
Delete(foo).Where(baz.Eq("gorp")).String(d)
```

Currently `sqlc` assumes that you want to generate prepared statements and (re)bind application parameters.

Aliasing
--------

`sqlc` allows you to alias your projections easily:

```go
// Renders `SELECT foo.bar AS x, foo.baz AS y FROM foo`
Select(bar.As("x"), baz.As("y")).From(foo).String(d)
```

By default, columns will be qualified by the name of their parent table. You can override this by aliasing the table, in addition to aliasing just the fields:

```go
// Renders `SELECT f.bar AS x, f.baz AS y FROM foo AS f`
Select(bar.As("x"), baz.As("y")).From(foo.As("f")).String(d)
```

Functions
---------

Functions can be applied to any field and they can be nested to any depth:

```go
// Renders `SELECT LOWER(HEX(MD5(foo.bar))) FROM foo`
Select(bar.Md5().Hex().Lower()).From(foo).String(d)
```

Joins
-----

There is basic support for support for joins:

```go
// Renders `SELECT foo.bar, quux.col FROM foo JOIN quux 
//          ON (quux.id = foo.bar AND quux.col = foo.baz)`
Select(bar, col).From(foo).Join(quux).On(id.IsEq(bar), col.IsEq(baz)).String(d)
```

In addition to INNER JOINs, LEFT OUTER JOINs are also supported:

```go
// Renders `SELECT foo.bar FROM foo LEFT OUTER JOIN quux ON quux.id = foo.bar`
Select(bar).From(foo).LeftOuterJoin(quux).
			On(id.IsEq(bar)).String(d)
```

An arbrirary number of joins can be constructed:

```go
// Renders `SELECT foo.bar FROM foo 
//          LEFT OUTER JOIN quux ON quux.id = foo.bar
//          LEFT OUTER JOIN gorp ON gorp.porg = foo.bar`
Select(bar).From(foo).LeftOuterJoin(quux).
			On(id.IsEq(bar)).LeftOuterJoin(gorp).
			On(porg.IsEq(bar)).String(d)
```

Returning From Insert
---------------------

(This is Postgres only feature)

You can specify a column from an INSERT to return back to the app:

```go
// Renders `INSERT INTO foo (bar) VALUES ($1) RETURNING id` on Postgres
InsertInto(foo).SetString(bar, "quux").Returning(id).String(d)
```

`Returning()` returns a fetchable row that you can bind from:

```go
var id int
row, _ := InsertInto(foo).SetString(bar, "quux").Returning(id).Fetch(d, db)
row.Scan(&id)
```

Code Generation
---------------

Install the `sqlc` command line tool:

	$ go get github.com/relops/sqlc

Make sure `sqlc` is on your PATH (usually $GOPATH/bin).

Then point `sqlc` at your sqlite DB file:

	$ sqlc -h
	Usage:
	  sqlc [OPTIONS]

	Application Options:
	  -f, --file=    The path to the sqlite file
      -u, --url=     The DB URL
      -o, --output=  The path to save the generated objects to
  	  -p, --package= The package to put the generated objects into
      -t, --type=    The type of the DB (mysql,postgres,sqlite)
      -s, --schema=  The target DB schema (required for MySQL and Postgres)

	Help Options:
	  -h, --help     Show this help message

Now you can use the generated objects in your app.

Database Support
----------------

* Sqlite
* MySQL
* Postgres

Features
--------

* SELECT ... FROM ... WHERE
* GROUP BY
* ORDER BY (assumes ASC right now)
* INSERTs
* UPDATEs
* DELETEs
* INNER JOINS (integration test only for single columns)
* LEFT OUTER JOINS (unit tested only, no integration test)
* Sub queries
* Data types currently implemented:
  * VARCHAR
  * INT
  * INTEGER
  * TIMESTAMP
* Functions (implemented on an as needs basis, easily extended):
  * COUNT
  * CAST
  * AVG
  * MAX
  * MIN
  * DIV
  * CEIL
  * MD5
  * LOWER
  * HEX
  * GROUP_CONCAT
  * TRUNC
* Statement rendering
* Querying via database/sql
* Code generation of table and field objects from an exising DB schema

Building
--------

To use the `sqlc` tool and runtime libraries, all that is required is a simple `go get`. However, if you want to build `sqlc` from scratch and run the integration tests, you'll need the following installed locally:

* [go-bindata](https://github.com/jteeuwen/go-bindata)
* Postgres
* MySQL

Status
------

Experimental - this is work in progress. Basically I'm trying to port [JOOQ][] to Go, but I don't know yet whether it will work.

[jooq]: http://jooq.org
