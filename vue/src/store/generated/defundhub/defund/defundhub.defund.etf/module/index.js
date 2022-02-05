// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateFund } from "./types/etf/tx";
import { MsgInvest } from "./types/etf/tx";
import { MsgUninvest } from "./types/etf/tx";
import { MsgCreateFund } from "./types/etf/tx";
const types = [
    ["/defundhub.defund.etf.MsgUpdateFund", MsgUpdateFund],
    ["/defundhub.defund.etf.MsgInvest", MsgInvest],
    ["/defundhub.defund.etf.MsgUninvest", MsgUninvest],
    ["/defundhub.defund.etf.MsgCreateFund", MsgCreateFund],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgUpdateFund: (data) => ({ typeUrl: "/defundhub.defund.etf.MsgUpdateFund", value: data }),
        msgInvest: (data) => ({ typeUrl: "/defundhub.defund.etf.MsgInvest", value: data }),
        msgUninvest: (data) => ({ typeUrl: "/defundhub.defund.etf.MsgUninvest", value: data }),
        msgCreateFund: (data) => ({ typeUrl: "/defundhub.defund.etf.MsgCreateFund", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
