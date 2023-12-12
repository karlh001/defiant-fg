/*
First mod run by Go
Main screen to accept user inputs
https://karlhunter.co.uk/defiant
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unicode/utf8"
)

const app_ver string = "1.2.2"
const app_date string = "2023-12-12"
const db_name string = "dfg.db"
const log_name string = "dfg.log"

func main() {

	// variables declaration
	var path string
	var logfile string
	var dbfile string
	noinfo := false
	version := false
	skip := false
	logon := false
	flag_help := false

	// flags declaration using flag package
	flag.StringVar(&path, "d", " ", "Specify directory")
	flag.BoolVar(&version, "version", false, "Print version information")
	flag.BoolVar(&skip, "s", false, "Skip confirmation message")
	flag.BoolVar(&logon, "l", false, "Output log file to the scanned directory")
	flag.StringVar(&logfile, "log", " ", "User defined log file")
	flag.StringVar(&dbfile, "db", " ", "User defined log file")
	flag.BoolVar(&noinfo, "e", false, "Skips info log entries, only shows errors")
	flag.BoolVar(&flag_help, "help", false, "Program help info")
	flag.Parse()

	// Enable writing log to file
	if logon == true {
		logging(path, noinfo, logon, logfile)
	}

	//path := filepath.Clean(flag_path)
	if version == true {
		// Display about info
		about_info()
	} else if flag_help {
		display_help()
	} else if path != "" {
		path = filepath.Clean(path)
		scan(path, skip, noinfo, logon, dbfile)
	}

}

func logging(path string, noinfo bool, logon bool, logfile string) {

	if logfile != " " {
		path = filepath.Clean(logfile)
	} else {
		path = filepath.Clean(path)
		path = filepath.Join(path, log_name)
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("error: Could not create log file")
	}

	log.SetOutput(file)

	if noinfo == true && logon == true {
		log.Println("info: Scan started")
	}

}

func scan(path string, skip bool, noinfo bool, logon bool, dbfile string) {

	// Check the directory path if exists
	if dir_exists(path) == 1 {
		// Directory exists
		if noinfo == false {
			log.Println("info: scanning directory:", path)
		}
	} else {
		// Does not exist
		// Quit application
		log.Fatal("Fatal: directory does not exist")
	}

	// Count the full path length provided to get to the
	// working directory. Later used as shortened path to
	// save into SQL database so that directory can move
	// round and be easily worked with
	path_count := utf8.RuneCountInString(path)

	// Did user choose own path?
	if dbfile != " " {
		dbfile = filepath.Clean(dbfile)
		if noinfo == false {
			log.Print("info: db location specified ", dbfile)
		}
	} else {
		dbfile = filepath.Join(path, db_name)
	}

	if is_file(dbfile) == 0 {
		// No database exists so create one
		// Double check with user just in case they
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
			db_output := create_database(path, dbfile)
			if db_output == 0 {
				if noinfo == false {
					log.Println("info: database file created")
				}
			} else {
				// For some reason database does not exist
				// Report error to the user and exit
				// Unable to continue with the program without db file
				log.Println("error: could not create database file. Can I write to this directory?")
				os.Exit(1)
			}

			// 0 means skip the missing file scan
			start_scan(path, path_count, 0, noinfo, logon, dbfile)

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

		start_scan(path, path_count, 1, noinfo, logon, dbfile)

	}

}

func start_scan(path string, path_count int, look_missing int, noinfo bool, logon bool, dbfile string) {

	// Db file is new, so no point looking for
	// missing files as this is first run on directory
	// Send path to function and cycle through all files
	// and directories to generate hashes
	iterate(path, path_count, noinfo, dbfile)

	// Run a scan to check for missing files
	// Only run is db was existing
	if look_missing != 0 {

		if noinfo == false {
			log.Println("info: checking for missing files")
		}

		missing_files_scan(path, dbfile)

	}

	// Make backup
	//backup_db(path)

	if noinfo == false {
		log.Println("info: finished")
	} else if noinfo == true && logon == true {
		log.Println("info: finished")
	}

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

	fmt.Println("DEFIANT File Guard (DEFIANT-FG)\nVersion", app_ver, "\nDate", app_date)
	fmt.Println("By Karl Hunter\nhttps://karlhunter.co.uk/defiant/\ndfg@karlhunter.co.uk")
	fmt.Println("To scan a directory, add the -d flag followed by directory\ne.g. dfg -d /path/to/dir/")

}
