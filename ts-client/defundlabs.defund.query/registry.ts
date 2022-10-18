import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateInterquery } from "./types/query/tx";
import { MsgCreateInterqueryResult } from "./types/query/tx";
import { MsgCreateInterqueryTimeout } from "./types/query/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/defundlabs.defund.query.MsgCreateInterquery", MsgCreateInterquery],
    ["/defundlabs.defund.query.MsgCreateInterqueryResult", MsgCreateInterqueryResult],
    ["/defundlabs.defund.query.MsgCreateInterqueryTimeout", MsgCreateInterqueryTimeout],
    
];

export { msgTypes }