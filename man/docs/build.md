# Build From Source

To build source to binary download and install Golang. On debian-based distros try:

     sudo apt install golang-go

Once set-up CD to the source code directory. 

	cd ~/Downloads/src

Run the build command:

	go build -o dfg.bin .

You will have a complied application named `dfg.bin`. Run a command to test it worked:

	./dfg.bin -version

If this works, copy to your system bin

	cp dfg.bin /usr/bin/dfg

Try the run command without the .bin extension

	dfg -version

If this works everything has been successful.  
