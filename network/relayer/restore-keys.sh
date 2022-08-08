#!/bin/bash
set -e

# Load shell variables
. ./network/relayer/variables.sh

# Delete keys if directory is present
if [ -d $RELAYER_DIRECTORY/keys ]; then
    rm -r $RELAYER_DIRECTORY/keys
fi

sleep 1s
### Restore Keys
rly keys restore defund defund "alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart" --home $RELAYER_DIRECTORY
sleep 5s

rly keys restore osmo-test-4 osmo "record gift you once hip style during joke field prize dust unique length more pencil transfer quit train device arrive energy sort steak upset" --home $RELAYER_DIRECTORY
sleep 5s
