/*
This mod checks whether a given string is a
directory.

Returns 1 for directory, 0 for not (likely file)

Karl Hunter 2023
2023-11-19
https://www.karlhunter.co.uk/defiant

*/

package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

func iterate(path string, path_count int, noinfo bool) int {

	// Path is full path to file; e.g. /dir/hello-world.txt
	// short_path is dir after or just file name
	// e.g. hello-world.txt
	// Create a map to store data
	hashmap := map[string]string{}
	full_path := path

	if noinfo == false {
		log.Println("info: scanning files")
	}

	// This function will cycle through the directory and print
	// files and directories.
	// If there was an error, e.g. permissions, then error message
	// output to the user
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		// Check for errors, and if short_path empty skip (liekly root dir)
		if err != nil {
			log.Println("error:", err.Error())
			return nil
		}

		// Feed the path and file through the hash function
		// This will return the hash value
		// hash.go
		file_hash := hash_function(path)

		// Grab the first x characters for path
		//full_path := path[0:path_count]
		// To remove the full path use:
		short_path := path[path_count:]

		// Skip DFG files
		if short_path == "" || short_path == "/"+db_name || short_path == "/"+log_name {
			return nil
		}

		// Check database to see if we have seen this
		// file before; need to use the short path
		// because it's the short path stored in sql
		file_check_result := check_file_sql(short_path, full_path, file_hash)

		// Runs this if statement if
		// is a file, not directory AND
		// file is not already known from db
		// if file_check_result returns 1, mean add to db
		if is_file(path) == 1 && file_check_result == 1 {

			// Add file with hash to the map
			// This will be sent later to insert
			// into the database later
			hashmap[short_path] = file_hash

			// REMOVED DUE TO BUG - UNABLE TO WRITE AND READ
			// AT THE SAME TIME. HELP NEEDED; PURPOSE TO SPEED
			// UP DATABASE WRITES

			// Check size of hashmap
			// if hash map greater than 9 and there
			// is no database lock continue
			// Or skip until the next cycle
			/*if len(hashmap) > 999999 && db_working == 0 {

				// Send hash map to SQL writer
				go write_files_sql(full_path, hashmap)

				// Clear hash map for further files
				hashmap = make(map[string]string)

			}*/
		}

		return nil
	})

	// Send map of new files to insert SQL function
	// Check if any more files to write
	write_files_sql(full_path, hashmap, noinfo)

	return 1

}

func is_file(path string) int {
	var is_file_check int

	// If the directory does not exist then return
	// 0, otherwise 1 means exists
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		is_file_check = 0
	} else if errors.Is(err, os.ErrNotExist) {
		is_file_check = 0
	} else {
		// Is a file
		is_file_check = 1
	}

	return is_file_check
}

func dir_exists(path string) int {
	var dir_exists_outcome int

	// If the directory does not exist then return
	// 0, otherwise 1 means exists
	if _, err := os.Stat(path); err == nil {
		dir_exists_outcome = 1
	} else if errors.Is(err, os.ErrNotExist) {
		dir_exists_outcome = 0
	} else {
		dir_exists_outcome = 0
	}

	return dir_exists_outcome
}
