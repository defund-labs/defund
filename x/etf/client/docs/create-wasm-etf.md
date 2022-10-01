## Create A WASM dETF

The below command is an example of creating a Cosmwasm based dETF on Defund via the Defund binary.

```bash
# store your dETF contract on chain
defundd tx wasm store [dETF contract path]

# instantiate your dETF
defundd tx wasm instantiate [code_id_int64] [json_encoded_init_args] --label [text] --admin [address,optional] --amount [coins,optional]

# execute a command on your dETF contract
defundd tx wasm execute [contract_addr_bech32] [json_encoded_send_args] --amount [coins,optional]
```