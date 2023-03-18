![alt text](https://defund.app/images/Defund-6-p-500.png)

**DeFund** is a blockchain that allows quantitative developers, financial institutions, ETF providers, financial advisors, and individual investors to build and invest in decentralized exchange traded funds that are completely trustless and decentralized with access to tokens from 60+ blockchains and support for 1000+ assets.

Use our base index fund smart contract to create a dETF from the command line or our UI in minutes. Alternatively, build out your own smart contract based dETF in hours, with direct support for 1000's of assets from 60+ chains.

## Set Up Validator/Node On Akash

Use the following [deploy.yaml](./testnet/private/deploy.yaml) as the yaml file for launching a DeFund node in Akash on testnet. We suggest using [Akashlytics](https://cloudmos.io/cloud-deploy).

## Install

You will need Golang, Rust, the Hermes IBC relayer, and the DeFund Golang relayer (https://github.com/defund-labs/relayer) installed to run in dev. 

```
git clone https://github.com/defund-labs/defund

cd defund

make install
```

## Getting Started in Dev/Local Mode

```bash
make init

make create-conn

# Wait for the connection to be acknowledged then edit ./network/relayer/config.yaml by changing the connection and client ids to match what your relayer returned
make init-rly

# start the relayer for interchain querying
make start-rly

# start hermes for ICA tx's
make start-hermes
```

## Bash Testing Framework
```bash
# test all main commands
make test-all

# test the create fund command
make test-create-fund

# test the create shares command
make test-create-shares

# test the redeem shares command
make test-redeem-shares
```

## Frontend
The following can be run to get started in dev on the frontend.
```bash
cd $HOME/defund/vue
npm i
npm run dev
```

### Planned Roadmap
* Fractionalized NFT's
* Perps
* Options
* Support for More Chains and Tokens Until All Tokens are Supported :)

### Potential Roadmap
* Add Support for More TradFi Based Smart Contract Language Support (Java?, Python?, Julia?)

## CLI Docs
- [Create Fund Command](./x/etf/client/docs/create-etf.md)
- [Create Shares Command](./x/etf/client/docs/create-shares.md)
- [Redeem Shares Command](./x/etf/client/docs/redeem-shares.md)
- [Cosmwasm dETF Commands](./x/etf/client/docs/create-wasm-etf.md)
