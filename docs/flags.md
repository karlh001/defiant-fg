# Flags

### Directory
`-d /path/to/dir/`

Choose directory to scan. For example:

	dfg -d /mnt/usb/pictures/

For documentation consistency, I always put the path flag at the end of the command, but this is not necessary. 

### Skip
`-s`

Skip any messages after executing application. Useful for scripting. Will skip confirmation message to create new database file

	dfg -s -d /mnt/usb/pictures/

### Skip info messages

`-e`

Only shows errors such as failed hashes and missing files. Does not show any info prompts during the scan.

Note: will also not show a line for new files.

### Logging

`-l`

Add the l flag to output to a log file. Your terminal will no longer show verbose. More on the [log file](log.md).

`-log`

Specify the location of the log file. Log entries append to previous log lines (log files do not overwrite).

Example:

	dfg -l -log /var/log/pictures.log -d /mnt/usb/pictures/

### Block Size

The default block size used to hash files is 64 MB. You can specify a custom block size, in MB.

`-b 64`

_Changing the block size does not affect existing hashes already saved in the database_

### Skip missing file scan

`--skip-missing`

Skips the missing file scan. You will not be notified about missing / deleted files which exist in the database.

### Database file location

`-db`

You can specify your own db location with name of the database file. To do this, use the `-db` flag, example:

	dfg -db /path/to/db.sql -d /path/to/files/

Ensure you specify the database name, not just the directory path.

More on [database](db.md).

### Version

`-version`

Outputs version of the application



