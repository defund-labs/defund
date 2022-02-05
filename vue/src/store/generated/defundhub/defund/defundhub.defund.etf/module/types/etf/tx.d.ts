import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "defundhub.defund.etf";
export interface MsgCreateFund {
    creator: string;
    symbol: string;
    name: string;
    description: string;
}
export interface MsgCreateFundResponse {
}
export interface MsgUpdateFund {
    creator: string;
    id: string;
    name: string;
    description: string;
}
export interface MsgUpdateFundResponse {
}
export interface MsgInvest {
    creator: string;
    fund: string;
    amount: string;
}
export interface MsgInvestResponse {
}
export interface MsgUninvest {
    creator: string;
    fund: string;
    amount: string;
}
export interface MsgUninvestResponse {
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
export declare const MsgInvest: {
    encode(message: MsgInvest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgInvest;
    fromJSON(object: any): MsgInvest;
    toJSON(message: MsgInvest): unknown;
    fromPartial(object: DeepPartial<MsgInvest>): MsgInvest;
};
export declare const MsgInvestResponse: {
    encode(_: MsgInvestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgInvestResponse;
    fromJSON(_: any): MsgInvestResponse;
    toJSON(_: MsgInvestResponse): unknown;
    fromPartial(_: DeepPartial<MsgInvestResponse>): MsgInvestResponse;
};
export declare const MsgUninvest: {
    encode(message: MsgUninvest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUninvest;
    fromJSON(object: any): MsgUninvest;
    toJSON(message: MsgUninvest): unknown;
    fromPartial(object: DeepPartial<MsgUninvest>): MsgUninvest;
};
export declare const MsgUninvestResponse: {
    encode(_: MsgUninvestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUninvestResponse;
    fromJSON(_: any): MsgUninvestResponse;
    toJSON(_: MsgUninvestResponse): unknown;
    fromPartial(_: DeepPartial<MsgUninvestResponse>): MsgUninvestResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateFund(request: MsgCreateFund): Promise<MsgCreateFundResponse>;
    UpdateFund(request: MsgUpdateFund): Promise<MsgUpdateFundResponse>;
    Invest(request: MsgInvest): Promise<MsgInvestResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    Uninvest(request: MsgUninvest): Promise<MsgUninvestResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateFund(request: MsgCreateFund): Promise<MsgCreateFundResponse>;
    UpdateFund(request: MsgUpdateFund): Promise<MsgUpdateFundResponse>;
    Invest(request: MsgInvest): Promise<MsgInvestResponse>;
    Uninvest(request: MsgUninvest): Promise<MsgUninvestResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
