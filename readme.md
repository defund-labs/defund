# Defund
**Defund** is a blockchain that allows the investment in and creation of diversified, autonomous exchange traded funds.

## Install

```
git clone https://github.com/defund-labs/defund

cd defund

make install
```

`make install` installs additional components including, the hermes relayer by Informal, the Osmosis binary, and the Cosmos Hub binary (gaia).

## Getting Started

```
make init

make start
```

`make init` initializes the Defund chain as well as the hermes relayer and a local implementation of Osmosis and Gaia (Cosmos Hub).

`make start` first starts a local version of Defund as well as the hermes relayer with a local version of Gaia (Cosmos Hub) and Osmosis.

## Frontend

To start the frontend in dev please run the following commands in another terminal:

```
cd frontend
npm install
npm run dev
```

## Creating A Fund

The following command is an example on creating a new dETF on Defund via the Defund binary:

```

```

## Invest In A Fund

The following command is an example on investing in a dETF on Defund via the Defund binary:

```

```

## Uninvest In A Fund

The following command is an example of redeeming an investment in a dETF on Defund via the Defund binary:

```

```

## Planned Roadmap
`Stage 1`

Specify crypto assets and weights on creation and autorebalance

`Stage 2` 

Add smart contracts allowing for programmatic trading of underlying assets of funds.

`Potential Other Additions`
* Superfluid Staking
* Add the ability to add instructions that power the trading of funds (if this then thats)
