/*
This mod connects to SQLite database

Karl Hunter 2023
2023-12-12
https://www.karlhunter.co.uk/defiant

*/

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/TwiN/go-color"
	_ "github.com/mattn/go-sqlite3"
)

// If database does not exist create new
func create_database(path string, dbfile string) int {

	var db_path string

	// User-defined db location
	if dbfile != " " {
		db_path = filepath.Clean(dbfile)
	} else {
		db_path = filepath.Join(path, db_name)
	}

	db, err := sql.Open("sqlite3", db_path)

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
        VALUES ("DEFIANTFG","http://karlhunter.co.uk/defiant","1","Karl Hunter");
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
func write_files_sql(path string, hashmap map[string]string, noinfo bool, dbfile string) int {

	db, err := sql.Open("sqlite3", filepath.Join(dbfile))

	if err != nil {
		log.Fatal("error: Could not open db, msg: ", err)
		return 1
	}

	defer db.Close()

	// Count the map
	new_files_count := len(hashmap)

	// Warn user; if more than 10 warn this may take a while
	if new_files_count > 0 {
		if noinfo == false {
			log.Println("info: writing block of hashes to database", new_files_count)
		}
	}

	// Loop through the hash map and insert into the
	// object table

	for key, value := range hashmap {

		// Remove the apostrophe if in file name
		// this will prevent SQL error on insertion to db
		clean_key := clean_string(key, 1)

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

	/*if new_files_count > 0 {
	    log.Println("info: database writes completed")
	}*/

	return 0

}

// Function to check if the file exists in the database
// This function is called from the iterate function found in
// the files.go package
func check_file_sql(short_path string, full_path string, hash string, dbfile string, logon bool) int {

	// Check if first record, the given diectory
	if short_path == "" {
		return 0
	}

	// Open database
	db, err := sql.Open("sqlite3", dbfile)

	if err != nil {
		log.Fatal("fatal: at db open, msg: ", err)
	}

	defer db.Close()

	// Writer file

	// Clean variable
	short_path = clean_string(short_path, 1)

	rows, err := db.Query("SELECT path, hash FROM objects WHERE objects.path = ? ORDER BY objects.ID_object DESC LIMIT 1;", short_path)

	if err != nil {
		log.Fatal("fatal: db query error: ", err)
	}

	defer rows.Close()

	for rows.Next() {
		var s_path string
		var s_hash string
		err = rows.Scan(&s_path, &s_hash)
		if err != nil {
			log.Println("fatal: SQL query error", err)
		}

		// Check the hash against the OS path and DB path
		if s_hash != hash {
			if logon == true {
				log.Println("fail: " + full_path + short_path)
			} else {
				log.Println(color.Red + "fail: " + full_path + short_path + color.Reset)
			}
		}

		return 0

	}

	return 1

}

// Query database and check files are still present
func missing_files_scan(full_path string, dbfile string, logon bool) int {

	// Iterate through directory using db records
	// if exists, skip. If record in db but not file
	// If in db but not in directory structure, warn user

	// Open database
	db, err := sql.Open("sqlite3", dbfile)

	if err != nil {
		log.Fatal("fatal: at db open, msg: ", err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT path FROM objects WHERE objects.enabled = 1 ORDER BY objects.ID_object DESC;")

	if err != nil {
		log.Fatal("fatal: db query error: ", err)
	}

	defer rows.Close()

	for rows.Next() {
		var s_path string
		err = rows.Scan(&s_path)

		// Add back special characters
		s_path = clean_string(s_path, 0)
		// System path, path and s_path combined
		sys_path := full_path + s_path

		// Check if the file exists
		// If not returns 0, the warn user
		if is_file(sys_path) != 1 {
			// Remove colour if log file
			if logon == true {
				log.Println("missing: " + s_path)
			} else {
				log.Println(color.Yellow + "missing: " + s_path + color.Reset)
			}
		}

		if err != nil {
			log.Println("fatal: SQL query error", err)
		}
	}

	return 0

}

func clean_string(filename string, do int) string {

	// To prevent SQL query errors, need to remove apostrophe
	// and <!. This function replaces them with xAPOSx so that
	// the apostrope can be added back to check later.

	if do == 1 {
		// Remove file characters
		filename = strings.Replace(filename, "'", "xAPOSx", -1)
		filename = strings.Replace(filename, "<!", "xBRACKx", -1)
	} else if do == 0 {
		filename = strings.Replace(filename, "xAPOSx", "'", -1)
		filename = strings.Replace(filename, "xBRACKx", "<!", -1)
	} else {
		// To be safe clean strings
		filename = strings.Replace(filename, "'", "xAPOSx", -1)
		filename = strings.Replace(filename, "<!", "xBRACKx", -1)
	}

	return filename

}

func disable_sql_func(dbfile string, db_ID int) {

	var db_path string
	db_path = filepath.Clean(dbfile)

	db, err := sql.Open("sqlite3", db_path)

	if err != nil {
		log.Fatal("fatal: could not connect to db", err)
	}

	defer db.Close()

	// Display file name to confirm with user if correct:
	rows, err := db.Query("SELECT path FROM main.objects WHERE main.objects.ID_object=?", db_ID)
	defer rows.Close()

	if err != nil {
		log.Fatal("Error: unable to locate ID; is it correct?")
		os.Exit(1)
	}

	for rows.Next() {
		var path string
		var choice string
		rows.Scan(&path)

		fmt.Println("File to delete:", path)
		fmt.Println("Correct? [y/n]")
		fmt.Scan(&choice)

		switch choice {
		case "y":
			fmt.Println("deleted!")
		default:
			fmt.Println("Delete operation was cancelled")
			os.Exit(1)
		}

	}

	/*
			// Write change to database
			sts := `
			UPDATE main.objects SET enabled = 0 WHERE main.objects.ID_object = 2
		        `
			_, err = db.Exec(sts)

			if err != nil {
				log.Fatal(err)
				return 1
			}
	*/
	//return 0

}
