/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Fund, FundPrice } from "../etf/fund";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

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

const baseQueryGetFundRequest: object = { symbol: "" };

export const QueryGetFundRequest = {
  encode(
    message: QueryGetFundRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetFundRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetFundRequest } as QueryGetFundRequest;
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
    const message = { ...baseQueryGetFundRequest } as QueryGetFundRequest;
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = String(object.symbol);
    } else {
      message.symbol = "";
    }
    return message;
  },

  toJSON(message: QueryGetFundRequest): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetFundRequest>): QueryGetFundRequest {
    const message = { ...baseQueryGetFundRequest } as QueryGetFundRequest;
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = object.symbol;
    } else {
      message.symbol = "";
    }
    return message;
  },
};

const baseQueryGetFundResponse: object = {};

export const QueryGetFundResponse = {
  encode(
    message: QueryGetFundResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.fund !== undefined) {
      Fund.encode(message.fund, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetFundResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetFundResponse } as QueryGetFundResponse;
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
    const message = { ...baseQueryGetFundResponse } as QueryGetFundResponse;
    if (object.fund !== undefined && object.fund !== null) {
      message.fund = Fund.fromJSON(object.fund);
    } else {
      message.fund = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetFundResponse): unknown {
    const obj: any = {};
    message.fund !== undefined &&
      (obj.fund = message.fund ? Fund.toJSON(message.fund) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetFundResponse>): QueryGetFundResponse {
    const message = { ...baseQueryGetFundResponse } as QueryGetFundResponse;
    if (object.fund !== undefined && object.fund !== null) {
      message.fund = Fund.fromPartial(object.fund);
    } else {
      message.fund = undefined;
    }
    return message;
  },
};

const baseQueryAllFundRequest: object = {};

export const QueryAllFundRequest = {
  encode(
    message: QueryAllFundRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllFundRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllFundRequest } as QueryAllFundRequest;
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
    const message = { ...baseQueryAllFundRequest } as QueryAllFundRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFundRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllFundRequest>): QueryAllFundRequest {
    const message = { ...baseQueryAllFundRequest } as QueryAllFundRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllFundResponse: object = {};

export const QueryAllFundResponse = {
  encode(
    message: QueryAllFundResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.fund) {
      Fund.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllFundResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllFundResponse } as QueryAllFundResponse;
    message.fund = [];
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
    const message = { ...baseQueryAllFundResponse } as QueryAllFundResponse;
    message.fund = [];
    if (object.fund !== undefined && object.fund !== null) {
      for (const e of object.fund) {
        message.fund.push(Fund.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFundResponse): unknown {
    const obj: any = {};
    if (message.fund) {
      obj.fund = message.fund.map((e) => (e ? Fund.toJSON(e) : undefined));
    } else {
      obj.fund = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllFundResponse>): QueryAllFundResponse {
    const message = { ...baseQueryAllFundResponse } as QueryAllFundResponse;
    message.fund = [];
    if (object.fund !== undefined && object.fund !== null) {
      for (const e of object.fund) {
        message.fund.push(Fund.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryFundPriceRequest: object = { symbol: "" };

export const QueryFundPriceRequest = {
  encode(
    message: QueryFundPriceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryFundPriceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryFundPriceRequest } as QueryFundPriceRequest;
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
    const message = { ...baseQueryFundPriceRequest } as QueryFundPriceRequest;
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = String(object.symbol);
    } else {
      message.symbol = "";
    }
    return message;
  },

  toJSON(message: QueryFundPriceRequest): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryFundPriceRequest>
  ): QueryFundPriceRequest {
    const message = { ...baseQueryFundPriceRequest } as QueryFundPriceRequest;
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = object.symbol;
    } else {
      message.symbol = "";
    }
    return message;
  },
};

const baseQueryFundPriceResponse: object = {};

export const QueryFundPriceResponse = {
  encode(
    message: QueryFundPriceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.price !== undefined) {
      FundPrice.encode(message.price, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryFundPriceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryFundPriceResponse } as QueryFundPriceResponse;
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
    const message = { ...baseQueryFundPriceResponse } as QueryFundPriceResponse;
    if (object.price !== undefined && object.price !== null) {
      message.price = FundPrice.fromJSON(object.price);
    } else {
      message.price = undefined;
    }
    return message;
  },

  toJSON(message: QueryFundPriceResponse): unknown {
    const obj: any = {};
    message.price !== undefined &&
      (obj.price = message.price ? FundPrice.toJSON(message.price) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryFundPriceResponse>
  ): QueryFundPriceResponse {
    const message = { ...baseQueryFundPriceResponse } as QueryFundPriceResponse;
    if (object.price !== undefined && object.price !== null) {
      message.price = FundPrice.fromPartial(object.price);
    } else {
      message.price = undefined;
    }
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
  }
  Fund(request: QueryGetFundRequest): Promise<QueryGetFundResponse> {
    const data = QueryGetFundRequest.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.etf.Query",
      "Fund",
      data
    );
    return promise.then((data) =>
      QueryGetFundResponse.decode(new Reader(data))
    );
  }

  FundAll(request: QueryAllFundRequest): Promise<QueryAllFundResponse> {
    const data = QueryAllFundRequest.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.etf.Query",
      "FundAll",
      data
    );
    return promise.then((data) =>
      QueryAllFundResponse.decode(new Reader(data))
    );
  }

  FundPrice(request: QueryFundPriceRequest): Promise<QueryFundPriceResponse> {
    const data = QueryFundPriceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.etf.Query",
      "FundPrice",
      data
    );
    return promise.then((data) =>
      QueryFundPriceResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
