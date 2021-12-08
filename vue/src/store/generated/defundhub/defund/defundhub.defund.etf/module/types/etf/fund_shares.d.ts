import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "defundhub.defund.etf";
export interface FundShares {
    amount: string;
    denom: string;
}
export declare const FundShares: {
    encode(message: FundShares, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): FundShares;
    fromJSON(object: any): FundShares;
    toJSON(message: FundShares): unknown;
    fromPartial(object: DeepPartial<FundShares>): FundShares;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
