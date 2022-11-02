import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRedeem } from "./types/etf/tx";
import { MsgCreateFund } from "./types/etf/tx";
import { MsgCreate } from "./types/etf/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/defundlabs.defund.etf.MsgRedeem", MsgRedeem],
    ["/defundlabs.defund.etf.MsgCreateFund", MsgCreateFund],
    ["/defundlabs.defund.etf.MsgCreate", MsgCreate],
    
];

export { msgTypes }