# Defund
**Defund** is a blockchain that allows quantitative developers, financial institutions, ETF providers, financial advisors, and individual investors to build and invest in decentralized exchange traded funds that are completely trustless and decentralized with access to tokens from 60+ blockchains and support for 1000+ assets.

Use our base index fund smart contract to create a dETF from the command line or our UI in minutes. Alternatively, build out your own smart contract based dETF in hours, with direct support for 1000's of assets from 60+ chains.

## Set Up Validator/Node On Akash

Details on setting up a Defund node and/or validator are coming soon.

## Install

You will need Golang, Rust, the Hermes IBC relayer, and the Defund Golang relayer (https://github.com/defund-labs/relayer) installed to run in dev. 

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

```bash
cd vue
npm i
npm run dev
```

## Creating A dETF

The following command is an example on creating a new dETF on Defund via the Defund binary:

```bash
# Create a new fund
# defundd tx etf create-fund [symbol] [name] [fund description] [base denom] [broker] [holdings] [rebalance period] [connection-id]
# holdings follow a comma seperated list format like so, denom:percent:poolId,denom:percent:poolId
defundd tx etf create-fund ATOM2 "ATOM Top 2" "The top 2 coins in the cosmos!" uatom gdex uatom:50:1,ibc/68A333688E5B07451F95555F8FE510E43EF9D3D44DF0909964F92081EF9BE5A7:50:2 10 connection-0 --from $KEY_NAME --keyring-backend test --home ./network/data/defund --gas auto
```

## Query dETF

```bash
defundd query etf fund ATOM2
```

All dETF's start at a price of 1 base denom. After the first investment all prices are proportially based on the underlying holdings
of that dETF.

## Query dETF Current Price

```bash
defundd query etf fund-price ATOM2
```

## Create Shares In A Fund

The following command is an example on investing in a dETF on Defund via the Defund binary:

```bash
# you must have the right portion of each asset in the fund for creating shares
defundd tx etf create-shares ATOM2
```

## Redeem Shares In A Fund

The following command is an example of redeeming an investment in a dETF on Defund via the Defund binary:

```bash
# you will receive each token that represents the funds holdings proportially to your ownership amount being redeemed
defundd tx etf redeem-shares ATOM2
```

## Planned Roadmap
* Add Market Cap Weighted to Base Index Contract
* Fractionalized NFT's
* Superfluid Staking - Fixed Income dETF's
* Support for More Chains and Tokens Until All Tokens are Supported :)

## Potential Roadmap
* Add Support for More TradFi Based Smart Contract Language Support (Java?, Python?, Julia?)
* Sub 1s Block Times for HFT?
