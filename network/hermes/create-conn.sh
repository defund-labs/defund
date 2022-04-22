#!/bin/bash
set -e

# Load shell variables
. ./network/hermes/variables.sh

### Configure the clients and connection
echo "Initiating connection handshake..."
hermes create connection defund-private-1 theta-testnet-001

sleep 2
