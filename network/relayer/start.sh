#!/bin/bash

# Load shell variables
. ./network/relayer/variables.sh

# Delete defundhub log file if present
if [ -f  $RELAYER_DIRECTORY/defundosmosis.log ]; then
    rm -r $RELAYER_DIRECTORY/defundosmosis.log
fi

# Start the Cosmos relayer for both paths
echo "Starting Cosmos relayer for Defund-ICA ----> Osmosis-ICA..."
$RELAYER_BINARY start defundosmosis --home $RELAYER_DIRECTORY > $RELAYER_DIRECTORY/defundosmosis.log 2>&1 &
