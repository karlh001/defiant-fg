package main

import (
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

// If database does not exist create new function
func create_new() {
	db, err := sql.Open("sqlite3", "test.db")

    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

	// Write to database
    sts := `
CREATE TABLE IF NOT EXISTS objects(id INTEGER PRIMARY KEY, path TEXT, hash TEXT);
`
    _, err = db.Exec(sts)

    if err != nil {
        log.Fatal(err)
    }
}

func add_file() {

	
		// need to cycle through array to add these in one
		// go after run through

		
}
