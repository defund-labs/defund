## Create A WASM dETF

The below command is an example of creating a Cosmwasm based dETF on Defund via the Defund binary.

```bash
# store your dETF contract on chain
defundd tx wasm store [dETF contract path]

# create your dETF
defundd tx etf create-fund WASM3 "The Top 3 Coins via a Wasm Contract" "The top 3 IBC coins in the cosmos via a Wasm contract!" axlUSDC [wasm_code_id_int64] 1 5000000 --from keyname --active

# execute a command on your dETF contract
defundd tx wasm execute [contract_addr_bech32] [json_encoded_send_args] --amount [coins,optional]
```

At each rebalance height (set to `1` above), Defund will automatically run your contract through the `Runner` entrypoint in your contract. It's that easy!