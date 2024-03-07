# Database

## About

DEFIANT-FG stores your file hashes into a database (db) file named `dfg.db`. It's formatted to SQLite. 

## Location

By default, db file is stored in the root of the given directory; for example, if you run on `/mnt/usb/Pictures` the db file will be located `/mnt/usb/Pictures/dfg.db`. Any sub folders will be included in the hash database.

You can specify your own db location with name of the database file. To do this, use the `-db` flag, example:

	dfg -db /path/to/db.sql -d /path/to/files/


## Portability

You are encouraged to keep a back-up of the database file. As long as you restore to the root of the original directory (or specify the exact location) it will work, even though the full path changes. For example, you can move the Pictures to: `/mnt/usb2/Pictures`; in fact, the picture directory name can be changed too.

## Structure

* dfg.db
	* objects [table]
		* ID_objects [column] (auto increment)
		* path [column] (subfolder1/filename)
		* hash [column] (SHA-256)
		* ts (time stamp) [column] (YYYYMMDDHHMMSS)
		* enabled [column] default 1

## File Name Limitations

The database will not accept an apostrophes or `<!` in the filename; if present in a file name, these will be replaced with a tag before storing in the database. The apostrophe will become `xAPOSx` and the less than bracket `<!` will become `xBRACKx`.

## Remove Entries
_as of version 1.3_

If you have removed an object and no longer wish it to remain in the database, you can remove it using follow commands.

First, you will need to determine the database ID. This will be printed in a log or terminal output after a scan:

	2024/03/02 16:19:46 missing: /Interesting_Document.pdf ID: 919

Copy / note the ID of the missing object, in this case 919.

	dfg -data del -db /path/to/dfg.db

The database tool will ask for ID of record to delete; this is this ID given in terminal output or log file, in this example it's 919.

	Database tool using the following database: /path/to/dfg.db

	Type ID of record to delete:
	919
	File entry to delete: /Interesting_Document.pdf
	Correct? [y/n]
	y
	Delete request completed



