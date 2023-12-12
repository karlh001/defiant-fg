/*
This script display help output to the
user.

Karl Hunter
2023-11-28
From version 1.2.2

*/

package main

import (
	"fmt"
)

func display_help() {

	fmt.Println("==========================================")
	fmt.Println("DEFIANT-FG HELP\nFor the manual run command 'man dfg'")
	fmt.Println("==========================================")
	fmt.Println("You are on version:", app_ver)
	fmt.Println("Build on:", app_date)
	fmt.Println("")
	fmt.Println("FG is a powerful, light-weight, open source, command line tool\nto monitor for corruption, changes or missing files\nin a directory structure.")
	fmt.Println("")
	fmt.Println("Basic usage: dfg -d /path/to/directory/")
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("")
	fmt.Println("-d			Directory selector")
	fmt.Println("-version		Output program current version")
	fmt.Println("-l			Create log file within directory scanned")
	fmt.Println("-e			Out log/output errors, ignores info:")
	fmt.Println("-s			Skips user input, such as confirm db creation")
	fmt.Println("-log			Location to the log file. Need to use -l too.")
	fmt.Println("-db			Location to the database; default root of directory")
	fmt.Println("")
	fmt.Println("dfg [-version | -help] [-l -e -s -d] /path/to/directory/")
	fmt.Println("")
	fmt.Println("Full help manual visit: https://karlhunter.co.uk/defiant")
	fmt.Println("")

}
