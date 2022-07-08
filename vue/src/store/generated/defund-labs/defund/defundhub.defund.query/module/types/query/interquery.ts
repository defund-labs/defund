/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "defundhub.defund.query";

export interface Interquery {
  creator: string;
  storeid: string;
  chainid: string;
  path: string;
  key: Uint8Array;
  timeoutHeight: number;
  connectionId: string;
}

export interface InterqueryResult {
  creator: string;
  storeid: string;
  chainid: string;
  data: Uint8Array;
  height: number;
  success: boolean;
  proved: boolean;
}

export interface InterqueryTimeoutResult {
  creator: string;
  storeid: string;
  timeoutHeight: number;
}

const baseInterquery: object = {
  creator: "",
  storeid: "",
  chainid: "",
  path: "",
  timeoutHeight: 0,
  connectionId: "",
};

export const Interquery = {
  encode(message: Interquery, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.storeid !== "") {
      writer.uint32(18).string(message.storeid);
    }
    if (message.chainid !== "") {
      writer.uint32(26).string(message.chainid);
    }
    if (message.path !== "") {
      writer.uint32(34).string(message.path);
    }
    if (message.key.length !== 0) {
      writer.uint32(42).bytes(message.key);
    }
    if (message.timeoutHeight !== 0) {
      writer.uint32(48).uint64(message.timeoutHeight);
    }
    if (message.connectionId !== "") {
      writer.uint32(58).string(message.connectionId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Interquery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseInterquery } as Interquery;
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
          message.path = reader.string();
          break;
        case 5:
          message.key = reader.bytes();
          break;
        case 6:
          message.timeoutHeight = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.connectionId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Interquery {
    const message = { ...baseInterquery } as Interquery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.storeid !== undefined && object.storeid !== null) {
      message.storeid = String(object.storeid);
    } else {
      message.storeid = "";
    }
    if (object.chainid !== undefined && object.chainid !== null) {
      message.chainid = String(object.chainid);
    } else {
      message.chainid = "";
    }
    if (object.path !== undefined && object.path !== null) {
      message.path = String(object.path);
    } else {
      message.path = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = bytesFromBase64(object.key);
    }
    if (object.timeoutHeight !== undefined && object.timeoutHeight !== null) {
      message.timeoutHeight = Number(object.timeoutHeight);
    } else {
      message.timeoutHeight = 0;
    }
    if (object.connectionId !== undefined && object.connectionId !== null) {
      message.connectionId = String(object.connectionId);
    } else {
      message.connectionId = "";
    }
    return message;
  },

  toJSON(message: Interquery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.storeid !== undefined && (obj.storeid = message.storeid);
    message.chainid !== undefined && (obj.chainid = message.chainid);
    message.path !== undefined && (obj.path = message.path);
    message.key !== undefined &&
      (obj.key = base64FromBytes(
        message.key !== undefined ? message.key : new Uint8Array()
      ));
    message.timeoutHeight !== undefined &&
      (obj.timeoutHeight = message.timeoutHeight);
    message.connectionId !== undefined &&
      (obj.connectionId = message.connectionId);
    return obj;
  },

  fromPartial(object: DeepPartial<Interquery>): Interquery {
    const message = { ...baseInterquery } as Interquery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.storeid !== undefined && object.storeid !== null) {
      message.storeid = object.storeid;
    } else {
      message.storeid = "";
    }
    if (object.chainid !== undefined && object.chainid !== null) {
      message.chainid = object.chainid;
    } else {
      message.chainid = "";
    }
    if (object.path !== undefined && object.path !== null) {
      message.path = object.path;
    } else {
      message.path = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = new Uint8Array();
    }
    if (object.timeoutHeight !== undefined && object.timeoutHeight !== null) {
      message.timeoutHeight = object.timeoutHeight;
    } else {
      message.timeoutHeight = 0;
    }
    if (object.connectionId !== undefined && object.connectionId !== null) {
      message.connectionId = object.connectionId;
    } else {
      message.connectionId = "";
    }
    return message;
  },
};

const baseInterqueryResult: object = {
  creator: "",
  storeid: "",
  chainid: "",
  height: 0,
  success: false,
  proved: false,
};

export const InterqueryResult = {
  encode(message: InterqueryResult, writer: Writer = Writer.create()): Writer {
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
    if (message.height !== 0) {
      writer.uint32(40).uint64(message.height);
    }
    if (message.success === true) {
      writer.uint32(48).bool(message.success);
    }
    if (message.proved === true) {
      writer.uint32(56).bool(message.proved);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): InterqueryResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseInterqueryResult } as InterqueryResult;
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
          message.height = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.success = reader.bool();
          break;
        case 7:
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
    const message = { ...baseInterqueryResult } as InterqueryResult;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.storeid !== undefined && object.storeid !== null) {
      message.storeid = String(object.storeid);
    } else {
      message.storeid = "";
    }
    if (object.chainid !== undefined && object.chainid !== null) {
      message.chainid = String(object.chainid);
    } else {
      message.chainid = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = bytesFromBase64(object.data);
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = Number(object.height);
    } else {
      message.height = 0;
    }
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    if (object.proved !== undefined && object.proved !== null) {
      message.proved = Boolean(object.proved);
    } else {
      message.proved = false;
    }
    return message;
  },

  toJSON(message: InterqueryResult): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.storeid !== undefined && (obj.storeid = message.storeid);
    message.chainid !== undefined && (obj.chainid = message.chainid);
    message.data !== undefined &&
      (obj.data = base64FromBytes(
        message.data !== undefined ? message.data : new Uint8Array()
      ));
    message.height !== undefined && (obj.height = message.height);
    message.success !== undefined && (obj.success = message.success);
    message.proved !== undefined && (obj.proved = message.proved);
    return obj;
  },

  fromPartial(object: DeepPartial<InterqueryResult>): InterqueryResult {
    const message = { ...baseInterqueryResult } as InterqueryResult;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.storeid !== undefined && object.storeid !== null) {
      message.storeid = object.storeid;
    } else {
      message.storeid = "";
    }
    if (object.chainid !== undefined && object.chainid !== null) {
      message.chainid = object.chainid;
    } else {
      message.chainid = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = new Uint8Array();
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = 0;
    }
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    if (object.proved !== undefined && object.proved !== null) {
      message.proved = object.proved;
    } else {
      message.proved = false;
    }
    return message;
  },
};

const baseInterqueryTimeoutResult: object = {
  creator: "",
  storeid: "",
  timeoutHeight: 0,
};

export const InterqueryTimeoutResult = {
  encode(
    message: InterqueryTimeoutResult,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.storeid !== "") {
      writer.uint32(18).string(message.storeid);
    }
    if (message.timeoutHeight !== 0) {
      writer.uint32(24).uint64(message.timeoutHeight);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): InterqueryTimeoutResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseInterqueryTimeoutResult,
    } as InterqueryTimeoutResult;
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
    const message = {
      ...baseInterqueryTimeoutResult,
    } as InterqueryTimeoutResult;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.storeid !== undefined && object.storeid !== null) {
      message.storeid = String(object.storeid);
    } else {
      message.storeid = "";
    }
    if (object.timeoutHeight !== undefined && object.timeoutHeight !== null) {
      message.timeoutHeight = Number(object.timeoutHeight);
    } else {
      message.timeoutHeight = 0;
    }
    return message;
  },

  toJSON(message: InterqueryTimeoutResult): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.storeid !== undefined && (obj.storeid = message.storeid);
    message.timeoutHeight !== undefined &&
      (obj.timeoutHeight = message.timeoutHeight);
    return obj;
  },

  fromPartial(
    object: DeepPartial<InterqueryTimeoutResult>
  ): InterqueryTimeoutResult {
    const message = {
      ...baseInterqueryTimeoutResult,
    } as InterqueryTimeoutResult;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.storeid !== undefined && object.storeid !== null) {
      message.storeid = object.storeid;
    } else {
      message.storeid = "";
    }
    if (object.timeoutHeight !== undefined && object.timeoutHeight !== null) {
      message.timeoutHeight = object.timeoutHeight;
    } else {
      message.timeoutHeight = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
