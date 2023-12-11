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