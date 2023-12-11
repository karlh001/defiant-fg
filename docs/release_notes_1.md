# Release Notes

## Version 1.2.2

2023-11-XX

* New flag: choose own [database](db.md) location
* New flag: choose own [log](logging.md) file location
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