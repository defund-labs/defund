/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "defundlabs.defund.broker";

export interface Source {
  poolId: number;
  interqueryId: string;
  status: string;
}

export interface Broker {
  id: string;
  connectionId: string;
  pools: Source[];
  status: string;
}

export interface Transfer {
  id: string;
  channel: string;
  sequence: number;
  status: string;
  token: Coin | undefined;
  sender: string;
  receiver: string;
  /** if we need to stake the transfer on completion or not */
  stake: boolean;
}

export interface Redeem {
  id: string;
  creator: string;
  fund: string;
  amount: Coin | undefined;
  status: string;
  type: string;
  /** if the type is staked, we need a timestamp for when it will be unstaked */
  unstakeTimestamp: string;
}

export interface Rebalance {
  id: string;
  fund: string;
  /** the height the rebalance was created */
  height: number;
  broker: string;
}

function createBaseSource(): Source {
  return { poolId: 0, interqueryId: "", status: "" };
}

export const Source = {
  encode(message: Source, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolId !== 0) {
      writer.uint32(8).uint64(message.poolId);
    }
    if (message.interqueryId !== "") {
      writer.uint32(18).string(message.interqueryId);
    }
    if (message.status !== "") {
      writer.uint32(26).string(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Source {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSource();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.interqueryId = reader.string();
          break;
        case 3:
          message.status = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Source {
    return {
      poolId: isSet(object.poolId) ? Number(object.poolId) : 0,
      interqueryId: isSet(object.interqueryId) ? String(object.interqueryId) : "",
      status: isSet(object.status) ? String(object.status) : "",
    };
  },

  toJSON(message: Source): unknown {
    const obj: any = {};
    message.poolId !== undefined && (obj.poolId = Math.round(message.poolId));
    message.interqueryId !== undefined && (obj.interqueryId = message.interqueryId);
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Source>, I>>(object: I): Source {
    const message = createBaseSource();
    message.poolId = object.poolId ?? 0;
    message.interqueryId = object.interqueryId ?? "";
    message.status = object.status ?? "";
    return message;
  },
};

function createBaseBroker(): Broker {
  return { id: "", connectionId: "", pools: [], status: "" };
}

export const Broker = {
  encode(message: Broker, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.connectionId !== "") {
      writer.uint32(18).string(message.connectionId);
    }
    for (const v of message.pools) {
      Source.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.status !== "") {
      writer.uint32(34).string(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Broker {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBroker();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.connectionId = reader.string();
          break;
        case 3:
          message.pools.push(Source.decode(reader, reader.uint32()));
          break;
        case 4:
          message.status = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Broker {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      connectionId: isSet(object.connectionId) ? String(object.connectionId) : "",
      pools: Array.isArray(object?.pools) ? object.pools.map((e: any) => Source.fromJSON(e)) : [],
      status: isSet(object.status) ? String(object.status) : "",
    };
  },

  toJSON(message: Broker): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.connectionId !== undefined && (obj.connectionId = message.connectionId);
    if (message.pools) {
      obj.pools = message.pools.map((e) => e ? Source.toJSON(e) : undefined);
    } else {
      obj.pools = [];
    }
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Broker>, I>>(object: I): Broker {
    const message = createBaseBroker();
    message.id = object.id ?? "";
    message.connectionId = object.connectionId ?? "";
    message.pools = object.pools?.map((e) => Source.fromPartial(e)) || [];
    message.status = object.status ?? "";
    return message;
  },
};

function createBaseTransfer(): Transfer {
  return { id: "", channel: "", sequence: 0, status: "", token: undefined, sender: "", receiver: "", stake: false };
}

export const Transfer = {
  encode(message: Transfer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.channel !== "") {
      writer.uint32(18).string(message.channel);
    }
    if (message.sequence !== 0) {
      writer.uint32(24).uint64(message.sequence);
    }
    if (message.status !== "") {
      writer.uint32(34).string(message.status);
    }
    if (message.token !== undefined) {
      Coin.encode(message.token, writer.uint32(42).fork()).ldelim();
    }
    if (message.sender !== "") {
      writer.uint32(50).string(message.sender);
    }
    if (message.receiver !== "") {
      writer.uint32(58).string(message.receiver);
    }
    if (message.stake === true) {
      writer.uint32(64).bool(message.stake);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Transfer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransfer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.channel = reader.string();
          break;
        case 3:
          message.sequence = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.status = reader.string();
          break;
        case 5:
          message.token = Coin.decode(reader, reader.uint32());
          break;
        case 6:
          message.sender = reader.string();
          break;
        case 7:
          message.receiver = reader.string();
          break;
        case 8:
          message.stake = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Transfer {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      channel: isSet(object.channel) ? String(object.channel) : "",
      sequence: isSet(object.sequence) ? Number(object.sequence) : 0,
      status: isSet(object.status) ? String(object.status) : "",
      token: isSet(object.token) ? Coin.fromJSON(object.token) : undefined,
      sender: isSet(object.sender) ? String(object.sender) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      stake: isSet(object.stake) ? Boolean(object.stake) : false,
    };
  },

  toJSON(message: Transfer): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.channel !== undefined && (obj.channel = message.channel);
    message.sequence !== undefined && (obj.sequence = Math.round(message.sequence));
    message.status !== undefined && (obj.status = message.status);
    message.token !== undefined && (obj.token = message.token ? Coin.toJSON(message.token) : undefined);
    message.sender !== undefined && (obj.sender = message.sender);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.stake !== undefined && (obj.stake = message.stake);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Transfer>, I>>(object: I): Transfer {
    const message = createBaseTransfer();
    message.id = object.id ?? "";
    message.channel = object.channel ?? "";
    message.sequence = object.sequence ?? 0;
    message.status = object.status ?? "";
    message.token = (object.token !== undefined && object.token !== null) ? Coin.fromPartial(object.token) : undefined;
    message.sender = object.sender ?? "";
    message.receiver = object.receiver ?? "";
    message.stake = object.stake ?? false;
    return message;
  },
};

function createBaseRedeem(): Redeem {
  return { id: "", creator: "", fund: "", amount: undefined, status: "", type: "", unstakeTimestamp: "" };
}

export const Redeem = {
  encode(message: Redeem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.fund !== "") {
      writer.uint32(26).string(message.fund);
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(34).fork()).ldelim();
    }
    if (message.status !== "") {
      writer.uint32(42).string(message.status);
    }
    if (message.type !== "") {
      writer.uint32(50).string(message.type);
    }
    if (message.unstakeTimestamp !== "") {
      writer.uint32(58).string(message.unstakeTimestamp);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Redeem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRedeem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.fund = reader.string();
          break;
        case 4:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        case 5:
          message.status = reader.string();
          break;
        case 6:
          message.type = reader.string();
          break;
        case 7:
          message.unstakeTimestamp = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Redeem {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      fund: isSet(object.fund) ? String(object.fund) : "",
      amount: isSet(object.amount) ? Coin.fromJSON(object.amount) : undefined,
      status: isSet(object.status) ? String(object.status) : "",
      type: isSet(object.type) ? String(object.type) : "",
      unstakeTimestamp: isSet(object.unstakeTimestamp) ? String(object.unstakeTimestamp) : "",
    };
  },

  toJSON(message: Redeem): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.creator !== undefined && (obj.creator = message.creator);
    message.fund !== undefined && (obj.fund = message.fund);
    message.amount !== undefined && (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.status !== undefined && (obj.status = message.status);
    message.type !== undefined && (obj.type = message.type);
    message.unstakeTimestamp !== undefined && (obj.unstakeTimestamp = message.unstakeTimestamp);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Redeem>, I>>(object: I): Redeem {
    const message = createBaseRedeem();
    message.id = object.id ?? "";
    message.creator = object.creator ?? "";
    message.fund = object.fund ?? "";
    message.amount = (object.amount !== undefined && object.amount !== null)
      ? Coin.fromPartial(object.amount)
      : undefined;
    message.status = object.status ?? "";
    message.type = object.type ?? "";
    message.unstakeTimestamp = object.unstakeTimestamp ?? "";
    return message;
  },
};

function createBaseRebalance(): Rebalance {
  return { id: "", fund: "", height: 0, broker: "" };
}

export const Rebalance = {
  encode(message: Rebalance, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.fund !== "") {
      writer.uint32(18).string(message.fund);
    }
    if (message.height !== 0) {
      writer.uint32(24).int64(message.height);
    }
    if (message.broker !== "") {
      writer.uint32(34).string(message.broker);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Rebalance {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRebalance();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.fund = reader.string();
          break;
        case 3:
          message.height = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.broker = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Rebalance {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      fund: isSet(object.fund) ? String(object.fund) : "",
      height: isSet(object.height) ? Number(object.height) : 0,
      broker: isSet(object.broker) ? String(object.broker) : "",
    };
  },

  toJSON(message: Rebalance): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.fund !== undefined && (obj.fund = message.fund);
    message.height !== undefined && (obj.height = Math.round(message.height));
    message.broker !== undefined && (obj.broker = message.broker);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Rebalance>, I>>(object: I): Rebalance {
    const message = createBaseRebalance();
    message.id = object.id ?? "";
    message.fund = object.fund ?? "";
    message.height = object.height ?? 0;
    message.broker = object.broker ?? "";
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
