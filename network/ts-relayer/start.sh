#!/bin/bash

CHAIN_DIR=./network/data
CHAINID_1=defund

echo "creating Defund relayer account and funding..."
defundd tx bank send defund1m9l358xunhhwds0568za49mzhvuxx9uxtnevlv defund1y295kyv2upsy6swhj0dulghf208ngec5k7zpjq 100000000ufetf --home $CHAIN_DIR/$CHAINID_1 --keyring-backend=test -y

# create connections
echo "Creating connections..."
node $HOME/ts-relayer/build/binary/ibc-setup/index.js connect --mnemonic "alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"

read srcconnection
read destconnection

# start up relayer polling every 6 seconds
echo "Starting Defund ts-relayer..."
node $HOME/ts-relayer/build/binary/ibc-relayer/index.js start --mnemonic "alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart" --src-connection $srcconnection --dest-connection $destconnection --poll 6