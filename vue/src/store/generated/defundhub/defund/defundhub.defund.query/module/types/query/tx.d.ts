import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "defundhub.defund.query";
export interface MsgCreateInterquery {
    creator: string;
    index: string;
    height: string;
    path: string;
    chainId: string;
    typeName: string;
}
export interface MsgCreateInterqueryResponse {
}
export interface MsgUpdateInterquery {
    creator: string;
    index: string;
    height: string;
    path: string;
    chainId: string;
    typeName: string;
}
export interface MsgUpdateInterqueryResponse {
}
export interface MsgDeleteInterquery {
    creator: string;
    index: string;
}
export interface MsgDeleteInterqueryResponse {
}
export declare const MsgCreateInterquery: {
    encode(message: MsgCreateInterquery, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateInterquery;
    fromJSON(object: any): MsgCreateInterquery;
    toJSON(message: MsgCreateInterquery): unknown;
    fromPartial(object: DeepPartial<MsgCreateInterquery>): MsgCreateInterquery;
};
export declare const MsgCreateInterqueryResponse: {
    encode(_: MsgCreateInterqueryResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateInterqueryResponse;
    fromJSON(_: any): MsgCreateInterqueryResponse;
    toJSON(_: MsgCreateInterqueryResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateInterqueryResponse>): MsgCreateInterqueryResponse;
};
export declare const MsgUpdateInterquery: {
    encode(message: MsgUpdateInterquery, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateInterquery;
    fromJSON(object: any): MsgUpdateInterquery;
    toJSON(message: MsgUpdateInterquery): unknown;
    fromPartial(object: DeepPartial<MsgUpdateInterquery>): MsgUpdateInterquery;
};
export declare const MsgUpdateInterqueryResponse: {
    encode(_: MsgUpdateInterqueryResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateInterqueryResponse;
    fromJSON(_: any): MsgUpdateInterqueryResponse;
    toJSON(_: MsgUpdateInterqueryResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateInterqueryResponse>): MsgUpdateInterqueryResponse;
};
export declare const MsgDeleteInterquery: {
    encode(message: MsgDeleteInterquery, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteInterquery;
    fromJSON(object: any): MsgDeleteInterquery;
    toJSON(message: MsgDeleteInterquery): unknown;
    fromPartial(object: DeepPartial<MsgDeleteInterquery>): MsgDeleteInterquery;
};
export declare const MsgDeleteInterqueryResponse: {
    encode(_: MsgDeleteInterqueryResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteInterqueryResponse;
    fromJSON(_: any): MsgDeleteInterqueryResponse;
    toJSON(_: MsgDeleteInterqueryResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteInterqueryResponse>): MsgDeleteInterqueryResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateInterquery(request: MsgCreateInterquery): Promise<MsgCreateInterqueryResponse>;
    UpdateInterquery(request: MsgUpdateInterquery): Promise<MsgUpdateInterqueryResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteInterquery(request: MsgDeleteInterquery): Promise<MsgDeleteInterqueryResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateInterquery(request: MsgCreateInterquery): Promise<MsgCreateInterqueryResponse>;
    UpdateInterquery(request: MsgUpdateInterquery): Promise<MsgUpdateInterqueryResponse>;
    DeleteInterquery(request: MsgDeleteInterquery): Promise<MsgDeleteInterqueryResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
