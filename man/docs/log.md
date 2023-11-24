# Logging

With the `-l` flag you can enable log file writing. If supplied, your terminal window will no longer show outputs, but instead this will be output to a log file within the path of the directory.

For example

	dfg -l -d /path/to/files

Will result in a log file:

	/path/to/files/dfg.log


The log file will append, which means every time you run the the application, new log entries will follow older entries. 