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
    "strings"
    "unicode/utf8"
)



// If database does not exist create new
func create_database(path string) int {
	db, err := sql.Open("sqlite3", path + "datafile.db")

    if err != nil {
        log.Fatal("fatal: could not connect to db", err)
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

// This function will take the map (array of all the filenames and hashes)
// and write to the SQLite3 database file
func write_files_sql(path string, hashmap map[string]string) int {

    db, err := sql.Open("sqlite3", path + "datafile.db")

    if err != nil {
        log.Fatal("error: at db open, msg: ", err)
        return 1
    }

    defer db.Close()

    // Count the map
    new_files_count := len(hashmap)

    // Warn user; if more than 10 warn this may take a while
    if new_files_count > 10 {
        log.Println("info: writing new files to database", new_files_count, "(this may take a while)")
    } else {
        log.Println("info: writing new files to database", new_files_count)
    }

    // Loop through the hash map and insert into the
    // object table
    for key, value := range hashmap {

        // Remove the apostrophe if in file name
        // this will prevent SQL error on insertion to db
        clean_key := clean_string(key)

        // Timestamp
        now := time.Now()
        timeStr := now.Format("20060102150405")

        // Build the SQL string
        sts := "INSERT INTO objects('ID_object','path','hash','ts','enabled') VALUES (NULL,'" + clean_key + "','" + value + "','" + timeStr + "',1);"
        

        // Error reporting
        _, err = db.Exec(sts)
        if err != nil {
            log.Println("error: cannot insert to data file, msg: ", err, "file:", key)
        }

    }

    log.Println("info: finished writing to database")

    return 0

}


// Function to check if the file exists in the database
// This function is called from the iterate function found in 
// the files.go package
func check_file_sql(short_path string, full_path string) int {

    var output string

    // Check if first record, the given diectory
    if short_path == "" {
        return 0
    }


    // Get user input path by removing short_name
    // from the long path
    file_name_count := utf8.RuneCountInString(short_path)
    db_path := full_path[:file_name_count]


    // Open database
    db, err := sql.Open("sqlite3", db_path + "datafile.db")

        if err != nil {
            log.Fatal("fatal: at db open, msg: ", err)
        }

        
    // Clean variable
    short_path = clean_string(short_path)
            
    // Prepare the query here
    // Searching for the short path against the path sent through function
    // In case of duplications, the query will select the first record, 
    // just in case of known hash changes
    query, err := db.Prepare("SELECT * FROM objects WHERE path = '" + short_path + "' ORDER BY ID_object DESC LIMIT 1;")

    // Look for SQL errors
    // Otherwise, query later will not work
    // Return 0 breaks and next file tried
    if err != nil {
        log.Println("error: could not query for file", err)
        return 0
    }

   
    defer query.Close()
       
    // Execute query
    err = query.QueryRow(short_path).Scan(&output)

    // Catch errors from query
    switch {
        case err == sql.ErrNoRows:
            // No record of the file in the db
            // this means that during iteration the file
            // found must be a new file
            log.Println("info: new file:", short_path)
            return 1
        case err != nil:
            log.Println("%s", err)
            return 0
        default:
            log.Println("Counted %s \n", output)
            return 0
    }
    
	return 0

}

// Query database and check files are still present
func missing_files_scan(full_path string, short_path string) int {

    return 0

}

// Clean string
func clean_string(filename string) string {

    filename = strings.Replace(filename, "'", "", -1)
    filename = strings.Replace(filename, "<!", "", -1)

    return filename
    
}