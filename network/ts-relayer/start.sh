#!/bin/bash

CHAIN_DIR=./network/data
RELAY_DIR=./network/ts-relayer
CHAINID_1=defund
MNEMONIC="alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"

#echo "creating Defund relayer account and funding..."
#defundd tx bank send defund1m9l358xunhhwds0568za49mzhvuxx9uxtnevlv defund1y295kyv2upsy6swhj0dulghf208ngec5k7zpjq 100000000ufetf --home $CHAIN_DIR/$CHAINID_1 --keyring-backend=test -y

# create connections
echo "Creating connections..."
ibc-setup connect --interactive --log-level debug

echo "What is the source connection to relay?"
read srcconnection
echo "What is the destination connection to relay?"
read destconnection

# start up relayer polling every 6 seconds
echo "Starting Defund ts-relayer for src:$srcconnection dest:$destconnection..."
ibc-relayer start --src-connection $srcconnection --dest-connection $destconnection --poll 6 --log-level debug --interactive > ./relayer.log 2>&1 &