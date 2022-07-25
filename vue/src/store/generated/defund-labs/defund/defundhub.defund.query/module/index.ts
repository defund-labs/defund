// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateInterquery } from "./types/query/tx";
import { MsgCreateInterqueryResult } from "./types/query/tx";
import { MsgCreateInterqueryTimeout } from "./types/query/tx";


const types = [
  ["/defundhub.defund.query.MsgCreateInterquery", MsgCreateInterquery],
  ["/defundhub.defund.query.MsgCreateInterqueryResult", MsgCreateInterqueryResult],
  ["/defundhub.defund.query.MsgCreateInterqueryTimeout", MsgCreateInterqueryTimeout],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateInterquery: (data: MsgCreateInterquery): EncodeObject => ({ typeUrl: "/defundhub.defund.query.MsgCreateInterquery", value: MsgCreateInterquery.fromPartial( data ) }),
    msgCreateInterqueryResult: (data: MsgCreateInterqueryResult): EncodeObject => ({ typeUrl: "/defundhub.defund.query.MsgCreateInterqueryResult", value: MsgCreateInterqueryResult.fromPartial( data ) }),
    msgCreateInterqueryTimeout: (data: MsgCreateInterqueryTimeout): EncodeObject => ({ typeUrl: "/defundhub.defund.query.MsgCreateInterqueryTimeout", value: MsgCreateInterqueryTimeout.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
