/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Fund, FundPrice } from "./fund";

export const protobufPackage = "defundlabs.defund.etf";

export interface QueryGetFundRequest {
  symbol: string;
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

export interface QueryFundPriceRequest {
  symbol: string;
}

export interface QueryFundPriceResponse {
  price: FundPrice | undefined;
}

function createBaseQueryGetFundRequest(): QueryGetFundRequest {
  return { symbol: "" };
}

export const QueryGetFundRequest = {
  encode(message: QueryGetFundRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetFundRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetFundRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbol = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFundRequest {
    return { symbol: isSet(object.symbol) ? String(object.symbol) : "" };
  },

  toJSON(message: QueryGetFundRequest): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetFundRequest>, I>>(object: I): QueryGetFundRequest {
    const message = createBaseQueryGetFundRequest();
    message.symbol = object.symbol ?? "";
    return message;
  },
};

function createBaseQueryGetFundResponse(): QueryGetFundResponse {
  return { fund: undefined };
}

export const QueryGetFundResponse = {
  encode(message: QueryGetFundResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fund !== undefined) {
      Fund.encode(message.fund, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetFundResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetFundResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fund = Fund.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFundResponse {
    return { fund: isSet(object.fund) ? Fund.fromJSON(object.fund) : undefined };
  },

  toJSON(message: QueryGetFundResponse): unknown {
    const obj: any = {};
    message.fund !== undefined && (obj.fund = message.fund ? Fund.toJSON(message.fund) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetFundResponse>, I>>(object: I): QueryGetFundResponse {
    const message = createBaseQueryGetFundResponse();
    message.fund = (object.fund !== undefined && object.fund !== null) ? Fund.fromPartial(object.fund) : undefined;
    return message;
  },
};

function createBaseQueryAllFundRequest(): QueryAllFundRequest {
  return { pagination: undefined };
}

export const QueryAllFundRequest = {
  encode(message: QueryAllFundRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllFundRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllFundRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFundRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllFundRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllFundRequest>, I>>(object: I): QueryAllFundRequest {
    const message = createBaseQueryAllFundRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllFundResponse(): QueryAllFundResponse {
  return { fund: [], pagination: undefined };
}

export const QueryAllFundResponse = {
  encode(message: QueryAllFundResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.fund) {
      Fund.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllFundResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllFundResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fund.push(Fund.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFundResponse {
    return {
      fund: Array.isArray(object?.fund) ? object.fund.map((e: any) => Fund.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllFundResponse): unknown {
    const obj: any = {};
    if (message.fund) {
      obj.fund = message.fund.map((e) => e ? Fund.toJSON(e) : undefined);
    } else {
      obj.fund = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllFundResponse>, I>>(object: I): QueryAllFundResponse {
    const message = createBaseQueryAllFundResponse();
    message.fund = object.fund?.map((e) => Fund.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryFundPriceRequest(): QueryFundPriceRequest {
  return { symbol: "" };
}

export const QueryFundPriceRequest = {
  encode(message: QueryFundPriceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryFundPriceRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryFundPriceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbol = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFundPriceRequest {
    return { symbol: isSet(object.symbol) ? String(object.symbol) : "" };
  },

  toJSON(message: QueryFundPriceRequest): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryFundPriceRequest>, I>>(object: I): QueryFundPriceRequest {
    const message = createBaseQueryFundPriceRequest();
    message.symbol = object.symbol ?? "";
    return message;
  },
};

function createBaseQueryFundPriceResponse(): QueryFundPriceResponse {
  return { price: undefined };
}

export const QueryFundPriceResponse = {
  encode(message: QueryFundPriceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.price !== undefined) {
      FundPrice.encode(message.price, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryFundPriceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryFundPriceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.price = FundPrice.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFundPriceResponse {
    return { price: isSet(object.price) ? FundPrice.fromJSON(object.price) : undefined };
  },

  toJSON(message: QueryFundPriceResponse): unknown {
    const obj: any = {};
    message.price !== undefined && (obj.price = message.price ? FundPrice.toJSON(message.price) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryFundPriceResponse>, I>>(object: I): QueryFundPriceResponse {
    const message = createBaseQueryFundPriceResponse();
    message.price = (object.price !== undefined && object.price !== null)
      ? FundPrice.fromPartial(object.price)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a fund by symbol. */
  Fund(request: QueryGetFundRequest): Promise<QueryGetFundResponse>;
  /** Queries a list of fund items. */
  FundAll(request: QueryAllFundRequest): Promise<QueryAllFundResponse>;
  /** Queries a list of fundPrice items. */
  FundPrice(request: QueryFundPriceRequest): Promise<QueryFundPriceResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Fund = this.Fund.bind(this);
    this.FundAll = this.FundAll.bind(this);
    this.FundPrice = this.FundPrice.bind(this);
  }
  Fund(request: QueryGetFundRequest): Promise<QueryGetFundResponse> {
    const data = QueryGetFundRequest.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.etf.Query", "Fund", data);
    return promise.then((data) => QueryGetFundResponse.decode(new _m0.Reader(data)));
  }

  FundAll(request: QueryAllFundRequest): Promise<QueryAllFundResponse> {
    const data = QueryAllFundRequest.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.etf.Query", "FundAll", data);
    return promise.then((data) => QueryAllFundResponse.decode(new _m0.Reader(data)));
  }

  FundPrice(request: QueryFundPriceRequest): Promise<QueryFundPriceResponse> {
    const data = QueryFundPriceRequest.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.etf.Query", "FundPrice", data);
    return promise.then((data) => QueryFundPriceResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
