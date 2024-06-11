<p align="center">
  <img src="https://cdn.prod.website-files.com/61ff2317a4bdb55b6494bb53/61ff26c9b9a7293f886ca79e_Defund%20(6)-p-500.png" alt="DeFund Finance"/>
</p>

**DeFund Finance** is a sovereign rollup built on Celestia using Rollkit, designed to create a rich DeFi ecosystem that serves as a hub for all modular liquidity. By leveraging the Celestia DA (Data Availability) layer, DeFund Finance aims to provide a comprehensive suite of DeFi products and services tailored to the modular blockchain ecosystem.

## Key Features

- **Spot Exchange**: DeFund Finance offers a state-of-the-art decentralized exchange that combines the best features of an orderbook and an automated market maker (AMM) model, providing deep liquidity, efficient price discovery, and minimal slippage for traders.

- **LP Vaults**: DeFund Finance's LP Vaults are an innovative solution to the problem of impermanent loss faced by liquidity providers in traditional AMMs. These vaults work in conjunction with the Spot Exchange, serving as the liquidity pools for the exchange while minimizing impermanent loss through an automated hedging mechanism.

- **Option Vaults**: Option Vaults bring the power of customizable risk management to the DeFi space. Users can mint customized options using a unique Dutch auction mechanism, providing a novel approach to hedging, yield generation, and market speculation.

- **Asset Abstraction**: DeFund Finance takes asset interoperability and accessibility to the next level through the integration of Calypso and the Cosmos Metamask Snap. These integrations enable users to perform any action with any asset using the most widely adopted wallet, Metamask.

## Architecture

DeFund Finance is built on a robust and scalable architecture that leverages the power of Rollkit and Celestia to create a high-performance, secure, and flexible platform for decentralized finance.

- **Rollkit + Celestia**: The combination of Rollkit and Celestia enables DeFund Finance to create a sovereign rollup that benefits from the security and scalability of Celestia's DA layer while maintaining the flexibility and customization offered by Rollkit.

- **Asset Abstraction**: DeFund Finance integrates Calypso and the Cosmos Metamask Snap to enable seamless interoperability between different blockchain networks and assets. Users can perform any action with any token, regardless of its native blockchain, using the familiar Metamask wallet.

## Getting Started

### Prerequisites

- Golang
- Rust
- Hermes IBC relayer
- DeFund Golang relayer (https://github.com/defund-labs/relayer)

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

## Roadmap

- [x] Structured Products
- [ ] Options + Option Vaults
- [ ] Hybrid DEX
- [ ] Revamped tokenomics (distribution + utility/function)
- [ ] Non-standard/Alternative Asset options
- [ ] Hedged Liquidity Pools
- [ ] Auto-trade/super-trade functionality
- [ ] Cross-VM outposts
- [ ] Support for more chains and tokens

## CLI Documentation (Deprecated)

- [Create Fund Command](docs/cli/create-fund.md)
- [Create Shares Command](docs/cli/create-shares.md)
- [Redeem Shares Command](docs/cli/redeem-shares.md)
- [Cosmwasm dETF Commands](docs/cli/cosmwasm-detf.md)