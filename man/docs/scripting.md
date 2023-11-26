# Scripting

You can use `dfg` as part of a script. There are three important flags to achieve this. 

Add these flags:

`-e`

`-s`

`-l`

The `-e` flag suppresses info messages only allowing error messages. Add the `-s` flag skips all messages (including errors) on the terminal so to get anything back you need to specify `-l` to output the log file. If you want info messages in the log omitt the `-s` flag.

How it looks together:

	dfg -s -e -l -d /path/to/files/


Any errors will return `1`, and success will return `0`.


