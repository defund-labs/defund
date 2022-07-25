# Install and Connect Chain to Defund
Defund has its own custom relayer here https://github.com/defund-labs/relayer that supports Defund's IBC interquerying. Right now this version of the Golang relayer does not support interchain accounts (ICS-27) but Hermes does. So for now, Defund requires both the Golang relayer and Hermes to run. This is good practice anyways as one acts as a backup to the other. We are in the process of adding ICS-27 support to our special version of the Golang relayer and thus will only need that relayer in the future.

This example connects Defund to the Gaia theta testnet (Cosmos Hub) which we use the Gravity Dex for liquidity.

## Install Gaia

```bash
cd $HOME
git clone -b v7.0.1 https://github.com/cosmos/gaia
cd gaia
make install
```

You will need to have uatom testnet tokens on Gaia. See here for more details https://github.com/cosmos/testnets/blob/master/v7-theta/public-testnet/README.md. Right now they took down the faucet Discord so messge them in the Cosmos Network discord for tokens https://discord.gg/cosmosnetwork.

## Setting Up Hermes

You will need Rust installed for Hermes. See here https://www.rust-lang.org/tools/install.

### Install Hermes

```bash
# See here for more details https://hermes.informal.systems/installation.html
cd $HOME
cargo install ibc-relayer-cli --bin hermes --locked
```

Make the Config Directory
```bash
mkdir $HOME/.hermes
cd $HOME/.hermes
wget https://raw.githubusercontent.com/defund-labs/defund/main/testnet/private/hermes/config.toml
```

If you do not have https://github.com/defund-labs/defund git cloned. Do so now.

```bash
cd defund
nano ./network/hermes/restore-keys.sh
# Either add the mnemonic you will be using for both chains where you see $MNEMONIC or set the MNEMONIC variables like below
export MNEMONIC="<your_mnemonic>"
# Restore the keys for hermes now
bash ./network/hermes/restore-keys.sh
```

### Create Service File and Startup Hermes

```bash
echo "[Unit]
Description=Defund daemon
After=network-online.target
[Service]
User=$USER
ExecStart=${HOME}/.cargo/bin/hermes start
Restart=always
RestartSec=3
LimitNOFILE=infinity
LimitNPROC=infinity
[Install]
WantedBy=multi-user.target
" >hermes.service
```

```bash
sudo mv hermes.service /lib/systemd/system/hermes.service
```

Reload and start the service:

```bash
sudo systemctl daemon-reload
systemctl restart systemd-journald
sudo systemctl start hermes
```

Check the status of the service:
```bash
sudo systemctl status hermes
```

Check the logs

```bash
journalctl -u hermes -f
```

`Welcome to Inter-Blockchain Communication :)`

## Setting Up Defund Golang Relayer

### Install Golang Relayer

```bash
cd $HOME
git clone https://github.com/defund-labs/relayer
cd relayer
make install
```

### Setup Config
```bash
rly config init
cd $HOME/.relayer/config
sudo rm config.yaml
wget https://raw.githubusercontent.com/defund-labs/defund/main/testnet/private/golang/config.yaml
sudo nano config.yaml
```

```yaml
# Change the key on both chains here to the key you are using
value:
    key: <keyname>
    chain-id: defund-private-1
value:
    key: <keyname>
    chain-id: theta-testnet-001
```

### Restore Keys

```bash
rly keys restore defund-private-1 <keyname> "$MNEMONIC"

rly keys restore theta-testnet-001 <keyname> "$MNEMONIC"
```

### Start Golang Relayer

```bash
rly start defundhub
```

We like to run it in a screen session to moniter logs actively.
```bash
cd $HOME
screen -S relayer
rly start defundhub
# To Exit Screen
CTLR+A+D
# To Reenter screen
screen -r relayer
```