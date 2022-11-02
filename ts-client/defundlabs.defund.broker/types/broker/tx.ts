/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "defundlabs.defund.broker";

export interface MsgAddLiquiditySource {
  creator: string;
  brokerId: string;
  poolId: number;
}

export interface MsgAddLiquiditySourceResponse {
}

export interface MsgAddConnectionBroker {
  creator: string;
  brokerId: string;
  connectionId: string;
}

export interface MsgAddConnectionBrokerResponse {
}

function createBaseMsgAddLiquiditySource(): MsgAddLiquiditySource {
  return { creator: "", brokerId: "", poolId: 0 };
}

export const MsgAddLiquiditySource = {
  encode(message: MsgAddLiquiditySource, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.brokerId !== "") {
      writer.uint32(18).string(message.brokerId);
    }
    if (message.poolId !== 0) {
      writer.uint32(24).uint64(message.poolId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddLiquiditySource {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddLiquiditySource();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.brokerId = reader.string();
          break;
        case 3:
          message.poolId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddLiquiditySource {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      brokerId: isSet(object.brokerId) ? String(object.brokerId) : "",
      poolId: isSet(object.poolId) ? Number(object.poolId) : 0,
    };
  },

  toJSON(message: MsgAddLiquiditySource): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.brokerId !== undefined && (obj.brokerId = message.brokerId);
    message.poolId !== undefined && (obj.poolId = Math.round(message.poolId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddLiquiditySource>, I>>(object: I): MsgAddLiquiditySource {
    const message = createBaseMsgAddLiquiditySource();
    message.creator = object.creator ?? "";
    message.brokerId = object.brokerId ?? "";
    message.poolId = object.poolId ?? 0;
    return message;
  },
};

function createBaseMsgAddLiquiditySourceResponse(): MsgAddLiquiditySourceResponse {
  return {};
}

export const MsgAddLiquiditySourceResponse = {
  encode(_: MsgAddLiquiditySourceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddLiquiditySourceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddLiquiditySourceResponse();
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

  fromJSON(_: any): MsgAddLiquiditySourceResponse {
    return {};
  },

  toJSON(_: MsgAddLiquiditySourceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddLiquiditySourceResponse>, I>>(_: I): MsgAddLiquiditySourceResponse {
    const message = createBaseMsgAddLiquiditySourceResponse();
    return message;
  },
};

function createBaseMsgAddConnectionBroker(): MsgAddConnectionBroker {
  return { creator: "", brokerId: "", connectionId: "" };
}

export const MsgAddConnectionBroker = {
  encode(message: MsgAddConnectionBroker, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.brokerId !== "") {
      writer.uint32(18).string(message.brokerId);
    }
    if (message.connectionId !== "") {
      writer.uint32(26).string(message.connectionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddConnectionBroker {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddConnectionBroker();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.brokerId = reader.string();
          break;
        case 3:
          message.connectionId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddConnectionBroker {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      brokerId: isSet(object.brokerId) ? String(object.brokerId) : "",
      connectionId: isSet(object.connectionId) ? String(object.connectionId) : "",
    };
  },

  toJSON(message: MsgAddConnectionBroker): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.brokerId !== undefined && (obj.brokerId = message.brokerId);
    message.connectionId !== undefined && (obj.connectionId = message.connectionId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddConnectionBroker>, I>>(object: I): MsgAddConnectionBroker {
    const message = createBaseMsgAddConnectionBroker();
    message.creator = object.creator ?? "";
    message.brokerId = object.brokerId ?? "";
    message.connectionId = object.connectionId ?? "";
    return message;
  },
};

function createBaseMsgAddConnectionBrokerResponse(): MsgAddConnectionBrokerResponse {
  return {};
}

export const MsgAddConnectionBrokerResponse = {
  encode(_: MsgAddConnectionBrokerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddConnectionBrokerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddConnectionBrokerResponse();
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

  fromJSON(_: any): MsgAddConnectionBrokerResponse {
    return {};
  },

  toJSON(_: MsgAddConnectionBrokerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddConnectionBrokerResponse>, I>>(_: I): MsgAddConnectionBrokerResponse {
    const message = createBaseMsgAddConnectionBrokerResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** Register defines a rpc handler for MsgAddLiquiditySource */
  AddLiquiditySource(request: MsgAddLiquiditySource): Promise<MsgAddLiquiditySourceResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddConnectionBroker(request: MsgAddConnectionBroker): Promise<MsgAddConnectionBrokerResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.AddLiquiditySource = this.AddLiquiditySource.bind(this);
    this.AddConnectionBroker = this.AddConnectionBroker.bind(this);
  }
  AddLiquiditySource(request: MsgAddLiquiditySource): Promise<MsgAddLiquiditySourceResponse> {
    const data = MsgAddLiquiditySource.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.broker.Msg", "AddLiquiditySource", data);
    return promise.then((data) => MsgAddLiquiditySourceResponse.decode(new _m0.Reader(data)));
  }

  AddConnectionBroker(request: MsgAddConnectionBroker): Promise<MsgAddConnectionBrokerResponse> {
    const data = MsgAddConnectionBroker.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.broker.Msg", "AddConnectionBroker", data);
    return promise.then((data) => MsgAddConnectionBrokerResponse.decode(new _m0.Reader(data)));
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
