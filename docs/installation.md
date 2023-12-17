# Installation

## On Linux 

### Deb Package

On Debian based operating systems (such as Ubuntu, Linux Mint), download the deb package. On desktop versions, open package and click `install package`. 

Or, via command line:

Download

	wget https://github.com/karlh001/defiant-fg/raw/main/bin/deb_pkgs/dfg_amd64.deb

Install (as root)

	dpkg --install dfg_amd64.deb

Done.

### Manual Install

However, if you are running another Linux operating system or prefer not use use deb packages follow these instructions:

For manual installation on Linux, download the source code from the releases section, extract the tar and change directory to the build directory (e.g. `cd ~/Downloads/defiant-fg-main/`.

	cd ~/Downloads
	
	wget https://github.com/karlh001/defiant-fg/blob/main/bin/linux_amd64/dfg.bin
	
	sudo cp dfg.bin /usr/bin/dfg
	
### Shell Script Install
	
To install via shell script, download the most recent release.

_Replace 1.X.X with version_

	tar -xzvf v1.X.X.tar.gz

	sh build/INSTALL_LINUX.sh

_Do not execute script from within build folder or it will fail_


## Test

Test install. Run:

	dfg -version

Done.

To skip installer, and copy as a stand-alone executable, extract the contents and copy dfg.bin anywhere you desire. Make executable with command:  

	chmod +x dfg.bin 


And run: 
 

	./dfg.bin 


## Next

* [Run](run.md)

