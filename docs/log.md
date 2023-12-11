# Logging

## Enable

With the `-l` flag you can enable log file writing. If supplied, your terminal window will no longer show outputs, but instead this will be output to a log file within the path of the directory. Useful for scripting.

For example

	dfg -l -d /path/to/files

Will result in a log file:

	/path/to/files/dfg.log


The log file will append, which means every time you run the the application, new log entries will follow older entries. 

## Location

The default location of the log file, if you specify -l only, will be the directory of the scanned path; however, you can specify your own directory by using the `-log` flag.

For example:

	dfg -l -log /var/log/pictures.log -d /mnt/usb/pictures/

## Timestmp

As log files append, you may want to create a new log file each time the program is run. To do this, add the date command to the `-log` flag:

	dfg -l -log /var/log/$(date +%F)_pictures.log -d /mnt/usb/pictures/

Output

	2023-12-11.log

More on [date](https://www.howtogeek.com/410442/how-to-display-the-date-and-time-in-the-linux-terminal-and-use-it-in-bash-scripts/#how-to-see-the-date-and-time-and-time-in-the-linux-terminal) outputs.


