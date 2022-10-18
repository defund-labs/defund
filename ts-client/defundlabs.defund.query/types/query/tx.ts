/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
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

export interface MsgCreateInterqueryResponse {}

export interface MsgCreateInterqueryResult {
  creator: string;
  storeid: string;
  /** data is submitted as a base64 encoded string but is broken down to bytes to be stored */
  data: string;
  height: Height | undefined;
  proof: ProofOps | undefined;
}

export interface MsgCreateInterqueryResultResponse {}

export interface MsgCreateInterqueryTimeout {
  creator: string;
  storeid: string;
  timeoutHeight: number;
}

export interface MsgCreateInterqueryTimeoutResponse {}

const baseMsgCreateInterquery: object = {
  creator: "",
  storeid: "",
  chainid: "",
  path: "",
  timeoutHeight: 0,
  connectionId: "",
};

export const MsgCreateInterquery = {
  encode(
    message: MsgCreateInterquery,
    writer: Writer = Writer.create()
  ): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): MsgCreateInterquery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateInterquery } as MsgCreateInterquery;
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
    const message = { ...baseMsgCreateInterquery } as MsgCreateInterquery;
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

  toJSON(message: MsgCreateInterquery): unknown {
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

  fromPartial(object: DeepPartial<MsgCreateInterquery>): MsgCreateInterquery {
    const message = { ...baseMsgCreateInterquery } as MsgCreateInterquery;
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

const baseMsgCreateInterqueryResponse: object = {};

export const MsgCreateInterqueryResponse = {
  encode(
    _: MsgCreateInterqueryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateInterqueryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateInterqueryResponse,
    } as MsgCreateInterqueryResponse;
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
    const message = {
      ...baseMsgCreateInterqueryResponse,
    } as MsgCreateInterqueryResponse;
    return message;
  },

  toJSON(_: MsgCreateInterqueryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateInterqueryResponse>
  ): MsgCreateInterqueryResponse {
    const message = {
      ...baseMsgCreateInterqueryResponse,
    } as MsgCreateInterqueryResponse;
    return message;
  },
};

const baseMsgCreateInterqueryResult: object = {
  creator: "",
  storeid: "",
  data: "",
};

export const MsgCreateInterqueryResult = {
  encode(
    message: MsgCreateInterqueryResult,
    writer: Writer = Writer.create()
  ): Writer {
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

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateInterqueryResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateInterqueryResult,
    } as MsgCreateInterqueryResult;
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
    const message = {
      ...baseMsgCreateInterqueryResult,
    } as MsgCreateInterqueryResult;
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
    if (object.data !== undefined && object.data !== null) {
      message.data = String(object.data);
    } else {
      message.data = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = Height.fromJSON(object.height);
    } else {
      message.height = undefined;
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = ProofOps.fromJSON(object.proof);
    } else {
      message.proof = undefined;
    }
    return message;
  },

  toJSON(message: MsgCreateInterqueryResult): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.storeid !== undefined && (obj.storeid = message.storeid);
    message.data !== undefined && (obj.data = message.data);
    message.height !== undefined &&
      (obj.height = message.height ? Height.toJSON(message.height) : undefined);
    message.proof !== undefined &&
      (obj.proof = message.proof ? ProofOps.toJSON(message.proof) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateInterqueryResult>
  ): MsgCreateInterqueryResult {
    const message = {
      ...baseMsgCreateInterqueryResult,
    } as MsgCreateInterqueryResult;
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
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = Height.fromPartial(object.height);
    } else {
      message.height = undefined;
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = ProofOps.fromPartial(object.proof);
    } else {
      message.proof = undefined;
    }
    return message;
  },
};

const baseMsgCreateInterqueryResultResponse: object = {};

export const MsgCreateInterqueryResultResponse = {
  encode(
    _: MsgCreateInterqueryResultResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateInterqueryResultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateInterqueryResultResponse,
    } as MsgCreateInterqueryResultResponse;
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
    const message = {
      ...baseMsgCreateInterqueryResultResponse,
    } as MsgCreateInterqueryResultResponse;
    return message;
  },

  toJSON(_: MsgCreateInterqueryResultResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateInterqueryResultResponse>
  ): MsgCreateInterqueryResultResponse {
    const message = {
      ...baseMsgCreateInterqueryResultResponse,
    } as MsgCreateInterqueryResultResponse;
    return message;
  },
};

const baseMsgCreateInterqueryTimeout: object = {
  creator: "",
  storeid: "",
  timeoutHeight: 0,
};

export const MsgCreateInterqueryTimeout = {
  encode(
    message: MsgCreateInterqueryTimeout,
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

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateInterqueryTimeout {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateInterqueryTimeout,
    } as MsgCreateInterqueryTimeout;
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
    const message = {
      ...baseMsgCreateInterqueryTimeout,
    } as MsgCreateInterqueryTimeout;
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

  toJSON(message: MsgCreateInterqueryTimeout): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.storeid !== undefined && (obj.storeid = message.storeid);
    message.timeoutHeight !== undefined &&
      (obj.timeoutHeight = message.timeoutHeight);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateInterqueryTimeout>
  ): MsgCreateInterqueryTimeout {
    const message = {
      ...baseMsgCreateInterqueryTimeout,
    } as MsgCreateInterqueryTimeout;
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

const baseMsgCreateInterqueryTimeoutResponse: object = {};

export const MsgCreateInterqueryTimeoutResponse = {
  encode(
    _: MsgCreateInterqueryTimeoutResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateInterqueryTimeoutResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateInterqueryTimeoutResponse,
    } as MsgCreateInterqueryTimeoutResponse;
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
    const message = {
      ...baseMsgCreateInterqueryTimeoutResponse,
    } as MsgCreateInterqueryTimeoutResponse;
    return message;
  },

  toJSON(_: MsgCreateInterqueryTimeoutResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateInterqueryTimeoutResponse>
  ): MsgCreateInterqueryTimeoutResponse {
    const message = {
      ...baseMsgCreateInterqueryTimeoutResponse,
    } as MsgCreateInterqueryTimeoutResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateInterquery(
    request: MsgCreateInterquery
  ): Promise<MsgCreateInterqueryResponse>;
  CreateInterqueryResult(
    request: MsgCreateInterqueryResult
  ): Promise<MsgCreateInterqueryResultResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateInterqueryTimeout(
    request: MsgCreateInterqueryTimeout
  ): Promise<MsgCreateInterqueryTimeoutResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateInterquery(
    request: MsgCreateInterquery
  ): Promise<MsgCreateInterqueryResponse> {
    const data = MsgCreateInterquery.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.query.Msg",
      "CreateInterquery",
      data
    );
    return promise.then((data) =>
      MsgCreateInterqueryResponse.decode(new Reader(data))
    );
  }

  CreateInterqueryResult(
    request: MsgCreateInterqueryResult
  ): Promise<MsgCreateInterqueryResultResponse> {
    const data = MsgCreateInterqueryResult.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.query.Msg",
      "CreateInterqueryResult",
      data
    );
    return promise.then((data) =>
      MsgCreateInterqueryResultResponse.decode(new Reader(data))
    );
  }

  CreateInterqueryTimeout(
    request: MsgCreateInterqueryTimeout
  ): Promise<MsgCreateInterqueryTimeoutResponse> {
    const data = MsgCreateInterqueryTimeout.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.query.Msg",
      "CreateInterqueryTimeout",
      data
    );
    return promise.then((data) =>
      MsgCreateInterqueryTimeoutResponse.decode(new Reader(data))
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
