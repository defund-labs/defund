/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Height } from "../ibc/core/client/v1/client";
import { ProofOps } from "../tendermint/crypto/proof";

export const protobufPackage = "defundlabs.defund.query";

export interface MsgCreateInterquery {
  creator: string;
  storeid: string;
  chainid: string;
  path: string;
  key: Uint8Array;
  timeoutHeight: number;
  connectionId: string;
}

export interface MsgCreateInterqueryResponse {
}

export interface MsgCreateInterqueryResult {
  creator: string;
  storeid: string;
  /** data is submitted as a base64 encoded string but is broken down to bytes to be stored */
  data: string;
  height: Height | undefined;
  proof: ProofOps | undefined;
}

export interface MsgCreateInterqueryResultResponse {
}

export interface MsgCreateInterqueryTimeout {
  creator: string;
  storeid: string;
  timeoutHeight: number;
}

export interface MsgCreateInterqueryTimeoutResponse {
}

function createBaseMsgCreateInterquery(): MsgCreateInterquery {
  return { creator: "", storeid: "", chainid: "", path: "", key: new Uint8Array(), timeoutHeight: 0, connectionId: "" };
}

export const MsgCreateInterquery = {
  encode(message: MsgCreateInterquery, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateInterquery {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateInterquery();
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

  fromJSON(object: any): MsgCreateInterquery {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      storeid: isSet(object.storeid) ? String(object.storeid) : "",
      chainid: isSet(object.chainid) ? String(object.chainid) : "",
      path: isSet(object.path) ? String(object.path) : "",
      key: isSet(object.key) ? bytesFromBase64(object.key) : new Uint8Array(),
      timeoutHeight: isSet(object.timeoutHeight) ? Number(object.timeoutHeight) : 0,
      connectionId: isSet(object.connectionId) ? String(object.connectionId) : "",
    };
  },

  toJSON(message: MsgCreateInterquery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.storeid !== undefined && (obj.storeid = message.storeid);
    message.chainid !== undefined && (obj.chainid = message.chainid);
    message.path !== undefined && (obj.path = message.path);
    message.key !== undefined
      && (obj.key = base64FromBytes(message.key !== undefined ? message.key : new Uint8Array()));
    message.timeoutHeight !== undefined && (obj.timeoutHeight = Math.round(message.timeoutHeight));
    message.connectionId !== undefined && (obj.connectionId = message.connectionId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateInterquery>, I>>(object: I): MsgCreateInterquery {
    const message = createBaseMsgCreateInterquery();
    message.creator = object.creator ?? "";
    message.storeid = object.storeid ?? "";
    message.chainid = object.chainid ?? "";
    message.path = object.path ?? "";
    message.key = object.key ?? new Uint8Array();
    message.timeoutHeight = object.timeoutHeight ?? 0;
    message.connectionId = object.connectionId ?? "";
    return message;
  },
};

function createBaseMsgCreateInterqueryResponse(): MsgCreateInterqueryResponse {
  return {};
}

export const MsgCreateInterqueryResponse = {
  encode(_: MsgCreateInterqueryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateInterqueryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateInterqueryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateInterqueryResponse {
    return {};
  },

  toJSON(_: MsgCreateInterqueryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateInterqueryResponse>, I>>(_: I): MsgCreateInterqueryResponse {
    const message = createBaseMsgCreateInterqueryResponse();
    return message;
  },
};

function createBaseMsgCreateInterqueryResult(): MsgCreateInterqueryResult {
  return { creator: "", storeid: "", data: "", height: undefined, proof: undefined };
}

export const MsgCreateInterqueryResult = {
  encode(message: MsgCreateInterqueryResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.storeid !== "") {
      writer.uint32(18).string(message.storeid);
    }
    if (message.data !== "") {
      writer.uint32(26).string(message.data);
    }
    if (message.height !== undefined) {
      Height.encode(message.height, writer.uint32(34).fork()).ldelim();
    }
    if (message.proof !== undefined) {
      ProofOps.encode(message.proof, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateInterqueryResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateInterqueryResult();
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
          message.data = reader.string();
          break;
        case 4:
          message.height = Height.decode(reader, reader.uint32());
          break;
        case 5:
          message.proof = ProofOps.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateInterqueryResult {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      storeid: isSet(object.storeid) ? String(object.storeid) : "",
      data: isSet(object.data) ? String(object.data) : "",
      height: isSet(object.height) ? Height.fromJSON(object.height) : undefined,
      proof: isSet(object.proof) ? ProofOps.fromJSON(object.proof) : undefined,
    };
  },

  toJSON(message: MsgCreateInterqueryResult): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.storeid !== undefined && (obj.storeid = message.storeid);
    message.data !== undefined && (obj.data = message.data);
    message.height !== undefined && (obj.height = message.height ? Height.toJSON(message.height) : undefined);
    message.proof !== undefined && (obj.proof = message.proof ? ProofOps.toJSON(message.proof) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateInterqueryResult>, I>>(object: I): MsgCreateInterqueryResult {
    const message = createBaseMsgCreateInterqueryResult();
    message.creator = object.creator ?? "";
    message.storeid = object.storeid ?? "";
    message.data = object.data ?? "";
    message.height = (object.height !== undefined && object.height !== null)
      ? Height.fromPartial(object.height)
      : undefined;
    message.proof = (object.proof !== undefined && object.proof !== null)
      ? ProofOps.fromPartial(object.proof)
      : undefined;
    return message;
  },
};

function createBaseMsgCreateInterqueryResultResponse(): MsgCreateInterqueryResultResponse {
  return {};
}

export const MsgCreateInterqueryResultResponse = {
  encode(_: MsgCreateInterqueryResultResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateInterqueryResultResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateInterqueryResultResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateInterqueryResultResponse {
    return {};
  },

  toJSON(_: MsgCreateInterqueryResultResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateInterqueryResultResponse>, I>>(
    _: I,
  ): MsgCreateInterqueryResultResponse {
    const message = createBaseMsgCreateInterqueryResultResponse();
    return message;
  },
};

function createBaseMsgCreateInterqueryTimeout(): MsgCreateInterqueryTimeout {
  return { creator: "", storeid: "", timeoutHeight: 0 };
}

export const MsgCreateInterqueryTimeout = {
  encode(message: MsgCreateInterqueryTimeout, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateInterqueryTimeout {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateInterqueryTimeout();
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

  fromJSON(object: any): MsgCreateInterqueryTimeout {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      storeid: isSet(object.storeid) ? String(object.storeid) : "",
      timeoutHeight: isSet(object.timeoutHeight) ? Number(object.timeoutHeight) : 0,
    };
  },

  toJSON(message: MsgCreateInterqueryTimeout): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.storeid !== undefined && (obj.storeid = message.storeid);
    message.timeoutHeight !== undefined && (obj.timeoutHeight = Math.round(message.timeoutHeight));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateInterqueryTimeout>, I>>(object: I): MsgCreateInterqueryTimeout {
    const message = createBaseMsgCreateInterqueryTimeout();
    message.creator = object.creator ?? "";
    message.storeid = object.storeid ?? "";
    message.timeoutHeight = object.timeoutHeight ?? 0;
    return message;
  },
};

function createBaseMsgCreateInterqueryTimeoutResponse(): MsgCreateInterqueryTimeoutResponse {
  return {};
}

export const MsgCreateInterqueryTimeoutResponse = {
  encode(_: MsgCreateInterqueryTimeoutResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateInterqueryTimeoutResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateInterqueryTimeoutResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateInterqueryTimeoutResponse {
    return {};
  },

  toJSON(_: MsgCreateInterqueryTimeoutResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateInterqueryTimeoutResponse>, I>>(
    _: I,
  ): MsgCreateInterqueryTimeoutResponse {
    const message = createBaseMsgCreateInterqueryTimeoutResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateInterquery(request: MsgCreateInterquery): Promise<MsgCreateInterqueryResponse>;
  CreateInterqueryResult(request: MsgCreateInterqueryResult): Promise<MsgCreateInterqueryResultResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateInterqueryTimeout(request: MsgCreateInterqueryTimeout): Promise<MsgCreateInterqueryTimeoutResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateInterquery = this.CreateInterquery.bind(this);
    this.CreateInterqueryResult = this.CreateInterqueryResult.bind(this);
    this.CreateInterqueryTimeout = this.CreateInterqueryTimeout.bind(this);
  }
  CreateInterquery(request: MsgCreateInterquery): Promise<MsgCreateInterqueryResponse> {
    const data = MsgCreateInterquery.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.query.Msg", "CreateInterquery", data);
    return promise.then((data) => MsgCreateInterqueryResponse.decode(new _m0.Reader(data)));
  }

  CreateInterqueryResult(request: MsgCreateInterqueryResult): Promise<MsgCreateInterqueryResultResponse> {
    const data = MsgCreateInterqueryResult.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.query.Msg", "CreateInterqueryResult", data);
    return promise.then((data) => MsgCreateInterqueryResultResponse.decode(new _m0.Reader(data)));
  }

  CreateInterqueryTimeout(request: MsgCreateInterqueryTimeout): Promise<MsgCreateInterqueryTimeoutResponse> {
    const data = MsgCreateInterqueryTimeout.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.query.Msg", "CreateInterqueryTimeout", data);
    return promise.then((data) => MsgCreateInterqueryTimeoutResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
