/* 
First mod run by Go
Main screen to accept user inputs
https://karlhunter.co.uk/defiant
*/

package main

import (

	"fmt"
	"os"
	"log"
	"unicode/utf8"
	"flag"
    "path/filepath"

)

const app_ver string = "1.0"
const app_date string = "2023-11-22"
const db_name string = "dfg.db"

func main() {

	// variables declaration  
	var path string
	version := false
	skip := false

	// flags declaration using flag package
	flag.StringVar(&path, "d", " ", "Specify directory")
	flag.BoolVar(&version, "version", false, "Print version information")
	flag.BoolVar(&skip, "s", false, "Skip confirmation message")
	flag.Parse()

	//path := filepath.Clean(flag_path)
	if version == true {
		// Display about info
		about_info()
	} else if path != "" {
		path = filepath.Clean(path)
		scan(path, skip)
	}
	

	
}

func scan(path string, skip bool) {

	// Do checks on path
	// Does it end with /
	//slash_check := path[len(path)-1:]
	

	// Check the directory path if exists
	if dir_exists(path) == 1 {
		// Directory exists
		log.Println("info: scanning directory:", path)
	} else {
		// Does not exist
		// Quit application
		log.Fatal("Fatal: directory does not exist")
	}

	//if slash_check != "/" {
		// Add the slash at the end of string
	//	slash_check = slash_check + "/"
	//	log.Println("info: added trailing slash to path", path + "/")
	//} 

	// Count the full path length provided to get to the 
	// working directory. Later used as shortened path to 
	// save into SQL database so that directory can move
	// round and be easily worked with
	path_count := utf8.RuneCountInString(path)

	// Check if there is a database file
	if is_file(filepath.Join(path, db_name)) == 0 {
		// No database exists so create one
		// Double check with user just in case they 
		// gave an incorrect directory to scan
		// Ask user whether they want to continue
		var choice string
		
		if skip == true {
			// User specified -s so do not bother them with
			// questions
			choice = "y"

		} else {
			// Confirmation message
			fmt.Println("No datbase file: do you want to continue (y/n)?")
			fmt.Scan(&choice)
		}


		switch choice {
		case "y":

			// Yes selected, so user happy to create the database
			// file and contine
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

			// 0 means skip the missing file scan
			start_scan(path, path_count, 0)

		case "n":
			os.Exit(1)

		default:
			log.Println("Incorrect option selected. Please choose only 'y' or 'n'")
			os.Exit(1)
		}

	} else {

		// This directory has been scanned before, so run
		// function and tell not to do the missing file scan
		// 1 sent through function to say do not run missing files

		start_scan(path, path_count, 1)

	}


}

func start_scan(path string, path_count int, look_missing int) {

	
		// Db file is new, so no point looking for
		// missing files as this is first run on directory
		// Send path to function and cycle through all files
		// and directories to generate hashes
		iterate(path, path_count)

		// Run a scan to check for missing files
		// Only run is db was existing
		if look_missing != 0 {
			log.Println("info: checking for missing files")
			missing_files_scan(path)
		}

		// Make backup
		//backup_db(path)

		log.Println("info: finished")

}


func backup_db(path string) {

	// BACKUP FEATURE NOT YET COMPLETED
	// PLEASE HELP IF CAN
	// WANT TO RUN AFTER SCAN, COMPRESS THE DATABASE AND 
	// STORE ~/.defiant-fg/backup
	// e.g. current working directory followed by db
	// YYYYMMDDHHMM_pictures_dfg.db
	// Cleaning perhaps keep last 5 copies

	// Get home dir
	dirname, err := os.UserHomeDir()
    if err != nil {
        log.Println(err)
		os.Exit(1)
    }

	// Check backup dir exists
	backup_dir := dirname + "/.defiantfg/backup"

	err = os.MkdirAll(backup_dir, 0775)
	if err != nil {
		log.Println("error: could not create backup directory")
	}


}


func about_info() {

	fmt.Println("DEFIANT File Guard (DEFIANT-FG)\nVersion", app_ver ,"\nDate", app_date)
	fmt.Println("By Karl Hunter\nhttps://karlhunter.co.uk/defiant/\nfg@karlhunter.co.uk")
	fmt.Println("To scan a directory, add the -d flag followed by directory\ne.g. dfg -d /path/to/dir/")

}