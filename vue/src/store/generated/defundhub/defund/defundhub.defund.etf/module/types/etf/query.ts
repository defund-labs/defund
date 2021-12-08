/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Fund } from "../etf/fund";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "defundhub.defund.etf";

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

const baseQueryGetFundRequest: object = { index: "" };

export const QueryGetFundRequest = {
  encode(
    message: QueryGetFundRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
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
          message.index = reader.string();
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
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetFundRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetFundRequest>): QueryGetFundRequest {
    const message = { ...baseQueryGetFundRequest } as QueryGetFundRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
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

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a fund by index. */
  Fund(request: QueryGetFundRequest): Promise<QueryGetFundResponse>;
  /** Queries a list of fund items. */
  FundAll(request: QueryAllFundRequest): Promise<QueryAllFundResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Fund(request: QueryGetFundRequest): Promise<QueryGetFundResponse> {
    const data = QueryGetFundRequest.encode(request).finish();
    const promise = this.rpc.request(
      "defundhub.defund.etf.Query",
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
      "defundhub.defund.etf.Query",
      "FundAll",
      data
    );
    return promise.then((data) =>
      QueryAllFundResponse.decode(new Reader(data))
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
