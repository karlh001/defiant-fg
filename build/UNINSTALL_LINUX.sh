#!/bin/sh

echo "Uninstall started"

# Deletes older installations


if [ -e /usr/bin/dfg ]
then
    sudo rm /usr/bin/dfg
fi


echo "Finished; removed dfg"
