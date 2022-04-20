# Defund Private Testnet

The purpose of the private Defund testnet is to begin to build a relationship between the Defund community, contributors, relayers and validators while testing each functionality in a staged process. The testnet begins with the initial testnet validator launch with all basic chain functionalities and the ability to create funds and query fund price (cross chain!).

The second stage of the private testnet, planned to come shortly after the first, is introducing the invest and uninvest cabilities into funds (including rebalancing of dETF's). This allows for staged testing of the underlying cross-chain investment process plus the testing of Defunds upgrade processes.

The third and final stage of the private testnet is introducing new broker chains plus adding all specialized modules that are needed for mainnet genesis like airdrop claims, automated governance, and emmissions. At the start of the private testnet we only support one broker chain, the Gravity Dex (on the Cosmos Hub), as GDex is the only chain with liquidity that supports ICA currently. By the third stage the hope will be that Osmosis, and potentially Crescent will have ICA capabilities and will then be included during this stage.

At the conclusion of the private testnet, there will be a public testnet that directly emulates Defunds mainnet genesis launch environment.

`Note:` Validators that participate in both the private and public testnet are eligble for a portion of the airdrop (5%) which is fairdropped equally to all validators during the testnets.

## Set Up Validator/Node On Akash

Details on setting up a Defund node and/or validator on Akash are coming very soon. We actively promote and encourage validators and runners of nodes in general, to use Akash maximizing the decentralization of Defund.

## Joining the Testnet

### Install the Defund binary

```
git clone https://github.com/defund-labs/defund

cd defund

make install
```

## Initialize Defund Node

```bash
defundd init NODE_NAME --chain-id=defund-private-1
```

Open up the config.toml to edit the seeds and persistent peers:

```bash
cd $HOME/.defundd/config
nano config.toml
```

Use page down or arrow keys to get to the line that says seeds = "" and replace it with the following:

```bash
seeds = "0f9a9c694c46bd28ad9ad6126e923993fc6c56b1@137.184.181.105:26656"
```

Next, add persistent peers:

```bash
persistent_peers = "4ab030b7fd75ed895c48bcc899b99c17a396736b@137.184.190.127:26656,3dbffa30baab16cc8597df02945dcee0aa0a4581@143.198.139.33:26656"
```

Then press ```Ctrl+O``` then enter to save, then ```Ctrl+X``` to exit

## Set Up Cosmovisor

Set up cosmovisor to automate many aspects of future defundd upgrades. To install Cosmovisor:

```bash
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0.0
```

Create the required directories:

```bash
mkdir -p ~/.defundd/cosmovisor
mkdir -p ~/.defundd/cosmovisor/genesis
mkdir -p ~/.defundd/cosmovisor/genesis/bin
mkdir -p ~/.defundd/cosmovisor/upgrades
```

Set the environment variables:

```bash
echo "# Setup Cosmovisor" >> ~/.profile
echo "export DAEMON_NAME=defundd" >> ~/.profile
echo "export DAEMON_HOME=$HOME/.defundd" >> ~/.profile
echo "export DAEMON_ALLOW_DOWNLOAD_BINARIES=false" >> ~/.profile
echo "export DAEMON_LOG_BUFFER_SIZE=512" >> ~/.profile
echo "export DAEMON_RESTART_AFTER_UPGRADE=true" >> ~/.profile
echo "export UNSAFE_SKIP_BACKUP=true" >> ~/.profile
source ~/.profile
```

Download and replace the genesis file:

```bash
# Need to Add
```

Copy the current defundd binary into the cosmovisor/genesis folder:

```bash
cp $GOPATH/bin/defundd ~/.defundd/cosmovisor/genesis/bin
```

To check your work, ensure the version of cosmovisor and defundd are the same:

```bash
cosmovisor version
defundd version
```

These two commands should both output `0.0.2`

Reset private validator file to genesis state:

```bash
defundd unsafe-reset-all
```

## Set Up Defund Service File

Set up a service to allow cosmovisor to run in the background as well as restart automatically if it runs into any problems:

```bash
echo "[Unit]
Description=Cosmovisor daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=defundd"
Environment="DAEMON_HOME=${HOME}/.defundd"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
Environment="DAEMON_LOG_BUFFER_SIZE=512"
Environment="UNSAFE_SKIP_BACKUP=true"
User=$USER
ExecStart=${HOME}/go/bin/cosmovisor start
Restart=always
RestartSec=3
LimitNOFILE=infinity
LimitNPROC=infinity
[Install]
WantedBy=multi-user.target
" >cosmovisor.service
```

Move this new file to the systemd directory:

```bash
sudo mv cosmovisor.service /lib/systemd/system/cosmovisor.service
```

## Start Defund Service

Reload and start the service:

```bash
sudo systemctl daemon-reload
systemctl restart systemd-journald
sudo systemctl start cosmovisor
```

Check the status of the service:

```bash
sudo systemctl status cosmovisor
```

To see live logs of the service:

```bash
journalctl -u cosmovisor -f
```

## Create Validator

Show your public key to be used in the create-validator command below

```bash
gaiad tendermint show-validator
```

Create your validator

```bash
gaiad tx staking create-validator \
  --amount=1000000uatom \
  --pubkey=$(gaiad tendermint show-validator) \
  --moniker="choose a moniker" \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1000000" \
  --gas="auto" \
  --gas-prices="0.0025uatom" \
  --from=<key_name>
```

Confirm your validator is running by using this command

```bash
gaiad query tendermint-validator-set | grep "$(gaiad tendermint show-address)"
```

Happy Investing!
