# Errors

Common errors and what they mean.

## Database

	error: could not create database file. Can I write to this directory?

As it states, is the directory you are scanning read only? FG needs access to write the database file.

	error: Could not open db

Does FG have read and write acces to the db? If that's the case, db may be damaged. Return a back-up copy.

	fatal: SQL query error

Although FG was able to open the database, it could not query the rows. This is bad. Likely damaged database file - restore from backup and try again.

## Hash

	error: failed hash on /path/to/file.JPG

This is what FG was designed to do. This messages means the file has been altered or became damaged. Restore the file from back and run the scan again.

## File System

	Fatal: directory does not exist

You provided a location that is not a directory. Check your input.

	missing: /path/to/file.PNG

The file is no longer present. It may have been deleted, moved or renamed. If you moved or renamed, the old file meta data will persist in the database. 
