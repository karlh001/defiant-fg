#!/bin/sh

echo "Installation started"

# Deletes older installations

echo "Checking for older installs"

if [ -e /usr/bin/dfg ]
then
    sudo rm /usr/bin/dfg
fi



## Copies data

echo "Copying new data"

sudo cp bin/arm_aarch64/dfg.bin /usr/bin/dfg
sudo cp docs/dfg.1 /usr/share/man/man1/
sudo chmod +x /usr/bin/dfg


# Finished

echo "Finished; you may test the install by running: dfg"

