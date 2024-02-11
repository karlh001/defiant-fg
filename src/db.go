/*
This go file is used to interact with the database
Such as remove entries.
*/

package main

import (
	"fmt"
	"os"
)

func db_tool_func(dbcommand string, dbfile string) {

	/* Commands

	clean (deletes all entries with 0) - to do
	del (disabled entry - sets to 0)
	hash (gives user hash of file) - to do

	*/

	if dbcommand == "del" {

		var choice string
		var db_ID int

		// This will disable the entry in given by the user

		fmt.Println("Database tool using the following database:", dbfile)

		fmt.Println("Type ID of record to delete:")
		fmt.Scan(&db_ID)

		fmt.Println("You have chosen:", db_ID)

		fmt.Println("Correct? [y/n]")
		fmt.Scan(&choice)

		switch choice {
		case "y":
			disable_sql_func(dbfile, db_ID)
		default:
			fmt.Println("Delete operation was cancelled")
			os.Exit(1)
		}

	}

}
