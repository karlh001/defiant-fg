/*
This mod checks whether a given string is a 
directory. Returns int of 1 (exists) and 
0 (does not exist or not directory)

Karl Hunter 2023
2023-11-12
https://www.karlhunter.co.uk/go

*/

package main

import (
	"os"
	"errors"
)

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