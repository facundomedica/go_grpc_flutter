package database

import (
	"database/sql"
	"log"
)

// This file is almost a copy of https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go

// DB is our database. Just explaining this because of the linter!
var DB *sql.DB

// InitDB creates the db file and the tables.
func InitDB() {
	// os.Remove("./foo.db") // uncomment this if you want to clean your database

	var err error
	DB, err = sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	create table users (username text not null primary key, password text);
	create table tasks (id text not null primary key, title text, timestamp integer, completed integer);
	` // completed is an integer because we don't have booleans in SQLite. No problem!

	// execute our query
	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
