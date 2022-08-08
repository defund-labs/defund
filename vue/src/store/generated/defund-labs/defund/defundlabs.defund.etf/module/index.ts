// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateFund } from "./types/etf/tx";
import { MsgRedeem } from "./types/etf/tx";
import { MsgCreate } from "./types/etf/tx";


const types = [
  ["/defundlabs.defund.etf.MsgCreateFund", MsgCreateFund],
  ["/defundlabs.defund.etf.MsgRedeem", MsgRedeem],
  ["/defundlabs.defund.etf.MsgCreate", MsgCreate],
  
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
    msgCreateFund: (data: MsgCreateFund): EncodeObject => ({ typeUrl: "/defundlabs.defund.etf.MsgCreateFund", value: MsgCreateFund.fromPartial( data ) }),
    msgRedeem: (data: MsgRedeem): EncodeObject => ({ typeUrl: "/defundlabs.defund.etf.MsgRedeem", value: MsgRedeem.fromPartial( data ) }),
    msgCreate: (data: MsgCreate): EncodeObject => ({ typeUrl: "/defundlabs.defund.etf.MsgCreate", value: MsgCreate.fromPartial( data ) }),
    
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
