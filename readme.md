# Defund
**Defund** is a blockchain that allows the investment in and creation of diversified, autonomous exchange traded funds.

## Set Up Validator/Node On Akash

Details on setting up a Defund node and/or validator are coming soon.

## Install

You will need Golang, Rust, the Hermes IBC relayer, and the Defund Golang relayer (https://github.com/defund-labs/relayer) installed to run in dev. 

```
git clone https://github.com/defund-labs/defund

cd defund

make install
```

## Getting Started in Dev Mode

```bash
make init

make create-conn

# Wait for the connection to be acknowledged then
# start the relayer for interquerying
make start-rly
```

## Frontend

```bash
#Coming Soon!
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

## Query dETF Historical Price

```bash
defundd query etf fund-prices ATOM2
```

## Invest In A Fund

The following command is an example on investing in a dETF on Defund via the Defund binary:

```bash
#Coming soon!
```

## Uninvest In A Fund

The following command is an example of redeeming an investment in a dETF on Defund via the Defund binary:

```bash
#Coming soon!
```

## Planned Roadmap
`Stage 1`

Specify crypto assets and weights on creation and autorebalance

`Stage 2` 

Add smart contracts allowing for programmatic trading of underlying assets of funds.

`Potential Other Additions`
* Superfluid Staking
* Add the ability to add instructions that power the trading of funds (if this then thats)
