<p align="center">
  <img src="https://github.com/defund-labs/defund/blob/d80730b57ed80a1f75c5d9990faa16efa7477e92/Defund-BrandAssets/Banners/DeFund-White.png" alt="DeFund Finance"/>
</p>

**DeFund Finance** is a sovereign rollup built on Celestia using Rollkit, designed to create a rich DeFi ecosystem that serves as a hub for all modular liquidity. By leveraging the Celestia DA (Data Availability) layer, DeFund Finance aims to provide a comprehensive suite of DeFi products and services tailored to the modular blockchain ecosystem.

## Key Features

- **Spot Exchange**: DeFund Finance offers a state-of-the-art decentralized exchange that combines the best features of an orderbook and an automated market maker (AMM) model, providing deep liquidity, efficient price discovery, and minimal slippage for traders.

- **Hedged Liquidity Pools**: DeFund Finance's Hedged Liquidity Pools are an innovative solution to the problem of impermanent loss faced by liquidity providers in traditional AMMs. These pools work in conjunction with the Spot Exchange, serving as the main source of liquidity for the exchange while minimizing impermanent loss through an automated hedging mechanism.

- **Option Vaults**: Option Vaults bring the power of customizable risk management to the DeFi space. Users can mint customized options using a unique Dutch auction mechanism, providing a novel approach to hedging, yield generation, and market speculation.

- **Asset Abstraction**: DeFund Finance takes asset interoperability and accessibility to the next level through the integration of Calypso and the Cosmos extension for Metamask. These integrations enable users to perform any action with any asset using the most widely adopted wallet, MetaMask.

## Architecture

DeFund Finance is built on a robust and scalable architecture that leverages the power of Rollkit and Celestia to create a high-performance, secure, and flexible platform for decentralized finance.

- **Rollkit + Celestia**: The combination of Rollkit and Celestia enables DeFund Finance to create a sovereign rollup that benefits from the security and scalability of Celestia's DA layer while maintaining the flexibility and customization offered by Rollkit.

- **Asset Abstraction**: DeFund Finance integrates Calypso and the Cosmos Metamask Snap to enable seamless interoperability between different blockchain networks and assets. Users can perform any action with any token, regardless of its native blockchain, using the familiar Metamask wallet.

## Getting Started

### Prerequisites

- Golang
- Kurtosis 1.0.0 (https://docs.kurtosis.com/install)
- Hermes IBC relayer (https://hermes.informal.systems/quick-start/installation.html)

## Install

You will need Golang and Kurtosis 1.0.0 installed. 

```
git clone https://github.com/defund-labs/defund

cd defund

# If you want the CLI installed
make install
```

## Getting Started in Dev/Local Mode

```bash
# Make sure you have kurtosis 1.0.0 installed and running
kurtosis run .
```

## Roadmap

- [x] Hybrid DEX
- [ ] Options + Option Vaults
- [ ] Hedged Liquidity Pools
- [ ] Rollkit Interoperability
- [ ] Cross-VM outposts
