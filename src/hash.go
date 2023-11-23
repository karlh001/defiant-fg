/*
This mod hashes a file, and returns the sha256 value
through the function as string.
Remember to call to get the file back:
the_hash := hash_function(file)

Karl Hunter 2023
2023-11-12
https://www.karlhunter.co.uk/defiant

*/

package main

import (

	"crypto/sha256"
	"io"
	"log"
	"os"
	"fmt"
	
)

// Hash the file function
// Add return value as string
func hash_function(file_to_run string) string {
	
	// 64 MB block size
	const BlockSize = 64
	
	// Open the file passed through function
	f, err := os.Open(file_to_run)
	
	// Error checking during hash
	if err != nil {
		log.Println("skip:", err)
	}
	defer f.Close()

	// The has function usng SHA-256
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		// Need to convert error to string to 
		// check if standard error about dir, so
		// we can ignore it. Only want to know about
		// failed hashing on files.
		var msg_error string
		msg_error = fmt.Sprintf("%v", err)
		msg_error = msg_error[9:]
		if msg_error == "directory" { log.Println("skip:", err) }
	}

	// Put the hash into variable to return
	actual_sum := fmt.Sprintf("%x", h.Sum(nil))
	
	// Returns the string back to the caller
	return actual_sum
  
}