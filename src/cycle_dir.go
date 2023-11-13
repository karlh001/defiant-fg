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

	// This function will cycle through the directory and print
	// files and directories

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
		
			if is_dir == 1 {

				// Feed the path and file through the hash function
				// This will return the hash value
				// hash.go
				file_hash := hash_function(path)

				log.Println("done:", short_path, file_hash)

			}

		return nil
    })
}