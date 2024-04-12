# Installation

## On Linux 

### Deb Package

On Debian based operating systems (such as Ubuntu, Linux Mint), download the deb package. On desktop versions, open package and click `install package`. 

Or, via command line:

Download

	wget https://github.com/karlh001/defiant-fg/raw/main/dist/dfg_amd64.deb

Install (as root)

	dpkg --install dfg_amd64.deb

Done.

### Manual Install

However, if you are running another Linux operating system or prefer not use use deb packages, find the build for your system by visiting the distribution directory:

[https://github.com/karlh001/defiant-fg/tree/main/dist/](https://github.com/karlh001/defiant-fg/tree/main/dist/)


	

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

