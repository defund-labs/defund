#!/bin/bash

defundd tx etf create-fund test test test "osmo" "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2:33:osmosis:1:spot,uosmo:34:osmosis:1:spot,ibc/1480B8FD20AD5FCAE81EA87584D269547DD4D436843C1D20F15E00EB64743EF4:33:osmosis:3:spot" 100 5000000 --from demowallet1 --keyring-backend test --home ./network/data/defund --gas auto -y
