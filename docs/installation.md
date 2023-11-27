# Installation

## On Linux 

On Debian based operating systems (such as Ubuntu, Linux Mint), download the deb package. On desktop versions, open package and click `install package`. 

Done.

However, if you are running another Linux operating system or prefer not use use deb packages follow this instructions:

For manual installation on Linux, download the source code from the releases section, extract the tar and change directory to the build directory (e.g. `cd ~/Downloads/defiant-fg-main/`, then run:

	sh build/INSTALL_LINUX.sh

_Do not execute script from within build folder or it will fail_

Test install. Run:

	dfg -version

Done.

To skip installer, and copy as a stand-alone executable, extract the contents and copy dfg.bin anywhere you desire. Make executable with command:  

	chmod +x dfg.bin 


And run: 
 

	./dfg.bin 


## Next

* [Run](run.md)