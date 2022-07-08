#!/bin/bash

# Load shell variables
. ./network/relayer/variables.sh

# Delete defundhub log file if present
if [ -f  $RELAYER_DIRECTORY/defundhub.log ]; then
    rm -r $RELAYER_DIRECTORY/defundhub.log
fi

# Delete hubdefund log file if present
if [ -f $RELAYER_DIRECTORY/hubdefund.log ]; then
    rm -r $RELAYER_DIRECTORY/hubdefund.log
fi

# Start the Cosmos relayer for both paths
echo "Starting Cosmos relayer for Defund-ICA ----> Osmosis-ICA..."
$RELAYER_BINARY start defundosmosis --home $RELAYER_DIRECTORY > $RELAYER_DIRECTORY/defundosmosis.log 2>&1 &

echo "Starting Cosmos relayer for Defund-Transfer ---> Osmosis-Transfer..."
$RELAYER_BINARY start osmosisdefund --home $RELAYER_DIRECTORY --debug-addr "localhost:7598" > $RELAYER_DIRECTORY/osmosisdefund.log 2>&1 &
