/* 
First mod run by Go
Main screen to accept user inputs
*/

package main

import (

	"fmt"
	"os"
	"log"
	"unicode/utf8"

)

func main() {

	ver_number := 0.1

	fmt.Println("Weclome to KHBackup", ver_number)

	// Get user argument
	path := "/media/karl/MassStor/test/sj/"

	// Do checks on path
	// Does it end with /
	slash_check := path[len(path)-1:]

	if slash_check != "/" {
		// Add the slash at the end of string
		slash_check = slash_check + "/"
		log.Println("info: added trailing slash to path", path + "/")
	} 

	// Count the full path length provided to get to the 
	// working directory. Later used as shortened path to 
	// save into SQL database so that directory can move
	// round and be easily worked with
	path_count := utf8.RuneCountInString(path)

	// Grab the first x characters for path
	//full_path := path[0:path_count]
	// To remove the full path use:
	short_path := path[path_count:]

	log.Fatal(short_path)

	// Check if there is a database file
	if is_file(path + "datafile.db") == 0 {
		// No database exists so create one
		db_output := create_database(path)
		if db_output == 0 {
			log.Println("info: database file created")
		} else {
			// For some reason database does not exist
			// Report error to the user and exit
			// Unable to continue with the program without db file
			log.Println("error: could not create database file. Can I write to this directory?")
			os.Exit(1)
		}
	}

	// Check the directory path if exists
	if dir_exists(path) == 1 {
		// Run through directory
		// cycle_dir.go
		log.Println("info: scanning directorty:", path)
		iterate(path, path_count)

	} else {
		// Does not exist
		// Quit application
		log.Println("Error: directory does not exist")
		os.Exit(1)

	}

	// Run a scan to check for missing files
	log.Println("info: checking for missing files")
	missing_files_scan(path, short_path)
	
	log.Println("info: finished")
}