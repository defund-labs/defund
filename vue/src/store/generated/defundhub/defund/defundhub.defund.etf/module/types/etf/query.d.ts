import { Reader, Writer } from "protobufjs/minimal";
import { Fund } from "../etf/fund";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "defundhub.defund.etf";
export interface QueryGetFundRequest {
    index: string;
}
export interface QueryGetFundResponse {
    fund: Fund | undefined;
}
export interface QueryAllFundRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllFundResponse {
    fund: Fund[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetFundRequest: {
    encode(message: QueryGetFundRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetFundRequest;
    fromJSON(object: any): QueryGetFundRequest;
    toJSON(message: QueryGetFundRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetFundRequest>): QueryGetFundRequest;
};
export declare const QueryGetFundResponse: {
    encode(message: QueryGetFundResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetFundResponse;
    fromJSON(object: any): QueryGetFundResponse;
    toJSON(message: QueryGetFundResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetFundResponse>): QueryGetFundResponse;
};
export declare const QueryAllFundRequest: {
    encode(message: QueryAllFundRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllFundRequest;
    fromJSON(object: any): QueryAllFundRequest;
    toJSON(message: QueryAllFundRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllFundRequest>): QueryAllFundRequest;
};
export declare const QueryAllFundResponse: {
    encode(message: QueryAllFundResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllFundResponse;
    fromJSON(object: any): QueryAllFundResponse;
    toJSON(message: QueryAllFundResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllFundResponse>): QueryAllFundResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a fund by index. */
    Fund(request: QueryGetFundRequest): Promise<QueryGetFundResponse>;
    /** Queries a list of fund items. */
    FundAll(request: QueryAllFundRequest): Promise<QueryAllFundResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Fund(request: QueryGetFundRequest): Promise<QueryGetFundResponse>;
    FundAll(request: QueryAllFundRequest): Promise<QueryAllFundResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
