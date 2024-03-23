# Release Notes

## Version 1.3.1

2024-04-XX

* Block size flag `--bs` 
* Minor code tidying
* SQLite3 upgraded to version 1.14.22

## Version 1.3

2024-03-14

* Database tools: allow removal of entry
* Database tools: find duplicates
* Removed "info: scanning files" from output
* Fixed typo

## Version 1.2.4

2024-03-07

* Added skip missing file scan flag
* Removed pre-build binary and deb package for ARM processors.

## Version 1.2.3

2023-12-17

* Colour coded terminal lines
	* red for failed hashes
	* orange for missing files
	* blue for new files

## Version 1.2.2

2023-12-12

* New flag: choose own [database](db.md) location
* New flag: choose own [log](log.md) file location
* New flag: show help
* Shows new files in log and in terminal
* Documentation update and error corrections

## Version 1.2.1

2023-11-28

* Rebuild binary to work with older Linux


## Version 1.2

2023-11-28

* Show only errors with `-e` flag
* In log more with `-e` still shows start / finish time

## Version 1.1

2023-11-25

* Added log [file](log.md) flag
* Deb packager installer

## Version 1.0.1

2023-11-24

* Minor typos fixes
* README typos fixed
* Fixed email address in about on application
* Built using earliest support Ubuntu LTS for compatibility 

## Version 1.0

2023-11-23

Initial public release.

Features

* Scan directory
* Hash files and store in database
* Check for missing files

Dependency versions

mattn [sqlite3 driver](https://github.com/mattn/go-sqlite3) `1.14.16`