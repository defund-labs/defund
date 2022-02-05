// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
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
        msgCreateInterquery: (data) => ({ typeUrl: "/defundhub.defund.query.MsgCreateInterquery", value: data }),
        msgUpdateInterquery: (data) => ({ typeUrl: "/defundhub.defund.query.MsgUpdateInterquery", value: data }),
        msgDeleteInterquery: (data) => ({ typeUrl: "/defundhub.defund.query.MsgDeleteInterquery", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
