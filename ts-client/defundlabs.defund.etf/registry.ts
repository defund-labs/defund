import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreate } from "./types/etf/tx";
import { MsgCreateFund } from "./types/etf/tx";
import { MsgRedeem } from "./types/etf/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/defundlabs.defund.etf.MsgCreate", MsgCreate],
    ["/defundlabs.defund.etf.MsgCreateFund", MsgCreateFund],
    ["/defundlabs.defund.etf.MsgRedeem", MsgRedeem],
    
];

export { msgTypes }