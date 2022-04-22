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
defundd config chain-id defund-private-1
defundd init NODE_NAME
```

Open up the config.toml to edit the seeds and persistent peers:

```bash
cd $HOME/.defund/config
nano config.toml
```

Use page down or arrow keys to get to the line that says seeds = "" and replace it with the following:

```bash
seeds = "1b3e596531dd8f36363b13339beed2364900e4c6@104.131.41.157:26656"
```

Next, add persistent peers:

```bash
persistent_peers = "111ba4e5ae97d5f294294ea6ca03c17506465ec5@208.68.39.221:26656"
```

Then press ```Ctrl+O``` then enter to save, then ```Ctrl+X``` to exit


## Genesis State

Download and replace the genesis file:

```bash
cd $HOME/.defund/config

wget https://raw.githubusercontent.com/schnetzlerjoe/defund/main/testnet/private/genesis.json
```

Reset private validator file to genesis state:

```bash
defundd tendermint unsafe-reset-all
```

## Add/Recover Keys
To create new keypair - make sure you save the mnemonics!
```bash
defundd keys add <key-name> 
```
Restore existing wallet with mnemonic seed phrase. You will be prompted to enter mnemonic seed. 
```bash
defundd keys add <key-name> --recover
```
Request tokens in [DeFund Discord](https://discord.com/invite/QuXAdnd7Pc)

## Set Up Defund Service File

Set up a service to allow Defund node to run in the background as well as restart automatically if it runs into any problems:

```bash
sudo tee /lib/systemd/system/defund.service > /dev/null <<EOF
[Unit]
Description=Defund daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=defundd"
Environment="DAEMON_HOME=${HOME}/.defundd"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
Environment="DAEMON_LOG_BUFFER_SIZE=512"
Environment="UNSAFE_SKIP_BACKUP=true"
User=$USER
ExecStart=${HOME}/go/bin/defundd start
Restart=always
RestartSec=3
LimitNOFILE=infinity
LimitNPROC=infinity
[Install]
WantedBy=multi-user.target
EOF
```


## Start Defund Service

Reload and start the service:

```bash
sudo systemctl daemon-reload
systemctl restart systemd-journald
sudo systemctl start defund
```

Check the status of the service:

```bash
sudo systemctl status defund
```

To see live logs of the service:

```bash
journalctl -f -n 100 -u defund -o cat
```

## Create Validator

Show your public key to be used in the create-validator command below

```bash
defundd tendermint show-validator
```

Create your validator

```bash
defundd tx staking create-validator \
  --amount=1000000ufetf \
  --pubkey=$(defundd tendermint show-validator) \
  --moniker="choose a moniker" \
  --chain-id=defund-private-1 \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1000000" \
  --gas="auto" \
  --from=<key_name>
```

Confirm your validator is running by using this command

```bash
defundd query tendermint-validator-set | grep "$(defundd tendermint show-address)"
```

## Useful commands

valoper addr
```bash
defundd keys show <key_name> --bech val -a
```

balance
```bash
defundd q bank balances <key_addr>
```

get commission
```bash
defundd tx distribution withdraw-rewards <valoper_addr> --from <key_name> --commission --gas auto -y
```

get rewards
```bash
defundd tx distribution withdraw-all-rewards --from <key_name> --gas auto -y
```

validators (active set)
```bash
defundd q staking validators --limit=2000 -oj \
| jq -r '.validators[] | select(.status=="BOND_STATUS_BONDED") | [(.tokens|tonumber / pow(10;6)), .description.moniker] | @csv' \
| column -t -s"," | tr -d '"'| sort -k1 -n -r | nl
```

delegate
```bash
defundd tx staking delegate <valoper_addr> <amout_tokens>ufetf --from <key_name> --gas auto -y
```

send
```bash
defundd tx bank send <key_name> <wallet_addr> <amout_tokens>ufetf --gas auto -y
```


Happy Investing!
