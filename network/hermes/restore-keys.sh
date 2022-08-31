#!/bin/bash
set -e

# Load shell variables
. ./network/hermes/variables.sh

### Sleep is needed otherwise the relayer crashes when trying to init
sleep 1s
### Restore Keys
$HERMES_BINARY --config $CONFIG_DIR keys add --key-name defund --mnemonic-file $HERMES_DIRECTORY/mnemonic-defund.txt --chain defund --overwrite
sleep 5s

$HERMES_BINARY --config $CONFIG_DIR keys add --key-name osmo-test-4 --mnemonic-file $HERMES_DIRECTORY/mnemonic-osmo.txt --chain osmo-test-4 --overwrite
sleep 5s