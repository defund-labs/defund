#!/bin/bash

defundd tx wasm store ./odd_number.wasm -y --from joe --chain-id defund --gas=10000000 --fees 10000000ufetf --broadcast-mode=block --keyring-backend test