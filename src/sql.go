package main

// github.com/mattn/go-sqlite3

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
  )
  

func main() {

	db, err := gorm.Open(sqlite.Open("/home/karl/go/projects/khb/src/test.db"), &gorm.Config{})
	
	// If there is an error, report to user
	if err != nil {
		log.Fatalf("db error", err)
		os.Exit(1)
	}

	log.Println("done", db)

}


