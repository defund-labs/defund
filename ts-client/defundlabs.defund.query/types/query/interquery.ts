/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Height } from "../ibc/core/client/v1/client";

export const protobufPackage = "defundlabs.defund.query";

export interface Interquery {
  storeid: string;
  chainid: string;
  path: string;
  key: Uint8Array;
  timeoutHeight: number;
  connectionId: string;
  clientId: string;
}

export interface InterqueryResult {
  creator: string;
  storeid: string;
  chainid: string;
  data: Uint8Array;
  /** queried chain height on submission */
  height:
    | Height
    | undefined;
  /** querying chain height on submission */
  localHeight: number;
  success: boolean;
  proved: boolean;
}

export interface InterqueryTimeoutResult {
  storeid: string;
  timeoutHeight: number;
}

function createBaseInterquery(): Interquery {
  return {
    storeid: "",
    chainid: "",
    path: "",
    key: new Uint8Array(),
    timeoutHeight: 0,
    connectionId: "",
    clientId: "",
  };
}

export const Interquery = {
  encode(message: Interquery, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.storeid !== "") {
      writer.uint32(10).string(message.storeid);
    }
    if (message.chainid !== "") {
      writer.uint32(18).string(message.chainid);
    }
    if (message.path !== "") {
      writer.uint32(26).string(message.path);
    }
    if (message.key.length !== 0) {
      writer.uint32(34).bytes(message.key);
    }
    if (message.timeoutHeight !== 0) {
      writer.uint32(40).uint64(message.timeoutHeight);
    }
    if (message.connectionId !== "") {
      writer.uint32(50).string(message.connectionId);
    }
    if (message.clientId !== "") {
      writer.uint32(58).string(message.clientId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Interquery {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInterquery();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storeid = reader.string();
          break;
        case 2:
          message.chainid = reader.string();
          break;
        case 3:
          message.path = reader.string();
          break;
        case 4:
          message.key = reader.bytes();
          break;
        case 5:
          message.timeoutHeight = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.connectionId = reader.string();
          break;
        case 7:
          message.clientId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Interquery {
    return {
      storeid: isSet(object.storeid) ? String(object.storeid) : "",
      chainid: isSet(object.chainid) ? String(object.chainid) : "",
      path: isSet(object.path) ? String(object.path) : "",
      key: isSet(object.key) ? bytesFromBase64(object.key) : new Uint8Array(),
      timeoutHeight: isSet(object.timeoutHeight) ? Number(object.timeoutHeight) : 0,
      connectionId: isSet(object.connectionId) ? String(object.connectionId) : "",
      clientId: isSet(object.clientId) ? String(object.clientId) : "",
    };
  },

  toJSON(message: Interquery): unknown {
    const obj: any = {};
    message.storeid !== undefined && (obj.storeid = message.storeid);
    message.chainid !== undefined && (obj.chainid = message.chainid);
    message.path !== undefined && (obj.path = message.path);
    message.key !== undefined
      && (obj.key = base64FromBytes(message.key !== undefined ? message.key : new Uint8Array()));
    message.timeoutHeight !== undefined && (obj.timeoutHeight = Math.round(message.timeoutHeight));
    message.connectionId !== undefined && (obj.connectionId = message.connectionId);
    message.clientId !== undefined && (obj.clientId = message.clientId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Interquery>, I>>(object: I): Interquery {
    const message = createBaseInterquery();
    message.storeid = object.storeid ?? "";
    message.chainid = object.chainid ?? "";
    message.path = object.path ?? "";
    message.key = object.key ?? new Uint8Array();
    message.timeoutHeight = object.timeoutHeight ?? 0;
    message.connectionId = object.connectionId ?? "";
    message.clientId = object.clientId ?? "";
    return message;
  },
};

function createBaseInterqueryResult(): InterqueryResult {
  return {
    creator: "",
    storeid: "",
    chainid: "",
    data: new Uint8Array(),
    height: undefined,
    localHeight: 0,
    success: false,
    proved: false,
  };
}

export const InterqueryResult = {
  encode(message: InterqueryResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.storeid !== "") {
      writer.uint32(18).string(message.storeid);
    }
    if (message.chainid !== "") {
      writer.uint32(26).string(message.chainid);
    }
    if (message.data.length !== 0) {
      writer.uint32(34).bytes(message.data);
    }
    if (message.height !== undefined) {
      Height.encode(message.height, writer.uint32(42).fork()).ldelim();
    }
    if (message.localHeight !== 0) {
      writer.uint32(48).uint64(message.localHeight);
    }
    if (message.success === true) {
      writer.uint32(56).bool(message.success);
    }
    if (message.proved === true) {
      writer.uint32(64).bool(message.proved);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InterqueryResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInterqueryResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.storeid = reader.string();
          break;
        case 3:
          message.chainid = reader.string();
          break;
        case 4:
          message.data = reader.bytes();
          break;
        case 5:
          message.height = Height.decode(reader, reader.uint32());
          break;
        case 6:
          message.localHeight = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.success = reader.bool();
          break;
        case 8:
          message.proved = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InterqueryResult {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      storeid: isSet(object.storeid) ? String(object.storeid) : "",
      chainid: isSet(object.chainid) ? String(object.chainid) : "",
      data: isSet(object.data) ? bytesFromBase64(object.data) : new Uint8Array(),
      height: isSet(object.height) ? Height.fromJSON(object.height) : undefined,
      localHeight: isSet(object.localHeight) ? Number(object.localHeight) : 0,
      success: isSet(object.success) ? Boolean(object.success) : false,
      proved: isSet(object.proved) ? Boolean(object.proved) : false,
    };
  },

  toJSON(message: InterqueryResult): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.storeid !== undefined && (obj.storeid = message.storeid);
    message.chainid !== undefined && (obj.chainid = message.chainid);
    message.data !== undefined
      && (obj.data = base64FromBytes(message.data !== undefined ? message.data : new Uint8Array()));
    message.height !== undefined && (obj.height = message.height ? Height.toJSON(message.height) : undefined);
    message.localHeight !== undefined && (obj.localHeight = Math.round(message.localHeight));
    message.success !== undefined && (obj.success = message.success);
    message.proved !== undefined && (obj.proved = message.proved);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InterqueryResult>, I>>(object: I): InterqueryResult {
    const message = createBaseInterqueryResult();
    message.creator = object.creator ?? "";
    message.storeid = object.storeid ?? "";
    message.chainid = object.chainid ?? "";
    message.data = object.data ?? new Uint8Array();
    message.height = (object.height !== undefined && object.height !== null)
      ? Height.fromPartial(object.height)
      : undefined;
    message.localHeight = object.localHeight ?? 0;
    message.success = object.success ?? false;
    message.proved = object.proved ?? false;
    return message;
  },
};

function createBaseInterqueryTimeoutResult(): InterqueryTimeoutResult {
  return { storeid: "", timeoutHeight: 0 };
}

export const InterqueryTimeoutResult = {
  encode(message: InterqueryTimeoutResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.storeid !== "") {
      writer.uint32(10).string(message.storeid);
    }
    if (message.timeoutHeight !== 0) {
      writer.uint32(16).uint64(message.timeoutHeight);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InterqueryTimeoutResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInterqueryTimeoutResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storeid = reader.string();
          break;
        case 2:
          message.timeoutHeight = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InterqueryTimeoutResult {
    return {
      storeid: isSet(object.storeid) ? String(object.storeid) : "",
      timeoutHeight: isSet(object.timeoutHeight) ? Number(object.timeoutHeight) : 0,
    };
  },

  toJSON(message: InterqueryTimeoutResult): unknown {
    const obj: any = {};
    message.storeid !== undefined && (obj.storeid = message.storeid);
    message.timeoutHeight !== undefined && (obj.timeoutHeight = Math.round(message.timeoutHeight));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InterqueryTimeoutResult>, I>>(object: I): InterqueryTimeoutResult {
    const message = createBaseInterqueryTimeoutResult();
    message.storeid = object.storeid ?? "";
    message.timeoutHeight = object.timeoutHeight ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
