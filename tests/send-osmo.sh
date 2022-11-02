echo "What is the transfer channel id on osmosis?"
read channel

osmosisd tx ibc-transfer transfer transfer $channel defund1m9l358xunhhwds0568za49mzhvuxx9uxtnevlv 10000000uosmo --from test --keyring-backend test --chain-id osmo-test-4 --node tcp://162.55.134.55:26657