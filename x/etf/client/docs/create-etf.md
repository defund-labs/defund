## Create A dETF

The below command is an example of creating a standard dETF on Defund via the Defund binary. Please see [link wasm command] to see how to create a Cosmwasm smart contract based fund (what we call a dFund).

Holdings follow a comma seperated list as a string like so `denom:percent:poolId:type,denom:percent:poolId:type` with no spaces.

```bash
defundd tx etf create-fund COSM2 "The Cosmos Top 2" "The top 2 coins in the cosmos!" uosmo uosmo:50:1:osmosis:spot,ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2:50:2:osmosis:spot 10 5000000 --from keyname
```