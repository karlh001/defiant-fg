/*
This go file is used to interact with the database
Such as remove entries.
*/

package main

import (
	"fmt"
)

func db_tool_func(dbcommand string, dbfile string) {

	/* Commands

	clean (deletes all entries with 0) - to do
	del (disabled entry - sets to 0)
	hash (gives user hash of file) - to do
	count (returns total active records)

	*/
	var err int
	err = 1

	if dbcommand == "del" || dbcommand == "delete" {
		var db_ID int
		fmt.Println("Database tool using the following database:", dbfile)
		fmt.Println("Type ID of record to delete:")
		fmt.Scan(&db_ID)
		disable_sql_func(dbfile, db_ID)
		err = 0
	}

	if dbcommand == "count" || dbcommand == "total" {
		count_sql_func(dbfile)
		err = 0
	}

	if dbcommand == "same" || dbcommand == "duplicate" {
		check_same_files(dbfile)
		err = 0
	}

	if err > 0 {
		fmt.Println("error: data command not recognised or was not declared.")
	}

}
