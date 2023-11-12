/* 
First mod run by Go
Main screen to accept user inputs
*/

package main

import (

	"fmt"
	"os"
	"log"

)

func main() {

	ver_number := 0.1

	fmt.Println("Weclome to KHBackup", ver_number)

	path := "/media/karl/MassStor/test/"

	// Check the directory path if exists
	if dir_exists(path) == 1 {
		// Run through directory
		// cycle_dir.go
		iterate(path)
	} else {
		// Does not exist
		// Quit application
		log.Println("Error: directory does not exist")
		os.Exit(1)
	}
	
}