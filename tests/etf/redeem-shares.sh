#!/bin/bash

defundd tx etf redeem ODD2 '875000etf/ODD2' '{"osmosisAddress": "osmo1m9l358xunhhwds0568za49mzhvuxx9uxtz8m2l"}' --from demowallet1 --keyring-backend test --home ./network/data/defund --gas auto --fees 1000000ufetf --gas-adjustment 1.5 --chain-id defund -y