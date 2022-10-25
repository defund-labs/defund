echo "What is the transfer channel id on osmosis?"
read channel

osmosisd tx ibc-transfer transfer transfer $channel defund1p0v6m6nu94xw5cm29qrdcsj2wudyts7pcj22s3 10000000uosmo --from defund --keyring-backend test --chain-id osmo-test-4 --node tcp://162.55.134.55:26657