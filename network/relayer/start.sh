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
echo "Starting Cosmos relayer for Defund ----> Cosmos Hub..."
$RELAYER_BINARY start defundhub --home $RELAYER_DIRECTORY > $RELAYER_DIRECTORY/defundhub.log 2>&1 &

echo "Starting Cosmos relayer for Cosmos Hub ----> Defund..."
$RELAYER_BINARY start hubdefund --home $RELAYER_DIRECTORY --debug-addr "localhost:7598" > $RELAYER_DIRECTORY/hubdefund.log 2>&1 &
