import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddLiquiditySource } from "./types/broker/tx";
import { MsgAddConnectionBroker } from "./types/broker/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/defundlabs.defund.broker.MsgAddLiquiditySource", MsgAddLiquiditySource],
    ["/defundlabs.defund.broker.MsgAddConnectionBroker", MsgAddConnectionBroker],
    
];

export { msgTypes }