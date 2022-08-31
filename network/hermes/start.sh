
#!/bin/bash

# Load shell variables
. ./network/hermes/variables.sh

echo "Hermes Relayer Version Check......"
$HERMES_BINARY version

# Start the hermes relayer in multi-paths mode
echo "Starting hermes relayer..."
$HERMES_BINARY --config $CONFIG_DIR start > $HERMES_DIRECTORY/hermes.log 2>&1 &