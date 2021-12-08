import { Reader, Writer } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";
export declare const protobufPackage = "defundhub.defund.etf";
export interface MsgCreateFund {
    creator: string;
    id: string;
    address: string;
    symbol: string;
    name: string;
    description: string;
    shares: Coin | undefined;
}
export interface MsgCreateFundResponse {
}
export interface MsgUpdateFund {
    creator: string;
    id: string;
    address: string;
    symbol: string;
    name: string;
    description: string;
    shares: Coin | undefined;
}
export interface MsgUpdateFundResponse {
}
export declare const MsgCreateFund: {
    encode(message: MsgCreateFund, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateFund;
    fromJSON(object: any): MsgCreateFund;
    toJSON(message: MsgCreateFund): unknown;
    fromPartial(object: DeepPartial<MsgCreateFund>): MsgCreateFund;
};
export declare const MsgCreateFundResponse: {
    encode(_: MsgCreateFundResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateFundResponse;
    fromJSON(_: any): MsgCreateFundResponse;
    toJSON(_: MsgCreateFundResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateFundResponse>): MsgCreateFundResponse;
};
export declare const MsgUpdateFund: {
    encode(message: MsgUpdateFund, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateFund;
    fromJSON(object: any): MsgUpdateFund;
    toJSON(message: MsgUpdateFund): unknown;
    fromPartial(object: DeepPartial<MsgUpdateFund>): MsgUpdateFund;
};
export declare const MsgUpdateFundResponse: {
    encode(_: MsgUpdateFundResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateFundResponse;
    fromJSON(_: any): MsgUpdateFundResponse;
    toJSON(_: MsgUpdateFundResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateFundResponse>): MsgUpdateFundResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateFund(request: MsgCreateFund): Promise<MsgCreateFundResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    UpdateFund(request: MsgUpdateFund): Promise<MsgUpdateFundResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateFund(request: MsgCreateFund): Promise<MsgCreateFundResponse>;
    UpdateFund(request: MsgUpdateFund): Promise<MsgUpdateFundResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
