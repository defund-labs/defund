import { Reader, Writer } from "protobufjs/minimal";
import { Interquery } from "../query/interquery";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "defundhub.defund.query";
export interface QueryGetInterqueryRequest {
    index: string;
}
export interface QueryGetInterqueryResponse {
    interquery: Interquery | undefined;
}
export interface QueryAllInterqueryRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllInterqueryResponse {
    interquery: Interquery[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetInterqueryRequest: {
    encode(message: QueryGetInterqueryRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetInterqueryRequest;
    fromJSON(object: any): QueryGetInterqueryRequest;
    toJSON(message: QueryGetInterqueryRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetInterqueryRequest>): QueryGetInterqueryRequest;
};
export declare const QueryGetInterqueryResponse: {
    encode(message: QueryGetInterqueryResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetInterqueryResponse;
    fromJSON(object: any): QueryGetInterqueryResponse;
    toJSON(message: QueryGetInterqueryResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetInterqueryResponse>): QueryGetInterqueryResponse;
};
export declare const QueryAllInterqueryRequest: {
    encode(message: QueryAllInterqueryRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllInterqueryRequest;
    fromJSON(object: any): QueryAllInterqueryRequest;
    toJSON(message: QueryAllInterqueryRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllInterqueryRequest>): QueryAllInterqueryRequest;
};
export declare const QueryAllInterqueryResponse: {
    encode(message: QueryAllInterqueryResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllInterqueryResponse;
    fromJSON(object: any): QueryAllInterqueryResponse;
    toJSON(message: QueryAllInterqueryResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllInterqueryResponse>): QueryAllInterqueryResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a interquery by index. */
    Interquery(request: QueryGetInterqueryRequest): Promise<QueryGetInterqueryResponse>;
    /** Queries a list of interquery items. */
    InterqueryAll(request: QueryAllInterqueryRequest): Promise<QueryAllInterqueryResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Interquery(request: QueryGetInterqueryRequest): Promise<QueryGetInterqueryResponse>;
    InterqueryAll(request: QueryAllInterqueryRequest): Promise<QueryAllInterqueryResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
