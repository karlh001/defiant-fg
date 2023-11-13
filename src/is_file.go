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

)

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

