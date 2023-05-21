#!/bin/bash
set -e

# Load shell variables
source ./network/relayer/variables.sh

# Delete previous config directory if present
if [ -d "$RELAYER_DIRECTORY/config" ]; then
    rm -r "$RELAYER_DIRECTORY/config"
fi

# Init the config
echo "Initializing config file..."
rly config init --home "$RELAYER_DIRECTORY"

# Move already provided config file to initialized config directory
rm "$CONFIG_DIRECTORY/config.yaml"
cp "$RELAYER_DIRECTORY/config.yaml" "$CONFIG_DIRECTORY"

sleep 2
