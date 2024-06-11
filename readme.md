![alt text](https://defund.app/images/Defund-6-p-500.png)

## What is DeFund

**DeFund** is a performance-optimized blockchain that fuels a rich ecosystem of DeFi products and applications, granting builders and investors access to unmatched cross-chain liquidity, minimized or maximized risk, and the full freedom to execute custom trading strategies across hundreds of blockchains and thousands of assets.


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

### Current + Planned Roadmap
* Structured Products - *current*
* Options + Option Vaults - *next to launch*
* Hybrid DEX - *next to launch*
* Revamped tokenomics (distribution + utility/function) - *coming soon*
* Non-standard/Alternative Asset options - *underway*
* Hedged Liquidity Pools - *underway*
* Auto-trade/super-trade functionality - *planned*
* Cross-VM outposts - *planned*
* Support for More Chains and Tokens Until All Tokens are Supported :)

### Potential Roadmap
* Add Support for More TradFi Based Smart Contract Language Support (Java?, Python?, Julia?)

## CLI Docs
- [Create Fund Command](./x/etf/client/docs/create-etf.md)
- [Create Shares Command](./x/etf/client/docs/create-shares.md)
- [Redeem Shares Command](./x/etf/client/docs/redeem-shares.md)
- [Cosmwasm dETF Commands](./x/etf/client/docs/create-wasm-etf.md)
