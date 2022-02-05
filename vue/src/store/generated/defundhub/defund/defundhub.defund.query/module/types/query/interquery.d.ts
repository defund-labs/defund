import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "defundhub.defund.query";
export interface Interquery {
    index: string;
    height: string;
    path: string;
    chainId: string;
    typeName: string;
    creator: string;
}
export declare const Interquery: {
    encode(message: Interquery, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Interquery;
    fromJSON(object: any): Interquery;
    toJSON(message: Interquery): unknown;
    fromPartial(object: DeepPartial<Interquery>): Interquery;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
