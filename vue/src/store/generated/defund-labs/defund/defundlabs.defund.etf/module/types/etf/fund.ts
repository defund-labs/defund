/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

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
  lastRebalanceHeight: number;
  balances: { [key: string]: Balances };
}

export interface Fund_BalancesEntry {
  key: string;
  value: Balances | undefined;
}

export interface Redeem {
  id: string;
  creator: string;
  fund: Fund | undefined;
  amount: Coin | undefined;
  status: string;
  type: string;
  /** if the type is staked, we need a timestamp for when it will be unstaked */
  unstakeTimestamp: string;
}

export interface Rebalance {
  id: string;
  fund: Fund | undefined;
  /** the height the rebalance was created */
  height: number;
  broker: string;
}

const baseBaseDenom: object = { onDefund: "", onBroker: "" };

export const BaseDenom = {
  encode(message: BaseDenom, writer: Writer = Writer.create()): Writer {
    if (message.onDefund !== "") {
      writer.uint32(10).string(message.onDefund);
    }
    if (message.onBroker !== "") {
      writer.uint32(18).string(message.onBroker);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): BaseDenom {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBaseDenom } as BaseDenom;
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
    const message = { ...baseBaseDenom } as BaseDenom;
    if (object.onDefund !== undefined && object.onDefund !== null) {
      message.onDefund = String(object.onDefund);
    } else {
      message.onDefund = "";
    }
    if (object.onBroker !== undefined && object.onBroker !== null) {
      message.onBroker = String(object.onBroker);
    } else {
      message.onBroker = "";
    }
    return message;
  },

  toJSON(message: BaseDenom): unknown {
    const obj: any = {};
    message.onDefund !== undefined && (obj.onDefund = message.onDefund);
    message.onBroker !== undefined && (obj.onBroker = message.onBroker);
    return obj;
  },

  fromPartial(object: DeepPartial<BaseDenom>): BaseDenom {
    const message = { ...baseBaseDenom } as BaseDenom;
    if (object.onDefund !== undefined && object.onDefund !== null) {
      message.onDefund = object.onDefund;
    } else {
      message.onDefund = "";
    }
    if (object.onBroker !== undefined && object.onBroker !== null) {
      message.onBroker = object.onBroker;
    } else {
      message.onBroker = "";
    }
    return message;
  },
};

const baseFundPrice: object = { id: "", height: 0, symbol: "" };

export const FundPrice = {
  encode(message: FundPrice, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.height !== 0) {
      writer.uint32(16).int64(message.height);
    }
    if (message.time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.time),
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(34).fork()).ldelim();
    }
    if (message.symbol !== "") {
      writer.uint32(42).string(message.symbol);
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
          message.height = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.time = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
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
    if (object.time !== undefined && object.time !== null) {
      message.time = fromJsonTimestamp(object.time);
    } else {
      message.time = undefined;
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
    message.time !== undefined &&
      (obj.time =
        message.time !== undefined ? message.time.toISOString() : null);
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
    if (object.time !== undefined && object.time !== null) {
      message.time = object.time;
    } else {
      message.time = undefined;
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

const baseBalances: object = {};

export const Balances = {
  encode(message: Balances, writer: Writer = Writer.create()): Writer {
    for (const v of message.balances) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Balances {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBalances } as Balances;
    message.balances = [];
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
    const message = { ...baseBalances } as Balances;
    message.balances = [];
    if (object.balances !== undefined && object.balances !== null) {
      for (const e of object.balances) {
        message.balances.push(Coin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Balances): unknown {
    const obj: any = {};
    if (message.balances) {
      obj.balances = message.balances.map((e) =>
        e ? Coin.toJSON(e) : undefined
      );
    } else {
      obj.balances = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Balances>): Balances {
    const message = { ...baseBalances } as Balances;
    message.balances = [];
    if (object.balances !== undefined && object.balances !== null) {
      for (const e of object.balances) {
        message.balances.push(Coin.fromPartial(e));
      }
    }
    return message;
  },
};

const baseHolding: object = {
  token: "",
  percent: 0,
  poolId: 0,
  brokerId: "",
  type: "",
};

export const Holding = {
  encode(message: Holding, writer: Writer = Writer.create()): Writer {
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
      message.poolId = Number(object.poolId);
    } else {
      message.poolId = 0;
    }
    if (object.brokerId !== undefined && object.brokerId !== null) {
      message.brokerId = String(object.brokerId);
    } else {
      message.brokerId = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
    }
    return message;
  },

  toJSON(message: Holding): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    message.percent !== undefined && (obj.percent = message.percent);
    message.poolId !== undefined && (obj.poolId = message.poolId);
    message.brokerId !== undefined && (obj.brokerId = message.brokerId);
    message.type !== undefined && (obj.type = message.type);
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
      message.poolId = 0;
    }
    if (object.brokerId !== undefined && object.brokerId !== null) {
      message.brokerId = object.brokerId;
    } else {
      message.brokerId = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
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
  creator: "",
  lastRebalanceHeight: 0,
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
    if (message.lastRebalanceHeight !== 0) {
      writer.uint32(88).int64(message.lastRebalanceHeight);
    }
    Object.entries(message.balances).forEach(([key, value]) => {
      Fund_BalancesEntry.encode(
        { key: key as any, value },
        writer.uint32(98).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Fund {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFund } as Fund;
    message.holdings = [];
    message.balances = {};
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
          message.lastRebalanceHeight = longToNumber(reader.int64() as Long);
          break;
        case 12:
          const entry12 = Fund_BalancesEntry.decode(reader, reader.uint32());
          if (entry12.value !== undefined) {
            message.balances[entry12.key] = entry12.value;
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
    const message = { ...baseFund } as Fund;
    message.holdings = [];
    message.balances = {};
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
      message.baseDenom = BaseDenom.fromJSON(object.baseDenom);
    } else {
      message.baseDenom = undefined;
    }
    if (object.startingPrice !== undefined && object.startingPrice !== null) {
      message.startingPrice = Coin.fromJSON(object.startingPrice);
    } else {
      message.startingPrice = undefined;
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (
      object.lastRebalanceHeight !== undefined &&
      object.lastRebalanceHeight !== null
    ) {
      message.lastRebalanceHeight = Number(object.lastRebalanceHeight);
    } else {
      message.lastRebalanceHeight = 0;
    }
    if (object.balances !== undefined && object.balances !== null) {
      Object.entries(object.balances).forEach(([key, value]) => {
        message.balances[key] = Balances.fromJSON(value);
      });
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
    if (message.holdings) {
      obj.holdings = message.holdings.map((e) =>
        e ? Holding.toJSON(e) : undefined
      );
    } else {
      obj.holdings = [];
    }
    message.rebalance !== undefined && (obj.rebalance = message.rebalance);
    message.baseDenom !== undefined &&
      (obj.baseDenom = message.baseDenom
        ? BaseDenom.toJSON(message.baseDenom)
        : undefined);
    message.startingPrice !== undefined &&
      (obj.startingPrice = message.startingPrice
        ? Coin.toJSON(message.startingPrice)
        : undefined);
    message.creator !== undefined && (obj.creator = message.creator);
    message.lastRebalanceHeight !== undefined &&
      (obj.lastRebalanceHeight = message.lastRebalanceHeight);
    obj.balances = {};
    if (message.balances) {
      Object.entries(message.balances).forEach(([k, v]) => {
        obj.balances[k] = Balances.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Fund>): Fund {
    const message = { ...baseFund } as Fund;
    message.holdings = [];
    message.balances = {};
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
      message.baseDenom = BaseDenom.fromPartial(object.baseDenom);
    } else {
      message.baseDenom = undefined;
    }
    if (object.startingPrice !== undefined && object.startingPrice !== null) {
      message.startingPrice = Coin.fromPartial(object.startingPrice);
    } else {
      message.startingPrice = undefined;
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (
      object.lastRebalanceHeight !== undefined &&
      object.lastRebalanceHeight !== null
    ) {
      message.lastRebalanceHeight = object.lastRebalanceHeight;
    } else {
      message.lastRebalanceHeight = 0;
    }
    if (object.balances !== undefined && object.balances !== null) {
      Object.entries(object.balances).forEach(([key, value]) => {
        if (value !== undefined) {
          message.balances[key] = Balances.fromPartial(value);
        }
      });
    }
    return message;
  },
};

const baseFund_BalancesEntry: object = { key: "" };

export const Fund_BalancesEntry = {
  encode(
    message: Fund_BalancesEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      Balances.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Fund_BalancesEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFund_BalancesEntry } as Fund_BalancesEntry;
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
    const message = { ...baseFund_BalancesEntry } as Fund_BalancesEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Balances.fromJSON(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },

  toJSON(message: Fund_BalancesEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value ? Balances.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Fund_BalancesEntry>): Fund_BalancesEntry {
    const message = { ...baseFund_BalancesEntry } as Fund_BalancesEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Balances.fromPartial(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },
};

const baseRedeem: object = {
  id: "",
  creator: "",
  status: "",
  type: "",
  unstakeTimestamp: "",
};

export const Redeem = {
  encode(message: Redeem, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): Redeem {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRedeem } as Redeem;
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
    const message = { ...baseRedeem } as Redeem;
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
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
    }
    if (
      object.unstakeTimestamp !== undefined &&
      object.unstakeTimestamp !== null
    ) {
      message.unstakeTimestamp = String(object.unstakeTimestamp);
    } else {
      message.unstakeTimestamp = "";
    }
    return message;
  },

  toJSON(message: Redeem): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.creator !== undefined && (obj.creator = message.creator);
    message.fund !== undefined &&
      (obj.fund = message.fund ? Fund.toJSON(message.fund) : undefined);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.status !== undefined && (obj.status = message.status);
    message.type !== undefined && (obj.type = message.type);
    message.unstakeTimestamp !== undefined &&
      (obj.unstakeTimestamp = message.unstakeTimestamp);
    return obj;
  },

  fromPartial(object: DeepPartial<Redeem>): Redeem {
    const message = { ...baseRedeem } as Redeem;
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
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
    }
    if (
      object.unstakeTimestamp !== undefined &&
      object.unstakeTimestamp !== null
    ) {
      message.unstakeTimestamp = object.unstakeTimestamp;
    } else {
      message.unstakeTimestamp = "";
    }
    return message;
  },
};

const baseRebalance: object = { id: "", height: 0, broker: "" };

export const Rebalance = {
  encode(message: Rebalance, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.fund !== undefined) {
      Fund.encode(message.fund, writer.uint32(18).fork()).ldelim();
    }
    if (message.height !== 0) {
      writer.uint32(24).int64(message.height);
    }
    if (message.broker !== "") {
      writer.uint32(34).string(message.broker);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Rebalance {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRebalance } as Rebalance;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.fund = Fund.decode(reader, reader.uint32());
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
    const message = { ...baseRebalance } as Rebalance;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.fund !== undefined && object.fund !== null) {
      message.fund = Fund.fromJSON(object.fund);
    } else {
      message.fund = undefined;
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = Number(object.height);
    } else {
      message.height = 0;
    }
    if (object.broker !== undefined && object.broker !== null) {
      message.broker = String(object.broker);
    } else {
      message.broker = "";
    }
    return message;
  },

  toJSON(message: Rebalance): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.fund !== undefined &&
      (obj.fund = message.fund ? Fund.toJSON(message.fund) : undefined);
    message.height !== undefined && (obj.height = message.height);
    message.broker !== undefined && (obj.broker = message.broker);
    return obj;
  },

  fromPartial(object: DeepPartial<Rebalance>): Rebalance {
    const message = { ...baseRebalance } as Rebalance;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.fund !== undefined && object.fund !== null) {
      message.fund = Fund.fromPartial(object.fund);
    } else {
      message.fund = undefined;
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = 0;
    }
    if (object.broker !== undefined && object.broker !== null) {
      message.broker = object.broker;
    } else {
      message.broker = "";
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

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
