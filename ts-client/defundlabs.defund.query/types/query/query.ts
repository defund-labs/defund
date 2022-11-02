/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Interquery, InterqueryResult, InterqueryTimeoutResult } from "./interquery";

export const protobufPackage = "defundlabs.defund.query";

export interface QueryGetInterqueryRequest {
  storeid: string;
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

export interface QueryGetInterqueryResultRequest {
  storeid: string;
}

export interface QueryGetInterqueryResultResponse {
  interqueryresult: InterqueryResult | undefined;
}

export interface QueryAllInterqueryResultRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllInterqueryResultResponse {
  interqueryresult: InterqueryResult[];
  pagination: PageResponse | undefined;
}

export interface QueryGetInterqueryTimeoutResultRequest {
  storeid: string;
}

export interface QueryGetInterqueryTimeoutResultResponse {
  interquerytimeoutresult: InterqueryTimeoutResult | undefined;
}

export interface QueryAllInterqueryTimeoutResultRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllInterqueryTimeoutResultResponse {
  interquerytimeoutresult: InterqueryTimeoutResult[];
  pagination: PageResponse | undefined;
}

function createBaseQueryGetInterqueryRequest(): QueryGetInterqueryRequest {
  return { storeid: "" };
}

export const QueryGetInterqueryRequest = {
  encode(message: QueryGetInterqueryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.storeid !== "") {
      writer.uint32(10).string(message.storeid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetInterqueryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetInterqueryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storeid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetInterqueryRequest {
    return { storeid: isSet(object.storeid) ? String(object.storeid) : "" };
  },

  toJSON(message: QueryGetInterqueryRequest): unknown {
    const obj: any = {};
    message.storeid !== undefined && (obj.storeid = message.storeid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetInterqueryRequest>, I>>(object: I): QueryGetInterqueryRequest {
    const message = createBaseQueryGetInterqueryRequest();
    message.storeid = object.storeid ?? "";
    return message;
  },
};

function createBaseQueryGetInterqueryResponse(): QueryGetInterqueryResponse {
  return { interquery: undefined };
}

export const QueryGetInterqueryResponse = {
  encode(message: QueryGetInterqueryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.interquery !== undefined) {
      Interquery.encode(message.interquery, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetInterqueryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetInterqueryResponse();
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
    return { interquery: isSet(object.interquery) ? Interquery.fromJSON(object.interquery) : undefined };
  },

  toJSON(message: QueryGetInterqueryResponse): unknown {
    const obj: any = {};
    message.interquery !== undefined
      && (obj.interquery = message.interquery ? Interquery.toJSON(message.interquery) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetInterqueryResponse>, I>>(object: I): QueryGetInterqueryResponse {
    const message = createBaseQueryGetInterqueryResponse();
    message.interquery = (object.interquery !== undefined && object.interquery !== null)
      ? Interquery.fromPartial(object.interquery)
      : undefined;
    return message;
  },
};

function createBaseQueryAllInterqueryRequest(): QueryAllInterqueryRequest {
  return { pagination: undefined };
}

export const QueryAllInterqueryRequest = {
  encode(message: QueryAllInterqueryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllInterqueryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllInterqueryRequest();
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
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllInterqueryRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllInterqueryRequest>, I>>(object: I): QueryAllInterqueryRequest {
    const message = createBaseQueryAllInterqueryRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllInterqueryResponse(): QueryAllInterqueryResponse {
  return { interquery: [], pagination: undefined };
}

export const QueryAllInterqueryResponse = {
  encode(message: QueryAllInterqueryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.interquery) {
      Interquery.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllInterqueryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllInterqueryResponse();
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
    return {
      interquery: Array.isArray(object?.interquery) ? object.interquery.map((e: any) => Interquery.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllInterqueryResponse): unknown {
    const obj: any = {};
    if (message.interquery) {
      obj.interquery = message.interquery.map((e) => e ? Interquery.toJSON(e) : undefined);
    } else {
      obj.interquery = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllInterqueryResponse>, I>>(object: I): QueryAllInterqueryResponse {
    const message = createBaseQueryAllInterqueryResponse();
    message.interquery = object.interquery?.map((e) => Interquery.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetInterqueryResultRequest(): QueryGetInterqueryResultRequest {
  return { storeid: "" };
}

export const QueryGetInterqueryResultRequest = {
  encode(message: QueryGetInterqueryResultRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.storeid !== "") {
      writer.uint32(10).string(message.storeid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetInterqueryResultRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetInterqueryResultRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storeid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetInterqueryResultRequest {
    return { storeid: isSet(object.storeid) ? String(object.storeid) : "" };
  },

  toJSON(message: QueryGetInterqueryResultRequest): unknown {
    const obj: any = {};
    message.storeid !== undefined && (obj.storeid = message.storeid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetInterqueryResultRequest>, I>>(
    object: I,
  ): QueryGetInterqueryResultRequest {
    const message = createBaseQueryGetInterqueryResultRequest();
    message.storeid = object.storeid ?? "";
    return message;
  },
};

function createBaseQueryGetInterqueryResultResponse(): QueryGetInterqueryResultResponse {
  return { interqueryresult: undefined };
}

export const QueryGetInterqueryResultResponse = {
  encode(message: QueryGetInterqueryResultResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.interqueryresult !== undefined) {
      InterqueryResult.encode(message.interqueryresult, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetInterqueryResultResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetInterqueryResultResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interqueryresult = InterqueryResult.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetInterqueryResultResponse {
    return {
      interqueryresult: isSet(object.interqueryresult) ? InterqueryResult.fromJSON(object.interqueryresult) : undefined,
    };
  },

  toJSON(message: QueryGetInterqueryResultResponse): unknown {
    const obj: any = {};
    message.interqueryresult !== undefined && (obj.interqueryresult = message.interqueryresult
      ? InterqueryResult.toJSON(message.interqueryresult)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetInterqueryResultResponse>, I>>(
    object: I,
  ): QueryGetInterqueryResultResponse {
    const message = createBaseQueryGetInterqueryResultResponse();
    message.interqueryresult = (object.interqueryresult !== undefined && object.interqueryresult !== null)
      ? InterqueryResult.fromPartial(object.interqueryresult)
      : undefined;
    return message;
  },
};

function createBaseQueryAllInterqueryResultRequest(): QueryAllInterqueryResultRequest {
  return { pagination: undefined };
}

export const QueryAllInterqueryResultRequest = {
  encode(message: QueryAllInterqueryResultRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllInterqueryResultRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllInterqueryResultRequest();
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

  fromJSON(object: any): QueryAllInterqueryResultRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllInterqueryResultRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllInterqueryResultRequest>, I>>(
    object: I,
  ): QueryAllInterqueryResultRequest {
    const message = createBaseQueryAllInterqueryResultRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllInterqueryResultResponse(): QueryAllInterqueryResultResponse {
  return { interqueryresult: [], pagination: undefined };
}

export const QueryAllInterqueryResultResponse = {
  encode(message: QueryAllInterqueryResultResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.interqueryresult) {
      InterqueryResult.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllInterqueryResultResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllInterqueryResultResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interqueryresult.push(InterqueryResult.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllInterqueryResultResponse {
    return {
      interqueryresult: Array.isArray(object?.interqueryresult)
        ? object.interqueryresult.map((e: any) => InterqueryResult.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllInterqueryResultResponse): unknown {
    const obj: any = {};
    if (message.interqueryresult) {
      obj.interqueryresult = message.interqueryresult.map((e) => e ? InterqueryResult.toJSON(e) : undefined);
    } else {
      obj.interqueryresult = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllInterqueryResultResponse>, I>>(
    object: I,
  ): QueryAllInterqueryResultResponse {
    const message = createBaseQueryAllInterqueryResultResponse();
    message.interqueryresult = object.interqueryresult?.map((e) => InterqueryResult.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetInterqueryTimeoutResultRequest(): QueryGetInterqueryTimeoutResultRequest {
  return { storeid: "" };
}

export const QueryGetInterqueryTimeoutResultRequest = {
  encode(message: QueryGetInterqueryTimeoutResultRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.storeid !== "") {
      writer.uint32(10).string(message.storeid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetInterqueryTimeoutResultRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetInterqueryTimeoutResultRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storeid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetInterqueryTimeoutResultRequest {
    return { storeid: isSet(object.storeid) ? String(object.storeid) : "" };
  },

  toJSON(message: QueryGetInterqueryTimeoutResultRequest): unknown {
    const obj: any = {};
    message.storeid !== undefined && (obj.storeid = message.storeid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetInterqueryTimeoutResultRequest>, I>>(
    object: I,
  ): QueryGetInterqueryTimeoutResultRequest {
    const message = createBaseQueryGetInterqueryTimeoutResultRequest();
    message.storeid = object.storeid ?? "";
    return message;
  },
};

function createBaseQueryGetInterqueryTimeoutResultResponse(): QueryGetInterqueryTimeoutResultResponse {
  return { interquerytimeoutresult: undefined };
}

export const QueryGetInterqueryTimeoutResultResponse = {
  encode(message: QueryGetInterqueryTimeoutResultResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.interquerytimeoutresult !== undefined) {
      InterqueryTimeoutResult.encode(message.interquerytimeoutresult, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetInterqueryTimeoutResultResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetInterqueryTimeoutResultResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interquerytimeoutresult = InterqueryTimeoutResult.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetInterqueryTimeoutResultResponse {
    return {
      interquerytimeoutresult: isSet(object.interquerytimeoutresult)
        ? InterqueryTimeoutResult.fromJSON(object.interquerytimeoutresult)
        : undefined,
    };
  },

  toJSON(message: QueryGetInterqueryTimeoutResultResponse): unknown {
    const obj: any = {};
    message.interquerytimeoutresult !== undefined && (obj.interquerytimeoutresult = message.interquerytimeoutresult
      ? InterqueryTimeoutResult.toJSON(message.interquerytimeoutresult)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetInterqueryTimeoutResultResponse>, I>>(
    object: I,
  ): QueryGetInterqueryTimeoutResultResponse {
    const message = createBaseQueryGetInterqueryTimeoutResultResponse();
    message.interquerytimeoutresult =
      (object.interquerytimeoutresult !== undefined && object.interquerytimeoutresult !== null)
        ? InterqueryTimeoutResult.fromPartial(object.interquerytimeoutresult)
        : undefined;
    return message;
  },
};

function createBaseQueryAllInterqueryTimeoutResultRequest(): QueryAllInterqueryTimeoutResultRequest {
  return { pagination: undefined };
}

export const QueryAllInterqueryTimeoutResultRequest = {
  encode(message: QueryAllInterqueryTimeoutResultRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllInterqueryTimeoutResultRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllInterqueryTimeoutResultRequest();
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

  fromJSON(object: any): QueryAllInterqueryTimeoutResultRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllInterqueryTimeoutResultRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllInterqueryTimeoutResultRequest>, I>>(
    object: I,
  ): QueryAllInterqueryTimeoutResultRequest {
    const message = createBaseQueryAllInterqueryTimeoutResultRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllInterqueryTimeoutResultResponse(): QueryAllInterqueryTimeoutResultResponse {
  return { interquerytimeoutresult: [], pagination: undefined };
}

export const QueryAllInterqueryTimeoutResultResponse = {
  encode(message: QueryAllInterqueryTimeoutResultResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.interquerytimeoutresult) {
      InterqueryTimeoutResult.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllInterqueryTimeoutResultResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllInterqueryTimeoutResultResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interquerytimeoutresult.push(InterqueryTimeoutResult.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllInterqueryTimeoutResultResponse {
    return {
      interquerytimeoutresult: Array.isArray(object?.interquerytimeoutresult)
        ? object.interquerytimeoutresult.map((e: any) => InterqueryTimeoutResult.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllInterqueryTimeoutResultResponse): unknown {
    const obj: any = {};
    if (message.interquerytimeoutresult) {
      obj.interquerytimeoutresult = message.interquerytimeoutresult.map((e) =>
        e ? InterqueryTimeoutResult.toJSON(e) : undefined
      );
    } else {
      obj.interquerytimeoutresult = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllInterqueryTimeoutResultResponse>, I>>(
    object: I,
  ): QueryAllInterqueryTimeoutResultResponse {
    const message = createBaseQueryAllInterqueryTimeoutResultResponse();
    message.interquerytimeoutresult = object.interquerytimeoutresult?.map((e) => InterqueryTimeoutResult.fromPartial(e))
      || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a interquery by index. */
  Interquery(request: QueryGetInterqueryRequest): Promise<QueryGetInterqueryResponse>;
  /** Queries a list of interquery items. */
  InterqueryAll(request: QueryAllInterqueryRequest): Promise<QueryAllInterqueryResponse>;
  /** Queries a interquery result by index. */
  InterqueryResult(request: QueryGetInterqueryResultRequest): Promise<QueryGetInterqueryResultResponse>;
  /** Queries a list of interquery result items. */
  InterqueryResultAll(request: QueryAllInterqueryResultRequest): Promise<QueryAllInterqueryResultResponse>;
  /** Queries a interquery timeout result by index. */
  InterqueryTimeoutResult(
    request: QueryGetInterqueryTimeoutResultRequest,
  ): Promise<QueryGetInterqueryTimeoutResultResponse>;
  /** Queries a list of interquery timeout result items. */
  InterqueryTimeoutResultAll(
    request: QueryAllInterqueryTimeoutResultRequest,
  ): Promise<QueryAllInterqueryTimeoutResultResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Interquery = this.Interquery.bind(this);
    this.InterqueryAll = this.InterqueryAll.bind(this);
    this.InterqueryResult = this.InterqueryResult.bind(this);
    this.InterqueryResultAll = this.InterqueryResultAll.bind(this);
    this.InterqueryTimeoutResult = this.InterqueryTimeoutResult.bind(this);
    this.InterqueryTimeoutResultAll = this.InterqueryTimeoutResultAll.bind(this);
  }
  Interquery(request: QueryGetInterqueryRequest): Promise<QueryGetInterqueryResponse> {
    const data = QueryGetInterqueryRequest.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.query.Query", "Interquery", data);
    return promise.then((data) => QueryGetInterqueryResponse.decode(new _m0.Reader(data)));
  }

  InterqueryAll(request: QueryAllInterqueryRequest): Promise<QueryAllInterqueryResponse> {
    const data = QueryAllInterqueryRequest.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.query.Query", "InterqueryAll", data);
    return promise.then((data) => QueryAllInterqueryResponse.decode(new _m0.Reader(data)));
  }

  InterqueryResult(request: QueryGetInterqueryResultRequest): Promise<QueryGetInterqueryResultResponse> {
    const data = QueryGetInterqueryResultRequest.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.query.Query", "InterqueryResult", data);
    return promise.then((data) => QueryGetInterqueryResultResponse.decode(new _m0.Reader(data)));
  }

  InterqueryResultAll(request: QueryAllInterqueryResultRequest): Promise<QueryAllInterqueryResultResponse> {
    const data = QueryAllInterqueryResultRequest.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.query.Query", "InterqueryResultAll", data);
    return promise.then((data) => QueryAllInterqueryResultResponse.decode(new _m0.Reader(data)));
  }

  InterqueryTimeoutResult(
    request: QueryGetInterqueryTimeoutResultRequest,
  ): Promise<QueryGetInterqueryTimeoutResultResponse> {
    const data = QueryGetInterqueryTimeoutResultRequest.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.query.Query", "InterqueryTimeoutResult", data);
    return promise.then((data) => QueryGetInterqueryTimeoutResultResponse.decode(new _m0.Reader(data)));
  }

  InterqueryTimeoutResultAll(
    request: QueryAllInterqueryTimeoutResultRequest,
  ): Promise<QueryAllInterqueryTimeoutResultResponse> {
    const data = QueryAllInterqueryTimeoutResultRequest.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.query.Query", "InterqueryTimeoutResultAll", data);
    return promise.then((data) => QueryAllInterqueryTimeoutResultResponse.decode(new _m0.Reader(data)));
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
