import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddConnectionBroker } from "./types/broker/tx";
import { MsgAddLiquiditySource } from "./types/broker/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/defundlabs.defund.broker.MsgAddConnectionBroker", MsgAddConnectionBroker],
    ["/defundlabs.defund.broker.MsgAddLiquiditySource", MsgAddLiquiditySource],
    
];

export { msgTypes }