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
	//"compress/gzip"
	//"bufio"
	//"io/ioutil"

)


func main() {

	var path string
	var choice int

	fmt.Println("Weclome to DEFIANT-FG ver 0.3")
	
	fmt.Printf("[1] Scan a directory\n[2] About")
	
	fmt.Println("\nChoose an option")

	fmt.Scan(&choice)

	// Menu
	switch choice {
	case 1:
		fmt.Println("Choose directory to scan:")
		fmt.Scan(&path)
		scan(path)
	case 2:
		about_info()
	}
		

}

func scan(path string) {

	// Get user argument
	//path = "/media/karl/MassStor/test/sj/"

	// Do checks on path
	// Does it end with /
	slash_check := path[len(path)-1:]
	
	// Check the directory path if exists
	if dir_exists(path) == 1 {
		log.Println("info: scanning directory:", path)
	} else {
		// Does not exist
		// Quit application
		log.Fatal("Fatal: directory does not exist")
	}
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

	// Send path to function and cycle through all files
	// and directories to generate hashes
	iterate(path, path_count)

	// Run a scan to check for missing files
	log.Println("info: checking for missing files")
	missing_files_scan(path)

	// Make backup
	//backup_db(path)

	log.Println("info: finished")

}

func backup_db(path string) {

	// Get home dir
	dirname, err := os.UserHomeDir()
    if err != nil {
        log.Println(err)
    }

	// Check backup dir exists

	backup_dir := dirname + "/.defiantfg/backup"

	err = os.MkdirAll(backup_dir, 0775)
	if err != nil {
		log.Println("error: could not create backup directory")
	} else {
		log.Println("info: db backup complete")
	}

	// Copy database file
	// Compress

}

func about_info() {

	// Information about the program
	fmt.Printf("===\nDEFIANT File Guard (DEFIANT-FG)\nBy Karl Hunter\nhttps://karlhunter.co.uk/defiant/\nfg@karlhunter.co.uk\n===\n")

}