/*
This mod connects to SQLite database

Karl Hunter 2023
2023-11-13
https://www.karlhunter.co.uk/go

*/

package main

import (
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
    "time"
)



// If database does not exist create new
func create_database(path string) int {
	db, err := sql.Open("sqlite3", path + "datafile.db")

    if err != nil {
        log.Fatal(err)
        return 1
    }

    defer db.Close()

	// Write to database
    sts := `
        CREATE TABLE IF NOT EXISTS 
        objects(ID_object INTEGER NOT NULL, 
        path TEXT,
        hash TEXT,
        ts TEXT,
        enabled INTEGER DEFAULT 1,
        PRIMARY KEY(ID_object AUTOINCREMENT));
        CREATE TABLE IF NOT EXISTS 
        about(program TEXT,
        website TEXT,
        version	TEXT,
        author TEXT);
        INSERT INTO about("program","website","version","author")
        VALUES ("KHBackup","http://karlhunter.co.uk/khb","0.1","Karl Hunter");
        `
    _, err = db.Exec(sts)

    if err != nil {
        log.Fatal(err)
        return 1
    }

    return 0
}

func write_files_sql(path string, hashmap map[string]string) int {

    db, err := sql.Open("sqlite3", path + "datafile.db")

    if err != nil {
        log.Fatal("error: at db open, msg: ", err)
        return 1
    }

    defer db.Close()

    // Loop through the hash map and insert into the
    // object table
    for key, value := range hashmap {

        // Timestamp
        now := time.Now()
        timeStr := now.String()

        // Build the SQL string
        sts := "INSERT INTO objects('ID_object','path','hash','ts','enabled') VALUES (NULL,'" + key + "','" + value + "','" + timeStr + "',1);"
        

        // Error reporting
        _, err = db.Exec(sts)
        if err != nil {
            log.Fatal("error: at insertion, msg: ", err)
            return 1
        }

    }

    return 0

}


// Function to check if the file exists in the database
func check_file_sql(short_path string) int {

    // Take the short path to query the database

    // Return value back to called funcion
    return 1

}
