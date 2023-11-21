#!/bin/sh

echo "Installation started"

# Deletes older installations

echo "Checking for older installs"

if [ -e /usr/lib/defiantfg ]
then
	echo "Found and deleted"
	rm -r /usr/lib/defiantfg
fi

if [ -e /usr/bin/dfg ]
then
    rm /usr/bin/dfg
fi





## Copies data

echo "Copying new data"

mkdir /usr/lib/defiantfg
cp -r * /usr/lib/defiantfg/
chmod +x /usr/lib/defiantfg/bin/dfg.bin
cp /usr/lib/defiantfg/bin/dfg.bin /usr/bin/dfg


# Finished

echo "Finished; you may test the install by running: dfg"
