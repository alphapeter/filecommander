#!/bin/sh

# Migrate config if necessary
if [ ! -d $SNAP_COMMON/settings.json ]; then
    cp $SNAP/conf/settings.json $SNAP_COMMON/settings.json
fi

$SNAP/bin/filecommander --settings $SNAP_COMMON/settings.json

