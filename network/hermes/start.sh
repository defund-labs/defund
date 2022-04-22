#!/bin/bash

echo "Hermes Relayer Version Check......"
hermes version

# Start the hermes relayer in multi-paths mode
echo "Starting hermes relayer..."
hermes start