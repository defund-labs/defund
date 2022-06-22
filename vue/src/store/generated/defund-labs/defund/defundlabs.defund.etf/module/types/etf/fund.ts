/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Broker } from "../broker/broker";

export const protobufPackage = "defundlabs.defund.etf";

export interface FundPrice {
  id: string;
  height: number;
  amount: Coin | undefined;
  symbol: string;
}

export interface Holding {
  token: string;
  percent: number;
  /** Pool ID of the Pool for this holding on Broker */
  poolId: string;
}

export interface Fund {
  symbol: string;
  address: string;
  name: string;
  description: string;
  shares: Coin | undefined;
  broker: Broker | undefined;
  holdings: Holding[];
  rebalance: number;
  baseDenom: string;
  connectionId: string;
  creator: string;
}

export interface Invest {
  id: string;
  creator: string;
  fund: Fund | undefined;
  amount: Coin | undefined;
  channel: string;
  sequence: string;
  status: string;
}

export interface Uninvest {
  id: string;
  creator: string;
  fund: Fund | undefined;
  amount: Coin | undefined;
  channel: string;
  sequence: string;
  status: string;
}

const baseFundPrice: object = { id: "", height: 0, symbol: "" };

export const FundPrice = {
  encode(message: FundPrice, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.height !== 0) {
      writer.uint32(16).uint64(message.height);
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(26).fork()).ldelim();
    }
    if (message.symbol !== "") {
      writer.uint32(34).string(message.symbol);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): FundPrice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFundPrice } as FundPrice;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.height = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.symbol = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FundPrice {
    const message = { ...baseFundPrice } as FundPrice;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = Number(object.height);
    } else {
      message.height = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromJSON(object.amount);
    } else {
      message.amount = undefined;
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = String(object.symbol);
    } else {
      message.symbol = "";
    }
    return message;
  },

  toJSON(message: FundPrice): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.height !== undefined && (obj.height = message.height);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    return obj;
  },

  fromPartial(object: DeepPartial<FundPrice>): FundPrice {
    const message = { ...baseFundPrice } as FundPrice;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromPartial(object.amount);
    } else {
      message.amount = undefined;
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = object.symbol;
    } else {
      message.symbol = "";
    }
    return message;
  },
};

const baseHolding: object = { token: "", percent: 0, poolId: "" };

export const Holding = {
  encode(message: Holding, writer: Writer = Writer.create()): Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.percent !== 0) {
      writer.uint32(16).int64(message.percent);
    }
    if (message.poolId !== "") {
      writer.uint32(26).string(message.poolId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Holding {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseHolding } as Holding;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        case 2:
          message.percent = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.poolId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Holding {
    const message = { ...baseHolding } as Holding;
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.percent !== undefined && object.percent !== null) {
      message.percent = Number(object.percent);
    } else {
      message.percent = 0;
    }
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = String(object.poolId);
    } else {
      message.poolId = "";
    }
    return message;
  },

  toJSON(message: Holding): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    message.percent !== undefined && (obj.percent = message.percent);
    message.poolId !== undefined && (obj.poolId = message.poolId);
    return obj;
  },

  fromPartial(object: DeepPartial<Holding>): Holding {
    const message = { ...baseHolding } as Holding;
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.percent !== undefined && object.percent !== null) {
      message.percent = object.percent;
    } else {
      message.percent = 0;
    }
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = object.poolId;
    } else {
      message.poolId = "";
    }
    return message;
  },
};

const baseFund: object = {
  symbol: "",
  address: "",
  name: "",
  description: "",
  rebalance: 0,
  baseDenom: "",
  connectionId: "",
  creator: "",
};

export const Fund = {
  encode(message: Fund, writer: Writer = Writer.create()): Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    if (message.shares !== undefined) {
      Coin.encode(message.shares, writer.uint32(42).fork()).ldelim();
    }
    if (message.broker !== undefined) {
      Broker.encode(message.broker, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.holdings) {
      Holding.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    if (message.rebalance !== 0) {
      writer.uint32(64).int64(message.rebalance);
    }
    if (message.baseDenom !== "") {
      writer.uint32(74).string(message.baseDenom);
    }
    if (message.connectionId !== "") {
      writer.uint32(82).string(message.connectionId);
    }
    if (message.creator !== "") {
      writer.uint32(90).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Fund {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFund } as Fund;
    message.holdings = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbol = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.name = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.shares = Coin.decode(reader, reader.uint32());
          break;
        case 6:
          message.broker = Broker.decode(reader, reader.uint32());
          break;
        case 7:
          message.holdings.push(Holding.decode(reader, reader.uint32()));
          break;
        case 8:
          message.rebalance = longToNumber(reader.int64() as Long);
          break;
        case 9:
          message.baseDenom = reader.string();
          break;
        case 10:
          message.connectionId = reader.string();
          break;
        case 11:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Fund {
    const message = { ...baseFund } as Fund;
    message.holdings = [];
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = String(object.symbol);
    } else {
      message.symbol = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = Coin.fromJSON(object.shares);
    } else {
      message.shares = undefined;
    }
    if (object.broker !== undefined && object.broker !== null) {
      message.broker = Broker.fromJSON(object.broker);
    } else {
      message.broker = undefined;
    }
    if (object.holdings !== undefined && object.holdings !== null) {
      for (const e of object.holdings) {
        message.holdings.push(Holding.fromJSON(e));
      }
    }
    if (object.rebalance !== undefined && object.rebalance !== null) {
      message.rebalance = Number(object.rebalance);
    } else {
      message.rebalance = 0;
    }
    if (object.baseDenom !== undefined && object.baseDenom !== null) {
      message.baseDenom = String(object.baseDenom);
    } else {
      message.baseDenom = "";
    }
    if (object.connectionId !== undefined && object.connectionId !== null) {
      message.connectionId = String(object.connectionId);
    } else {
      message.connectionId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: Fund): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.address !== undefined && (obj.address = message.address);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    message.shares !== undefined &&
      (obj.shares = message.shares ? Coin.toJSON(message.shares) : undefined);
    message.broker !== undefined &&
      (obj.broker = message.broker ? Broker.toJSON(message.broker) : undefined);
    if (message.holdings) {
      obj.holdings = message.holdings.map((e) =>
        e ? Holding.toJSON(e) : undefined
      );
    } else {
      obj.holdings = [];
    }
    message.rebalance !== undefined && (obj.rebalance = message.rebalance);
    message.baseDenom !== undefined && (obj.baseDenom = message.baseDenom);
    message.connectionId !== undefined &&
      (obj.connectionId = message.connectionId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Fund>): Fund {
    const message = { ...baseFund } as Fund;
    message.holdings = [];
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = object.symbol;
    } else {
      message.symbol = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = Coin.fromPartial(object.shares);
    } else {
      message.shares = undefined;
    }
    if (object.broker !== undefined && object.broker !== null) {
      message.broker = Broker.fromPartial(object.broker);
    } else {
      message.broker = undefined;
    }
    if (object.holdings !== undefined && object.holdings !== null) {
      for (const e of object.holdings) {
        message.holdings.push(Holding.fromPartial(e));
      }
    }
    if (object.rebalance !== undefined && object.rebalance !== null) {
      message.rebalance = object.rebalance;
    } else {
      message.rebalance = 0;
    }
    if (object.baseDenom !== undefined && object.baseDenom !== null) {
      message.baseDenom = object.baseDenom;
    } else {
      message.baseDenom = "";
    }
    if (object.connectionId !== undefined && object.connectionId !== null) {
      message.connectionId = object.connectionId;
    } else {
      message.connectionId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseInvest: object = {
  id: "",
  creator: "",
  channel: "",
  sequence: "",
  status: "",
};

export const Invest = {
  encode(message: Invest, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.fund !== undefined) {
      Fund.encode(message.fund, writer.uint32(26).fork()).ldelim();
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(34).fork()).ldelim();
    }
    if (message.channel !== "") {
      writer.uint32(42).string(message.channel);
    }
    if (message.sequence !== "") {
      writer.uint32(50).string(message.sequence);
    }
    if (message.status !== "") {
      writer.uint32(58).string(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Invest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseInvest } as Invest;
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
          message.fund = Fund.decode(reader, reader.uint32());
          break;
        case 4:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        case 5:
          message.channel = reader.string();
          break;
        case 6:
          message.sequence = reader.string();
          break;
        case 7:
          message.status = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Invest {
    const message = { ...baseInvest } as Invest;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.fund !== undefined && object.fund !== null) {
      message.fund = Fund.fromJSON(object.fund);
    } else {
      message.fund = undefined;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromJSON(object.amount);
    } else {
      message.amount = undefined;
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = String(object.channel);
    } else {
      message.channel = "";
    }
    if (object.sequence !== undefined && object.sequence !== null) {
      message.sequence = String(object.sequence);
    } else {
      message.sequence = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    return message;
  },

  toJSON(message: Invest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.creator !== undefined && (obj.creator = message.creator);
    message.fund !== undefined &&
      (obj.fund = message.fund ? Fund.toJSON(message.fund) : undefined);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.channel !== undefined && (obj.channel = message.channel);
    message.sequence !== undefined && (obj.sequence = message.sequence);
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(object: DeepPartial<Invest>): Invest {
    const message = { ...baseInvest } as Invest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.fund !== undefined && object.fund !== null) {
      message.fund = Fund.fromPartial(object.fund);
    } else {
      message.fund = undefined;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromPartial(object.amount);
    } else {
      message.amount = undefined;
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = object.channel;
    } else {
      message.channel = "";
    }
    if (object.sequence !== undefined && object.sequence !== null) {
      message.sequence = object.sequence;
    } else {
      message.sequence = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    return message;
  },
};

const baseUninvest: object = {
  id: "",
  creator: "",
  channel: "",
  sequence: "",
  status: "",
};

export const Uninvest = {
  encode(message: Uninvest, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.fund !== undefined) {
      Fund.encode(message.fund, writer.uint32(26).fork()).ldelim();
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(34).fork()).ldelim();
    }
    if (message.channel !== "") {
      writer.uint32(42).string(message.channel);
    }
    if (message.sequence !== "") {
      writer.uint32(50).string(message.sequence);
    }
    if (message.status !== "") {
      writer.uint32(58).string(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Uninvest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUninvest } as Uninvest;
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
          message.fund = Fund.decode(reader, reader.uint32());
          break;
        case 4:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        case 5:
          message.channel = reader.string();
          break;
        case 6:
          message.sequence = reader.string();
          break;
        case 7:
          message.status = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Uninvest {
    const message = { ...baseUninvest } as Uninvest;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.fund !== undefined && object.fund !== null) {
      message.fund = Fund.fromJSON(object.fund);
    } else {
      message.fund = undefined;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromJSON(object.amount);
    } else {
      message.amount = undefined;
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = String(object.channel);
    } else {
      message.channel = "";
    }
    if (object.sequence !== undefined && object.sequence !== null) {
      message.sequence = String(object.sequence);
    } else {
      message.sequence = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    return message;
  },

  toJSON(message: Uninvest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.creator !== undefined && (obj.creator = message.creator);
    message.fund !== undefined &&
      (obj.fund = message.fund ? Fund.toJSON(message.fund) : undefined);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.channel !== undefined && (obj.channel = message.channel);
    message.sequence !== undefined && (obj.sequence = message.sequence);
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(object: DeepPartial<Uninvest>): Uninvest {
    const message = { ...baseUninvest } as Uninvest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.fund !== undefined && object.fund !== null) {
      message.fund = Fund.fromPartial(object.fund);
    } else {
      message.fund = undefined;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromPartial(object.amount);
    } else {
      message.amount = undefined;
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = object.channel;
    } else {
      message.channel = "";
    }
    if (object.sequence !== undefined && object.sequence !== null) {
      message.sequence = object.sequence;
    } else {
      message.sequence = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
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
