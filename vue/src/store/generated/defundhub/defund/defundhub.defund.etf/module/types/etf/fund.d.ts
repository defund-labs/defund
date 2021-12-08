import { Coin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "defundhub.defund.etf";
export interface Fund {
    id: string;
    address: string;
    symbol: string;
    name: string;
    description: string;
    shares: Coin | undefined;
    creator: string;
}
export declare const Fund: {
    encode(message: Fund, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Fund;
    fromJSON(object: any): Fund;
    toJSON(message: Fund): unknown;
    fromPartial(object: DeepPartial<Fund>): Fund;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
