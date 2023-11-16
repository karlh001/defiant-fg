/*
This mod checks whether a given string is a 
directory.

Returns 1 for directory, 0 for not (likely file)

Karl Hunter 2023
2023-11-12
https://www.karlhunter.co.uk/go

*/

package main

import (
	"os"
	"errors"
	"log"
	"path/filepath"
)


func iterate(path string, path_count int) {

	// Create a map to store data
	hashmap := map[string]string{}

	// This function will cycle through the directory and print
	// files and directories.
	// If there was an error, e.g. permissions, then error message
	// output to the user
    filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

        if err != nil {
            log.Println("error:", err.Error())
			return nil
        }
	
		// Ask if the file now exist
		// if Dir, will look like
		// pathpath with same name added on top
		// is_dir.go
		is_dir := is_file(path)

		// Feed the path and file through the hash function
		// This will return the hash value
		// hash.go
		file_hash := hash_function(path)

		// Check database to see if we have seen this 
		// file before; need to use the short path
		// because it's the short path stored in sql
		file_check_result := check_file_sql(short_path, path)

		// Runs this if statement if
		// is a file, not directory AND
		// file is not already known from db
		if is_dir == 1 && file_check_result == 1 {
		//if is_dir == 1 {

			// Add file with hash to the map
			// This will be sent later to insert
			// into the database later
			hashmap[short_path] = file_hash

		}		

		return nil
    })

	// Send map of new files to insert SQL function
	write_files_sql(path, hashmap)

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