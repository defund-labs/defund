/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import {
  Interquery,
  InterqueryResult,
  InterqueryTimeoutResult,
} from "../query/interquery";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "defundhub.defund.query";

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

const baseQueryGetInterqueryRequest: object = { storeid: "" };

export const QueryGetInterqueryRequest = {
  encode(
    message: QueryGetInterqueryRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.storeid !== "") {
      writer.uint32(10).string(message.storeid);
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
    const message = {
      ...baseQueryGetInterqueryRequest,
    } as QueryGetInterqueryRequest;
    if (object.storeid !== undefined && object.storeid !== null) {
      message.storeid = String(object.storeid);
    } else {
      message.storeid = "";
    }
    return message;
  },

  toJSON(message: QueryGetInterqueryRequest): unknown {
    const obj: any = {};
    message.storeid !== undefined && (obj.storeid = message.storeid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetInterqueryRequest>
  ): QueryGetInterqueryRequest {
    const message = {
      ...baseQueryGetInterqueryRequest,
    } as QueryGetInterqueryRequest;
    if (object.storeid !== undefined && object.storeid !== null) {
      message.storeid = object.storeid;
    } else {
      message.storeid = "";
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

const baseQueryGetInterqueryResultRequest: object = { storeid: "" };

export const QueryGetInterqueryResultRequest = {
  encode(
    message: QueryGetInterqueryResultRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.storeid !== "") {
      writer.uint32(10).string(message.storeid);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetInterqueryResultRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetInterqueryResultRequest,
    } as QueryGetInterqueryResultRequest;
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
    const message = {
      ...baseQueryGetInterqueryResultRequest,
    } as QueryGetInterqueryResultRequest;
    if (object.storeid !== undefined && object.storeid !== null) {
      message.storeid = String(object.storeid);
    } else {
      message.storeid = "";
    }
    return message;
  },

  toJSON(message: QueryGetInterqueryResultRequest): unknown {
    const obj: any = {};
    message.storeid !== undefined && (obj.storeid = message.storeid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetInterqueryResultRequest>
  ): QueryGetInterqueryResultRequest {
    const message = {
      ...baseQueryGetInterqueryResultRequest,
    } as QueryGetInterqueryResultRequest;
    if (object.storeid !== undefined && object.storeid !== null) {
      message.storeid = object.storeid;
    } else {
      message.storeid = "";
    }
    return message;
  },
};

const baseQueryGetInterqueryResultResponse: object = {};

export const QueryGetInterqueryResultResponse = {
  encode(
    message: QueryGetInterqueryResultResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.interqueryresult !== undefined) {
      InterqueryResult.encode(
        message.interqueryresult,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetInterqueryResultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetInterqueryResultResponse,
    } as QueryGetInterqueryResultResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interqueryresult = InterqueryResult.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetInterqueryResultResponse {
    const message = {
      ...baseQueryGetInterqueryResultResponse,
    } as QueryGetInterqueryResultResponse;
    if (
      object.interqueryresult !== undefined &&
      object.interqueryresult !== null
    ) {
      message.interqueryresult = InterqueryResult.fromJSON(
        object.interqueryresult
      );
    } else {
      message.interqueryresult = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetInterqueryResultResponse): unknown {
    const obj: any = {};
    message.interqueryresult !== undefined &&
      (obj.interqueryresult = message.interqueryresult
        ? InterqueryResult.toJSON(message.interqueryresult)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetInterqueryResultResponse>
  ): QueryGetInterqueryResultResponse {
    const message = {
      ...baseQueryGetInterqueryResultResponse,
    } as QueryGetInterqueryResultResponse;
    if (
      object.interqueryresult !== undefined &&
      object.interqueryresult !== null
    ) {
      message.interqueryresult = InterqueryResult.fromPartial(
        object.interqueryresult
      );
    } else {
      message.interqueryresult = undefined;
    }
    return message;
  },
};

const baseQueryAllInterqueryResultRequest: object = {};

export const QueryAllInterqueryResultRequest = {
  encode(
    message: QueryAllInterqueryResultRequest,
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
  ): QueryAllInterqueryResultRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllInterqueryResultRequest,
    } as QueryAllInterqueryResultRequest;
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
    const message = {
      ...baseQueryAllInterqueryResultRequest,
    } as QueryAllInterqueryResultRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllInterqueryResultRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllInterqueryResultRequest>
  ): QueryAllInterqueryResultRequest {
    const message = {
      ...baseQueryAllInterqueryResultRequest,
    } as QueryAllInterqueryResultRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllInterqueryResultResponse: object = {};

export const QueryAllInterqueryResultResponse = {
  encode(
    message: QueryAllInterqueryResultResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.interqueryresult) {
      InterqueryResult.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllInterqueryResultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllInterqueryResultResponse,
    } as QueryAllInterqueryResultResponse;
    message.interqueryresult = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interqueryresult.push(
            InterqueryResult.decode(reader, reader.uint32())
          );
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
    const message = {
      ...baseQueryAllInterqueryResultResponse,
    } as QueryAllInterqueryResultResponse;
    message.interqueryresult = [];
    if (
      object.interqueryresult !== undefined &&
      object.interqueryresult !== null
    ) {
      for (const e of object.interqueryresult) {
        message.interqueryresult.push(InterqueryResult.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllInterqueryResultResponse): unknown {
    const obj: any = {};
    if (message.interqueryresult) {
      obj.interqueryresult = message.interqueryresult.map((e) =>
        e ? InterqueryResult.toJSON(e) : undefined
      );
    } else {
      obj.interqueryresult = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllInterqueryResultResponse>
  ): QueryAllInterqueryResultResponse {
    const message = {
      ...baseQueryAllInterqueryResultResponse,
    } as QueryAllInterqueryResultResponse;
    message.interqueryresult = [];
    if (
      object.interqueryresult !== undefined &&
      object.interqueryresult !== null
    ) {
      for (const e of object.interqueryresult) {
        message.interqueryresult.push(InterqueryResult.fromPartial(e));
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

const baseQueryGetInterqueryTimeoutResultRequest: object = { storeid: "" };

export const QueryGetInterqueryTimeoutResultRequest = {
  encode(
    message: QueryGetInterqueryTimeoutResultRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.storeid !== "") {
      writer.uint32(10).string(message.storeid);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetInterqueryTimeoutResultRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetInterqueryTimeoutResultRequest,
    } as QueryGetInterqueryTimeoutResultRequest;
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
    const message = {
      ...baseQueryGetInterqueryTimeoutResultRequest,
    } as QueryGetInterqueryTimeoutResultRequest;
    if (object.storeid !== undefined && object.storeid !== null) {
      message.storeid = String(object.storeid);
    } else {
      message.storeid = "";
    }
    return message;
  },

  toJSON(message: QueryGetInterqueryTimeoutResultRequest): unknown {
    const obj: any = {};
    message.storeid !== undefined && (obj.storeid = message.storeid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetInterqueryTimeoutResultRequest>
  ): QueryGetInterqueryTimeoutResultRequest {
    const message = {
      ...baseQueryGetInterqueryTimeoutResultRequest,
    } as QueryGetInterqueryTimeoutResultRequest;
    if (object.storeid !== undefined && object.storeid !== null) {
      message.storeid = object.storeid;
    } else {
      message.storeid = "";
    }
    return message;
  },
};

const baseQueryGetInterqueryTimeoutResultResponse: object = {};

export const QueryGetInterqueryTimeoutResultResponse = {
  encode(
    message: QueryGetInterqueryTimeoutResultResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.interquerytimeoutresult !== undefined) {
      InterqueryTimeoutResult.encode(
        message.interquerytimeoutresult,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetInterqueryTimeoutResultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetInterqueryTimeoutResultResponse,
    } as QueryGetInterqueryTimeoutResultResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interquerytimeoutresult = InterqueryTimeoutResult.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetInterqueryTimeoutResultResponse {
    const message = {
      ...baseQueryGetInterqueryTimeoutResultResponse,
    } as QueryGetInterqueryTimeoutResultResponse;
    if (
      object.interquerytimeoutresult !== undefined &&
      object.interquerytimeoutresult !== null
    ) {
      message.interquerytimeoutresult = InterqueryTimeoutResult.fromJSON(
        object.interquerytimeoutresult
      );
    } else {
      message.interquerytimeoutresult = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetInterqueryTimeoutResultResponse): unknown {
    const obj: any = {};
    message.interquerytimeoutresult !== undefined &&
      (obj.interquerytimeoutresult = message.interquerytimeoutresult
        ? InterqueryTimeoutResult.toJSON(message.interquerytimeoutresult)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetInterqueryTimeoutResultResponse>
  ): QueryGetInterqueryTimeoutResultResponse {
    const message = {
      ...baseQueryGetInterqueryTimeoutResultResponse,
    } as QueryGetInterqueryTimeoutResultResponse;
    if (
      object.interquerytimeoutresult !== undefined &&
      object.interquerytimeoutresult !== null
    ) {
      message.interquerytimeoutresult = InterqueryTimeoutResult.fromPartial(
        object.interquerytimeoutresult
      );
    } else {
      message.interquerytimeoutresult = undefined;
    }
    return message;
  },
};

const baseQueryAllInterqueryTimeoutResultRequest: object = {};

export const QueryAllInterqueryTimeoutResultRequest = {
  encode(
    message: QueryAllInterqueryTimeoutResultRequest,
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
  ): QueryAllInterqueryTimeoutResultRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllInterqueryTimeoutResultRequest,
    } as QueryAllInterqueryTimeoutResultRequest;
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
    const message = {
      ...baseQueryAllInterqueryTimeoutResultRequest,
    } as QueryAllInterqueryTimeoutResultRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllInterqueryTimeoutResultRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllInterqueryTimeoutResultRequest>
  ): QueryAllInterqueryTimeoutResultRequest {
    const message = {
      ...baseQueryAllInterqueryTimeoutResultRequest,
    } as QueryAllInterqueryTimeoutResultRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllInterqueryTimeoutResultResponse: object = {};

export const QueryAllInterqueryTimeoutResultResponse = {
  encode(
    message: QueryAllInterqueryTimeoutResultResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.interquerytimeoutresult) {
      InterqueryTimeoutResult.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllInterqueryTimeoutResultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllInterqueryTimeoutResultResponse,
    } as QueryAllInterqueryTimeoutResultResponse;
    message.interquerytimeoutresult = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interquerytimeoutresult.push(
            InterqueryTimeoutResult.decode(reader, reader.uint32())
          );
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
    const message = {
      ...baseQueryAllInterqueryTimeoutResultResponse,
    } as QueryAllInterqueryTimeoutResultResponse;
    message.interquerytimeoutresult = [];
    if (
      object.interquerytimeoutresult !== undefined &&
      object.interquerytimeoutresult !== null
    ) {
      for (const e of object.interquerytimeoutresult) {
        message.interquerytimeoutresult.push(
          InterqueryTimeoutResult.fromJSON(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
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
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllInterqueryTimeoutResultResponse>
  ): QueryAllInterqueryTimeoutResultResponse {
    const message = {
      ...baseQueryAllInterqueryTimeoutResultResponse,
    } as QueryAllInterqueryTimeoutResultResponse;
    message.interquerytimeoutresult = [];
    if (
      object.interquerytimeoutresult !== undefined &&
      object.interquerytimeoutresult !== null
    ) {
      for (const e of object.interquerytimeoutresult) {
        message.interquerytimeoutresult.push(
          InterqueryTimeoutResult.fromPartial(e)
        );
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
  /** Queries a interquery result by index. */
  InterqueryResult(
    request: QueryGetInterqueryResultRequest
  ): Promise<QueryGetInterqueryResultResponse>;
  /** Queries a list of interquery result items. */
  InterqueryResultAll(
    request: QueryAllInterqueryResultRequest
  ): Promise<QueryAllInterqueryResultResponse>;
  /** Queries a interquery timeout result by index. */
  InterqueryTimeoutResult(
    request: QueryGetInterqueryTimeoutResultRequest
  ): Promise<QueryGetInterqueryTimeoutResultResponse>;
  /** Queries a list of interquery timeout result items. */
  InterqueryTimeoutResultAll(
    request: QueryAllInterqueryTimeoutResultRequest
  ): Promise<QueryAllInterqueryTimeoutResultResponse>;
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

  InterqueryResult(
    request: QueryGetInterqueryResultRequest
  ): Promise<QueryGetInterqueryResultResponse> {
    const data = QueryGetInterqueryResultRequest.encode(request).finish();
    const promise = this.rpc.request(
      "defundhub.defund.query.Query",
      "InterqueryResult",
      data
    );
    return promise.then((data) =>
      QueryGetInterqueryResultResponse.decode(new Reader(data))
    );
  }

  InterqueryResultAll(
    request: QueryAllInterqueryResultRequest
  ): Promise<QueryAllInterqueryResultResponse> {
    const data = QueryAllInterqueryResultRequest.encode(request).finish();
    const promise = this.rpc.request(
      "defundhub.defund.query.Query",
      "InterqueryResultAll",
      data
    );
    return promise.then((data) =>
      QueryAllInterqueryResultResponse.decode(new Reader(data))
    );
  }

  InterqueryTimeoutResult(
    request: QueryGetInterqueryTimeoutResultRequest
  ): Promise<QueryGetInterqueryTimeoutResultResponse> {
    const data = QueryGetInterqueryTimeoutResultRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "defundhub.defund.query.Query",
      "InterqueryTimeoutResult",
      data
    );
    return promise.then((data) =>
      QueryGetInterqueryTimeoutResultResponse.decode(new Reader(data))
    );
  }

  InterqueryTimeoutResultAll(
    request: QueryAllInterqueryTimeoutResultRequest
  ): Promise<QueryAllInterqueryTimeoutResultResponse> {
    const data = QueryAllInterqueryTimeoutResultRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "defundhub.defund.query.Query",
      "InterqueryTimeoutResultAll",
      data
    );
    return promise.then((data) =>
      QueryAllInterqueryTimeoutResultResponse.decode(new Reader(data))
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
