#!/bin/bash

BINARY=defundd
CHAIN_DIR=./network/data
CHAINID_1=defund
CHAINID_2=osmo-test-4
VAL_MNEMONIC_1="clock post desk civil pottery foster expand merit dash seminar song memory figure uniform spice circle try happy obvious trash crime hybrid hood cushion"
DEMO_MNEMONIC_1="banner spread envelope side kite person disagree path silver will brother under couch edit food venture squirrel civil budget number acquire point work mass"
DEMO_MNEMONIC_2="veteran try aware erosion drink dance decade comic dawn museum release episode original list ability owner size tuition surface ceiling depth seminar capable only"
RLY_MNEMONIC_1="alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"
RLY_MNEMONIC_2="record gift you once hip style during joke field prize dust unique length more pencil transfer quit train device arrive energy sort steak upset"

# Stop interammd if it is already running 
if pgrep -x "$BINARY" >/dev/null; then
    echo "Terminating $BINARY..."
    killall $BINARY
fi

echo "Removing previous data..."
rm -rf $CHAIN_DIR/$CHAINID_1 &> /dev/null

# Add directories for all chains, exit if an error occurs
if ! mkdir -p $CHAIN_DIR/$CHAINID_1 2>/dev/null; then
    echo "Failed to create chain folder. Aborting..."
    exit 1
fi

cd $HOME/defund
echo "Initializing $CHAINID_1..."
$BINARY init test --home $CHAIN_DIR/$CHAINID_1 --chain-id=$CHAINID_1

echo "Adding accounts..."
echo $VAL_MNEMONIC_1 | $BINARY keys add val --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test
echo $DEMO_MNEMONIC_1 | $BINARY keys add demowallet1 --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test
echo $DEMO_MNEMONIC_2 | osmosisd keys add demowallet2 --home $CHAIN_DIR/$CHAINID_2 --recover --keyring-backend=test
echo $RLY_MNEMONIC_1 | $BINARY keys add defund --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test 
echo $RLY_MNEMONIC_2 | osmosisd keys add osmo --home $CHAIN_DIR/$CHAINID_2 --recover --keyring-backend=test

$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show val --keyring-backend test -a) 100000000000ufetf  --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show demowallet1 --keyring-backend test -a) 100000000000ufetf  --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show defund --keyring-backend test -a) 100000000000ufetf  --home $CHAIN_DIR/$CHAINID_1

echo "Creating and collecting gentx..."
$BINARY gentx val 7000000000ufetf --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --keyring-backend test
$BINARY collect-gentxs --home $CHAIN_DIR/$CHAINID_1

echo "Changing defaults and ports in app.toml and config.toml files..."
sed -i -e 's/cors_allowed_origins = []/cors_allowed_origins = [*]/g' $CHAIN_DIR/$CHAINID_1/config/config.toml
sed -i -e 's/enable = false/enable = true/g' $CHAIN_DIR/$CHAINID_1/config/app.toml
sed -i -e 's/swagger = false/swagger = true/g' $CHAIN_DIR/$CHAINID_1/config/app.toml
sed -i -e 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $CHAIN_DIR/$CHAINID_1/config/app.toml
sed -i -e 's/enable-unsafe-cors = false/enable-unsafe-cors = true/g' $CHAIN_DIR/$CHAINID_1/config/app.toml

echo "Changing Genesis File"
# Change crisis fee denom to fake detf
contents="$(jq '.app_state.crisis.constant_fee.denom = "ufetf"' ./network/data/defund/config/genesis.json)" && \
echo -E "${contents}" > ./network/data/defund/config/genesis.json
# Change mint denom to fake detf
contents="$(jq '.app_state.mint.params.mint_denom = "ufetf"' ./network/data/defund/config/genesis.json)" && \
echo -E "${contents}" > ./network/data/defund/config/genesis.json
# Change staking denom to fake detf
contents="$(jq '.app_state.staking.params.bond_denom = "ufetf"' ./network/data/defund/config/genesis.json)" && \
echo -E "${contents}" > ./network/data/defund/config/genesis.json
# Change gov deposit denom to fake detf
contents="$(jq '.app_state.gov.deposit_params.min_deposit[0].denom = "ufetf"' ./network/data/defund/config/genesis.json)" && \
echo -E "${contents}" > ./network/data/defund/config/genesis.json
# Change brokers in broker in genesis file to blank so it will run broker function
contents="$(jq '.app_state.broker.brokers = []' ./network/data/defund/config/genesis.json)" && \
echo -E "${contents}" > ./network/data/defund/config/genesis.json
# Change base denoms in genesis file
contents="$(jq '.app_state.broker.params.base_denoms.AtomTrace.path = "transfer/channel-0/transfer/channel-0"' ./network/data/defund/config/genesis.json)" && \
echo -E "${contents}" > ./network/data/defund/config/genesis.json
contents="$(jq '.app_state.broker.params.base_denoms.OsmoTrace.path = "transfer/channel-0"' ./network/data/defund/config/genesis.json)" && \
echo -E "${contents}" > ./network/data/defund/config/genesis.json
contents="$(jq '.consensus_params.block.time_iota_ms = "1"' ./network/data/defund/config/genesis.json)" && \
echo -E "${contents}" > ./network/data/defund/config/genesis.json
