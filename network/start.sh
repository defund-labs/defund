#!/bin/bash

BINARY=defundd
CHAIN_DIR=./network/data
CHAINID_1=defund
CHAINID_2=osmosis
CHAINID_3=gaia
GRPCPORT_1=7090
GRPCPORT_2=8090
GRPCPORT_3=9090
GRPCWEB_1=7091
GRPCWEB_2=8091
GRPCWEB_3=9091


echo "Starting $CHAINID_1 in $CHAIN_DIR..."
echo "Creating log file at $CHAIN_DIR/$CHAINID_1.log"
$BINARY start --log_level trace --log_format json --home $CHAIN_DIR/$CHAINID_1 --pruning=nothing --grpc.address="0.0.0.0:$GRPCPORT_1" --grpc-web.address="0.0.0.0:$GRPCWEB_1" > $CHAIN_DIR/$CHAINID_1.log 2>&1 &

echo "Starting Osmosis in $CHAIN_DIR..."
echo "Creating log file at $CHAIN_DIR/osmosis.log"
osmosisd start --log_level trace --log_format json --home $CHAIN_DIR/$CHAINID_2 --pruning=nothing --grpc.address="0.0.0.0:$GRPCPORT_2" --grpc-web.address="0.0.0.0:$GRPCWEB_2" > $CHAIN_DIR/$CHAINID_2.log 2>&1 &

echo "Starting Cosmos/Gaia in $CHAIN_DIR..."
echo "Creating log file at $CHAIN_DIR/gaia.log"
gaiad start --log_level trace --log_format json --home $CHAIN_DIR/$CHAINID_3 --pruning=nothing --grpc.address="0.0.0.0:$GRPCPORT_3" --grpc-web.address="0.0.0.0:$GRPCWEB_3" > $CHAIN_DIR/$CHAINID_3.log 2>&1 &
