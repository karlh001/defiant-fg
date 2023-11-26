# Scripting

You can use `dfg` as part of a script. There are three important flags to achieve this. 

Add these flags:

`-e`

`-s`

`-l`

The `-e` flag suppresses info messages only allowing error messages. Add the `-s` flag skips all messages (including errors) on the terminal so to get anything back you need to specify `-l` to output the log file. If you want info messages in the log omitt the `-s` flag.

How it looks together:

	dfg -s -e -l -d /path/to/files/


If you supply both the `-l` and `-e` flags together this will result in the following being written to the log file (need to know time scan started or finished). If there are any errors these will display between the start and end sections of the log.

	2023/11/26 13:35:59 info: Scan started
	2023/11/26 13:35:59 info: finished

If you want to check if script was successful look for a `0` return:

Any errors will return `1`, and success will return `0`.


