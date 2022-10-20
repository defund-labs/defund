#!/bin/bash

BINARY=defundd
CHAIN_DIR=./network/data
CHAINID_1=defund

echo "Starting $CHAINID_1 in $CHAIN_DIR..."
echo "Creating log file at $CHAIN_DIR/$CHAINID_1.log"
$BINARY start --log_level info --log_format json --home $CHAIN_DIR/$CHAINID_1 --pruning=nothing > $CHAIN_DIR/$CHAINID_1.log 2>&1 &