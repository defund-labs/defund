#!/bin/bash

echo "What timeout height for the packet will you like (4-7354590)?"
read packettimeoutheight

defundd tx etf create IBC3 "10000000ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518" "channel-0" --from demowallet1 --keyring-backend test --home ./network/data/defund --packet-timeout-height $packettimeoutheight --packet-timeout-timestamp 0 --gas auto -y