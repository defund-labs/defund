#!/bin/bash
set -e

### Sleep is needed otherwise the relayer crashes when trying to init
sleep 1s
### Restore Keys
hermes keys restore defund-private-1 -m "$MNEMONIC"
sleep 5s

hermes keys restore theta-testnet-001 -m "$MNEMONIC"
sleep 5s