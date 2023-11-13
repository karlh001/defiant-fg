/*
This mod cycles through directory structure

Karl Hunter 2023
2023-11-12
https://www.karlhunter.co.uk/go

*/

package main

import (

		"log"
		"os"
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

			// Remove full path for db entry
			short_path := path[path_count:] 
			
			// Check database to see if we have seen this 
			// file before; need to use the short path
			// because it's the short path stored in sql
			file_check_result := check_file_sql(short_path)

			// Runs this if statement if
			// is a file, not directory AND
			// file is not already known from db
			if is_dir == 1 && file_check_result == 1 {

				// Feed the path and file through the hash function
				// This will return the hash value
				// hash.go
				file_hash := hash_function(path)

				// Add file with hash to the map
				hashmap[short_path] = file_hash

				// Output result to log
				log.Println("new:", short_path, file_hash)

			}

		return nil
    })

}