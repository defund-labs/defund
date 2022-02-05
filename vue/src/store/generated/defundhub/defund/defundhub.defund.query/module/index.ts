// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateInterquery } from "./types/query/tx";
import { MsgUpdateInterquery } from "./types/query/tx";
import { MsgDeleteInterquery } from "./types/query/tx";


const types = [
  ["/defundhub.defund.query.MsgCreateInterquery", MsgCreateInterquery],
  ["/defundhub.defund.query.MsgUpdateInterquery", MsgUpdateInterquery],
  ["/defundhub.defund.query.MsgDeleteInterquery", MsgDeleteInterquery],
  
];
export const MissingWalletError = new Error("wallet is required");

const registry = new Registry(<any>types);

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

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateInterquery: (data: MsgCreateInterquery): EncodeObject => ({ typeUrl: "/defundhub.defund.query.MsgCreateInterquery", value: data }),
    msgUpdateInterquery: (data: MsgUpdateInterquery): EncodeObject => ({ typeUrl: "/defundhub.defund.query.MsgUpdateInterquery", value: data }),
    msgDeleteInterquery: (data: MsgDeleteInterquery): EncodeObject => ({ typeUrl: "/defundhub.defund.query.MsgDeleteInterquery", value: data }),
    
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
