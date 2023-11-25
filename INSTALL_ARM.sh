#!/bin/sh

echo "Installation started"

# Deletes older installations

echo "Checking for older installs"

if [ -e /usr/lib/defiantfg ]
then
	echo "Found and deleted"
	sudo rm -r /usr/lib/defiantfg
fi

if [ -e /usr/bin/dfg ]
then
    sudo rm /usr/bin/dfg
fi



## Copies data

echo "Copying new data"

sudo mkdir /usr/lib/defiantfg
sudo cp -r * /usr/lib/defiantfg/
sudo chmod +x /usr/lib/defiantfg/bin/arm_aarch64/dfg.bin
sudo cp /usr/lib/defiantfg/bin/arm_aarch64/dfg.bin /usr/bin/dfg
sudo cp /usr/lib/defiantfg/man/dfg.1 /usr/share/man/man1/

# Finished

echo "Finished; you may test the install by running: dfg"

