#!/bin/bash

BINARY=defundd
CHAIN_DIR=./network/data
CHAINID_1=defund

echo "Starting $CHAINID_1 in $CHAIN_DIR..."
$BINARY start --rpc.laddr tcp://0.0.0.0:26657 --log_level debug --log_format json --home $CHAIN_DIR/$CHAINID_1 --pruning=nothing