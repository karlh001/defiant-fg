# Flags

### Directory
`-d /path/to/dir/`

Choose directory to scan. For example:

	dfg -d /mnt/usb/pictures/

Note: the -d should be the last argument. Use other flags before this. 

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

`-log'

Specify the location of the log file. Log entries append to previous log lines (log files do not overwrite).

Example:

	dfg -l -log /var/log/pictures.log -d /mnt/usb/pictures/

### Version

`-version`

Outputs version of the application



