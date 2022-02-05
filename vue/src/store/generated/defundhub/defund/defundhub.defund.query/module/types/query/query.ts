/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Interquery } from "../query/interquery";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "defundhub.defund.query";

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

const baseQueryGetInterqueryRequest: object = { index: "" };

export const QueryGetInterqueryRequest = {
  encode(
    message: QueryGetInterqueryRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetInterqueryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetInterqueryRequest,
    } as QueryGetInterqueryRequest;
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

  fromJSON(object: any): QueryGetInterqueryRequest {
    const message = {
      ...baseQueryGetInterqueryRequest,
    } as QueryGetInterqueryRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetInterqueryRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetInterqueryRequest>
  ): QueryGetInterqueryRequest {
    const message = {
      ...baseQueryGetInterqueryRequest,
    } as QueryGetInterqueryRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetInterqueryResponse: object = {};

export const QueryGetInterqueryResponse = {
  encode(
    message: QueryGetInterqueryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.interquery !== undefined) {
      Interquery.encode(message.interquery, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetInterqueryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetInterqueryResponse,
    } as QueryGetInterqueryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interquery = Interquery.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetInterqueryResponse {
    const message = {
      ...baseQueryGetInterqueryResponse,
    } as QueryGetInterqueryResponse;
    if (object.interquery !== undefined && object.interquery !== null) {
      message.interquery = Interquery.fromJSON(object.interquery);
    } else {
      message.interquery = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetInterqueryResponse): unknown {
    const obj: any = {};
    message.interquery !== undefined &&
      (obj.interquery = message.interquery
        ? Interquery.toJSON(message.interquery)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetInterqueryResponse>
  ): QueryGetInterqueryResponse {
    const message = {
      ...baseQueryGetInterqueryResponse,
    } as QueryGetInterqueryResponse;
    if (object.interquery !== undefined && object.interquery !== null) {
      message.interquery = Interquery.fromPartial(object.interquery);
    } else {
      message.interquery = undefined;
    }
    return message;
  },
};

const baseQueryAllInterqueryRequest: object = {};

export const QueryAllInterqueryRequest = {
  encode(
    message: QueryAllInterqueryRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllInterqueryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllInterqueryRequest,
    } as QueryAllInterqueryRequest;
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

  fromJSON(object: any): QueryAllInterqueryRequest {
    const message = {
      ...baseQueryAllInterqueryRequest,
    } as QueryAllInterqueryRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllInterqueryRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllInterqueryRequest>
  ): QueryAllInterqueryRequest {
    const message = {
      ...baseQueryAllInterqueryRequest,
    } as QueryAllInterqueryRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllInterqueryResponse: object = {};

export const QueryAllInterqueryResponse = {
  encode(
    message: QueryAllInterqueryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.interquery) {
      Interquery.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllInterqueryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllInterqueryResponse,
    } as QueryAllInterqueryResponse;
    message.interquery = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interquery.push(Interquery.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllInterqueryResponse {
    const message = {
      ...baseQueryAllInterqueryResponse,
    } as QueryAllInterqueryResponse;
    message.interquery = [];
    if (object.interquery !== undefined && object.interquery !== null) {
      for (const e of object.interquery) {
        message.interquery.push(Interquery.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllInterqueryResponse): unknown {
    const obj: any = {};
    if (message.interquery) {
      obj.interquery = message.interquery.map((e) =>
        e ? Interquery.toJSON(e) : undefined
      );
    } else {
      obj.interquery = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllInterqueryResponse>
  ): QueryAllInterqueryResponse {
    const message = {
      ...baseQueryAllInterqueryResponse,
    } as QueryAllInterqueryResponse;
    message.interquery = [];
    if (object.interquery !== undefined && object.interquery !== null) {
      for (const e of object.interquery) {
        message.interquery.push(Interquery.fromPartial(e));
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
  /** Queries a interquery by index. */
  Interquery(
    request: QueryGetInterqueryRequest
  ): Promise<QueryGetInterqueryResponse>;
  /** Queries a list of interquery items. */
  InterqueryAll(
    request: QueryAllInterqueryRequest
  ): Promise<QueryAllInterqueryResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Interquery(
    request: QueryGetInterqueryRequest
  ): Promise<QueryGetInterqueryResponse> {
    const data = QueryGetInterqueryRequest.encode(request).finish();
    const promise = this.rpc.request(
      "defundhub.defund.query.Query",
      "Interquery",
      data
    );
    return promise.then((data) =>
      QueryGetInterqueryResponse.decode(new Reader(data))
    );
  }

  InterqueryAll(
    request: QueryAllInterqueryRequest
  ): Promise<QueryAllInterqueryResponse> {
    const data = QueryAllInterqueryRequest.encode(request).finish();
    const promise = this.rpc.request(
      "defundhub.defund.query.Query",
      "InterqueryAll",
      data
    );
    return promise.then((data) =>
      QueryAllInterqueryResponse.decode(new Reader(data))
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
