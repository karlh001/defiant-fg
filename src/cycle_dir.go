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


func iterate(path string) {
    filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Fatalf(err.Error())
			return nil
        }
        
		// Ask if the file now exist
		// if Dir, will look like
		// pathpath with same name added on top
		// is_dir.go
		is_dir := dir_exists(path)

			if is_dir == 1 {

				// Feed the path and file through the hash function
				// hash.go
				file_hash := hash_function(path)

				log.Println(path, file_hash)

			} else {

				log.Println("error: ", path)
			}

		return nil
    })

}