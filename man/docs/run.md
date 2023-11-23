# Run

To use FG, feed dfg with a directory to work with using the -d flag:

	dfg -d /path/to/dir/ 
 
If this is the first time you have scanned a directory you will be asked to confirm. This is to ensure you have not accidentally typed an incorrect address. To confirm, type `y` and press enter (to skip visit [flags](flags.md) section). If a new database is created you will see:

	info: database file created


The scan will begin. To cancel, hold down the Ctrl key and press C (Ctl+C).

When the scan is complete you will get an confirmation:


For errors that may occur, read the [error page](errors.md).

## What is happening?

During the scan, FG will cycle through the directory and sub-directories. When it finds a file the following occurs: 

 

1. Is the object a file? (i.e. do not process a directory) 

2. Is this file in the database already? (does FG know about it, could be a new file) 
	1. If yes, calculates the file's hash then compares against the database. If no match, reports to user: "error: failed file hash" 
	2. If no, calculates the file's hash and adds to a queue to write to db later 

After this process has finished, FG will then cycle through the database to check for missing files. 

1. Is the file a file 
2. Is the file present 
	1. If yes, continues with scan
	2. if no, reports to user "error: file missing" 

The scan will complete. 

If this is the first time you are running a scan, and have thousands of files, it may take a few hours to write to the database. Further scans on the same directory will work much faster. 
