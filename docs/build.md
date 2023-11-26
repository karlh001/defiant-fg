# Build From Source

To build source code into a binary you need to download and install Golang. On debian-based distros try:

     sudo apt install golang-go

Once set-up change directory to the source code directory. 

	cd ~/Downloads/XYZ/src

Run the build command:

	go build -o dfg.bin .

During the build, Go will download dependencies. 

You will have a complied application named `dfg.bin`. Run a command to test it worked:

	./dfg.bin -version

If this works, copy to your system bin

	cp dfg.bin /usr/bin/dfg

Try the run command without the .bin extension

	dfg -version

If this works everything has been successful.  
