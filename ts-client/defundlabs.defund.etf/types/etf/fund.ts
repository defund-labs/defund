/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Timestamp } from "../google/protobuf/timestamp";

export const protobufPackage = "defundlabs.defund.etf";

export interface BaseDenom {
  onDefund: string;
  onBroker: string;
}

export interface FundPrice {
  id: string;
  height: number;
  time: Date | undefined;
  amount: Coin | undefined;
  symbol: string;
}

export interface Balances {
  balances: Coin[];
}

export interface Holding {
  token: string;
  percent: number;
  /** Pool ID of the Pool for this holding on Broker */
  poolId: number;
  /** Broker Id for the Broker */
  brokerId: string;
  /** the type of the asset. Valid types are spot, staked */
  type: string;
}

export interface Fund {
  symbol: string;
  address: string;
  name: string;
  description: string;
  shares: Coin | undefined;
  holdings: Holding[];
  rebalance: number;
  baseDenom: BaseDenom | undefined;
  startingPrice: Coin | undefined;
  creator: string;
  rebalancing: boolean;
  lastRebalanceHeight: number;
  balances: { [key: string]: Balances };
}

export interface Fund_BalancesEntry {
  key: string;
  value: Balances | undefined;
}

function createBaseBaseDenom(): BaseDenom {
  return { onDefund: "", onBroker: "" };
}

export const BaseDenom = {
  encode(message: BaseDenom, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.onDefund !== "") {
      writer.uint32(10).string(message.onDefund);
    }
    if (message.onBroker !== "") {
      writer.uint32(18).string(message.onBroker);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BaseDenom {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBaseDenom();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.onDefund = reader.string();
          break;
        case 2:
          message.onBroker = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BaseDenom {
    return {
      onDefund: isSet(object.onDefund) ? String(object.onDefund) : "",
      onBroker: isSet(object.onBroker) ? String(object.onBroker) : "",
    };
  },

  toJSON(message: BaseDenom): unknown {
    const obj: any = {};
    message.onDefund !== undefined && (obj.onDefund = message.onDefund);
    message.onBroker !== undefined && (obj.onBroker = message.onBroker);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BaseDenom>, I>>(object: I): BaseDenom {
    const message = createBaseBaseDenom();
    message.onDefund = object.onDefund ?? "";
    message.onBroker = object.onBroker ?? "";
    return message;
  },
};

function createBaseFundPrice(): FundPrice {
  return { id: "", height: 0, time: undefined, amount: undefined, symbol: "" };
}

export const FundPrice = {
  encode(message: FundPrice, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.height !== 0) {
      writer.uint32(16).int64(message.height);
    }
    if (message.time !== undefined) {
      Timestamp.encode(toTimestamp(message.time), writer.uint32(26).fork()).ldelim();
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(34).fork()).ldelim();
    }
    if (message.symbol !== "") {
      writer.uint32(42).string(message.symbol);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FundPrice {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFundPrice();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.height = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.time = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 4:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        case 5:
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
    return {
      id: isSet(object.id) ? String(object.id) : "",
      height: isSet(object.height) ? Number(object.height) : 0,
      time: isSet(object.time) ? fromJsonTimestamp(object.time) : undefined,
      amount: isSet(object.amount) ? Coin.fromJSON(object.amount) : undefined,
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
    };
  },

  toJSON(message: FundPrice): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.height !== undefined && (obj.height = Math.round(message.height));
    message.time !== undefined && (obj.time = message.time.toISOString());
    message.amount !== undefined && (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FundPrice>, I>>(object: I): FundPrice {
    const message = createBaseFundPrice();
    message.id = object.id ?? "";
    message.height = object.height ?? 0;
    message.time = object.time ?? undefined;
    message.amount = (object.amount !== undefined && object.amount !== null)
      ? Coin.fromPartial(object.amount)
      : undefined;
    message.symbol = object.symbol ?? "";
    return message;
  },
};

function createBaseBalances(): Balances {
  return { balances: [] };
}

export const Balances = {
  encode(message: Balances, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.balances) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Balances {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBalances();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.balances.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Balances {
    return { balances: Array.isArray(object?.balances) ? object.balances.map((e: any) => Coin.fromJSON(e)) : [] };
  },

  toJSON(message: Balances): unknown {
    const obj: any = {};
    if (message.balances) {
      obj.balances = message.balances.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.balances = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Balances>, I>>(object: I): Balances {
    const message = createBaseBalances();
    message.balances = object.balances?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseHolding(): Holding {
  return { token: "", percent: 0, poolId: 0, brokerId: "", type: "" };
}

export const Holding = {
  encode(message: Holding, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.percent !== 0) {
      writer.uint32(16).int64(message.percent);
    }
    if (message.poolId !== 0) {
      writer.uint32(24).uint64(message.poolId);
    }
    if (message.brokerId !== "") {
      writer.uint32(34).string(message.brokerId);
    }
    if (message.type !== "") {
      writer.uint32(42).string(message.type);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Holding {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHolding();
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
          message.poolId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.brokerId = reader.string();
          break;
        case 5:
          message.type = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Holding {
    return {
      token: isSet(object.token) ? String(object.token) : "",
      percent: isSet(object.percent) ? Number(object.percent) : 0,
      poolId: isSet(object.poolId) ? Number(object.poolId) : 0,
      brokerId: isSet(object.brokerId) ? String(object.brokerId) : "",
      type: isSet(object.type) ? String(object.type) : "",
    };
  },

  toJSON(message: Holding): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    message.percent !== undefined && (obj.percent = Math.round(message.percent));
    message.poolId !== undefined && (obj.poolId = Math.round(message.poolId));
    message.brokerId !== undefined && (obj.brokerId = message.brokerId);
    message.type !== undefined && (obj.type = message.type);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Holding>, I>>(object: I): Holding {
    const message = createBaseHolding();
    message.token = object.token ?? "";
    message.percent = object.percent ?? 0;
    message.poolId = object.poolId ?? 0;
    message.brokerId = object.brokerId ?? "";
    message.type = object.type ?? "";
    return message;
  },
};

function createBaseFund(): Fund {
  return {
    symbol: "",
    address: "",
    name: "",
    description: "",
    shares: undefined,
    holdings: [],
    rebalance: 0,
    baseDenom: undefined,
    startingPrice: undefined,
    creator: "",
    rebalancing: false,
    lastRebalanceHeight: 0,
    balances: {},
  };
}

export const Fund = {
  encode(message: Fund, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    for (const v of message.holdings) {
      Holding.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    if (message.rebalance !== 0) {
      writer.uint32(56).int64(message.rebalance);
    }
    if (message.baseDenom !== undefined) {
      BaseDenom.encode(message.baseDenom, writer.uint32(66).fork()).ldelim();
    }
    if (message.startingPrice !== undefined) {
      Coin.encode(message.startingPrice, writer.uint32(74).fork()).ldelim();
    }
    if (message.creator !== "") {
      writer.uint32(82).string(message.creator);
    }
    if (message.rebalancing === true) {
      writer.uint32(88).bool(message.rebalancing);
    }
    if (message.lastRebalanceHeight !== 0) {
      writer.uint32(96).int64(message.lastRebalanceHeight);
    }
    Object.entries(message.balances).forEach(([key, value]) => {
      Fund_BalancesEntry.encode({ key: key as any, value }, writer.uint32(106).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Fund {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFund();
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
          message.holdings.push(Holding.decode(reader, reader.uint32()));
          break;
        case 7:
          message.rebalance = longToNumber(reader.int64() as Long);
          break;
        case 8:
          message.baseDenom = BaseDenom.decode(reader, reader.uint32());
          break;
        case 9:
          message.startingPrice = Coin.decode(reader, reader.uint32());
          break;
        case 10:
          message.creator = reader.string();
          break;
        case 11:
          message.rebalancing = reader.bool();
          break;
        case 12:
          message.lastRebalanceHeight = longToNumber(reader.int64() as Long);
          break;
        case 13:
          const entry13 = Fund_BalancesEntry.decode(reader, reader.uint32());
          if (entry13.value !== undefined) {
            message.balances[entry13.key] = entry13.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Fund {
    return {
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      address: isSet(object.address) ? String(object.address) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      shares: isSet(object.shares) ? Coin.fromJSON(object.shares) : undefined,
      holdings: Array.isArray(object?.holdings) ? object.holdings.map((e: any) => Holding.fromJSON(e)) : [],
      rebalance: isSet(object.rebalance) ? Number(object.rebalance) : 0,
      baseDenom: isSet(object.baseDenom) ? BaseDenom.fromJSON(object.baseDenom) : undefined,
      startingPrice: isSet(object.startingPrice) ? Coin.fromJSON(object.startingPrice) : undefined,
      creator: isSet(object.creator) ? String(object.creator) : "",
      rebalancing: isSet(object.rebalancing) ? Boolean(object.rebalancing) : false,
      lastRebalanceHeight: isSet(object.lastRebalanceHeight) ? Number(object.lastRebalanceHeight) : 0,
      balances: isObject(object.balances)
        ? Object.entries(object.balances).reduce<{ [key: string]: Balances }>((acc, [key, value]) => {
          acc[key] = Balances.fromJSON(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: Fund): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.address !== undefined && (obj.address = message.address);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.shares !== undefined && (obj.shares = message.shares ? Coin.toJSON(message.shares) : undefined);
    if (message.holdings) {
      obj.holdings = message.holdings.map((e) => e ? Holding.toJSON(e) : undefined);
    } else {
      obj.holdings = [];
    }
    message.rebalance !== undefined && (obj.rebalance = Math.round(message.rebalance));
    message.baseDenom !== undefined
      && (obj.baseDenom = message.baseDenom ? BaseDenom.toJSON(message.baseDenom) : undefined);
    message.startingPrice !== undefined
      && (obj.startingPrice = message.startingPrice ? Coin.toJSON(message.startingPrice) : undefined);
    message.creator !== undefined && (obj.creator = message.creator);
    message.rebalancing !== undefined && (obj.rebalancing = message.rebalancing);
    message.lastRebalanceHeight !== undefined && (obj.lastRebalanceHeight = Math.round(message.lastRebalanceHeight));
    obj.balances = {};
    if (message.balances) {
      Object.entries(message.balances).forEach(([k, v]) => {
        obj.balances[k] = Balances.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Fund>, I>>(object: I): Fund {
    const message = createBaseFund();
    message.symbol = object.symbol ?? "";
    message.address = object.address ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.shares = (object.shares !== undefined && object.shares !== null)
      ? Coin.fromPartial(object.shares)
      : undefined;
    message.holdings = object.holdings?.map((e) => Holding.fromPartial(e)) || [];
    message.rebalance = object.rebalance ?? 0;
    message.baseDenom = (object.baseDenom !== undefined && object.baseDenom !== null)
      ? BaseDenom.fromPartial(object.baseDenom)
      : undefined;
    message.startingPrice = (object.startingPrice !== undefined && object.startingPrice !== null)
      ? Coin.fromPartial(object.startingPrice)
      : undefined;
    message.creator = object.creator ?? "";
    message.rebalancing = object.rebalancing ?? false;
    message.lastRebalanceHeight = object.lastRebalanceHeight ?? 0;
    message.balances = Object.entries(object.balances ?? {}).reduce<{ [key: string]: Balances }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = Balances.fromPartial(value);
        }
        return acc;
      },
      {},
    );
    return message;
  },
};

function createBaseFund_BalancesEntry(): Fund_BalancesEntry {
  return { key: "", value: undefined };
}

export const Fund_BalancesEntry = {
  encode(message: Fund_BalancesEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      Balances.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Fund_BalancesEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFund_BalancesEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = Balances.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Fund_BalancesEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? Balances.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: Fund_BalancesEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? Balances.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Fund_BalancesEntry>, I>>(object: I): Fund_BalancesEntry {
    const message = createBaseFund_BalancesEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? Balances.fromPartial(object.value)
      : undefined;
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
