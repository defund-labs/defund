/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "defundlabs.defund.etf";

export interface MsgCreateFund {
  creator: string;
  symbol: string;
  name: string;
  description: string;
  broker: string;
  holdings: string;
  rebalance: number;
  baseDenom: string;
  startingPrice: string;
}

export interface MsgCreateFundResponse {}

export interface MsgCreate {
  creator: string;
  fund: string;
  tokens: Coin[];
  channel: string;
  /**
   * Timeout height relative to the current block height.
   * The timeout is disabled when set to 0.
   */
  timeout_height: string;
  /**
   * Timeout timestamp in absolute nanoseconds since unix epoch.
   * The timeout is disabled when set to 0.
   */
  timeout_timestamp: number;
}

export interface MsgCreateResponse {}

export interface MsgRedeem {
  creator: string;
  fund: string;
  amount: Coin | undefined;
  channel: string;
}

export interface MsgRedeemResponse {}

const baseMsgCreateFund: object = {
  creator: "",
  symbol: "",
  name: "",
  description: "",
  broker: "",
  holdings: "",
  rebalance: 0,
  baseDenom: "",
  startingPrice: "",
};

export const MsgCreateFund = {
  encode(message: MsgCreateFund, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.symbol !== "") {
      writer.uint32(18).string(message.symbol);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    if (message.broker !== "") {
      writer.uint32(42).string(message.broker);
    }
    if (message.holdings !== "") {
      writer.uint32(50).string(message.holdings);
    }
    if (message.rebalance !== 0) {
      writer.uint32(56).int64(message.rebalance);
    }
    if (message.baseDenom !== "") {
      writer.uint32(66).string(message.baseDenom);
    }
    if (message.startingPrice !== "") {
      writer.uint32(74).string(message.startingPrice);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateFund {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateFund } as MsgCreateFund;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.symbol = reader.string();
          break;
        case 3:
          message.name = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.broker = reader.string();
          break;
        case 6:
          message.holdings = reader.string();
          break;
        case 7:
          message.rebalance = longToNumber(reader.int64() as Long);
          break;
        case 8:
          message.baseDenom = reader.string();
          break;
        case 9:
          message.startingPrice = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateFund {
    const message = { ...baseMsgCreateFund } as MsgCreateFund;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = String(object.symbol);
    } else {
      message.symbol = "";
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
    if (object.broker !== undefined && object.broker !== null) {
      message.broker = String(object.broker);
    } else {
      message.broker = "";
    }
    if (object.holdings !== undefined && object.holdings !== null) {
      message.holdings = String(object.holdings);
    } else {
      message.holdings = "";
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
    if (object.startingPrice !== undefined && object.startingPrice !== null) {
      message.startingPrice = String(object.startingPrice);
    } else {
      message.startingPrice = "";
    }
    return message;
  },

  toJSON(message: MsgCreateFund): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    message.broker !== undefined && (obj.broker = message.broker);
    message.holdings !== undefined && (obj.holdings = message.holdings);
    message.rebalance !== undefined && (obj.rebalance = message.rebalance);
    message.baseDenom !== undefined && (obj.baseDenom = message.baseDenom);
    message.startingPrice !== undefined &&
      (obj.startingPrice = message.startingPrice);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateFund>): MsgCreateFund {
    const message = { ...baseMsgCreateFund } as MsgCreateFund;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = object.symbol;
    } else {
      message.symbol = "";
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
    if (object.broker !== undefined && object.broker !== null) {
      message.broker = object.broker;
    } else {
      message.broker = "";
    }
    if (object.holdings !== undefined && object.holdings !== null) {
      message.holdings = object.holdings;
    } else {
      message.holdings = "";
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
    if (object.startingPrice !== undefined && object.startingPrice !== null) {
      message.startingPrice = object.startingPrice;
    } else {
      message.startingPrice = "";
    }
    return message;
  },
};

const baseMsgCreateFundResponse: object = {};

export const MsgCreateFundResponse = {
  encode(_: MsgCreateFundResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateFundResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateFundResponse } as MsgCreateFundResponse;
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

  fromJSON(_: any): MsgCreateFundResponse {
    const message = { ...baseMsgCreateFundResponse } as MsgCreateFundResponse;
    return message;
  },

  toJSON(_: MsgCreateFundResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreateFundResponse>): MsgCreateFundResponse {
    const message = { ...baseMsgCreateFundResponse } as MsgCreateFundResponse;
    return message;
  },
};

const baseMsgCreate: object = {
  creator: "",
  fund: "",
  channel: "",
  timeout_height: "",
  timeout_timestamp: 0,
};

export const MsgCreate = {
  encode(message: MsgCreate, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.fund !== "") {
      writer.uint32(18).string(message.fund);
    }
    for (const v of message.tokens) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.channel !== "") {
      writer.uint32(34).string(message.channel);
    }
    if (message.timeout_height !== "") {
      writer.uint32(42).string(message.timeout_height);
    }
    if (message.timeout_timestamp !== 0) {
      writer.uint32(48).uint64(message.timeout_timestamp);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreate {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreate } as MsgCreate;
    message.tokens = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.fund = reader.string();
          break;
        case 3:
          message.tokens.push(Coin.decode(reader, reader.uint32()));
          break;
        case 4:
          message.channel = reader.string();
          break;
        case 5:
          message.timeout_height = reader.string();
          break;
        case 6:
          message.timeout_timestamp = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreate {
    const message = { ...baseMsgCreate } as MsgCreate;
    message.tokens = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.fund !== undefined && object.fund !== null) {
      message.fund = String(object.fund);
    } else {
      message.fund = "";
    }
    if (object.tokens !== undefined && object.tokens !== null) {
      for (const e of object.tokens) {
        message.tokens.push(Coin.fromJSON(e));
      }
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = String(object.channel);
    } else {
      message.channel = "";
    }
    if (object.timeout_height !== undefined && object.timeout_height !== null) {
      message.timeout_height = String(object.timeout_height);
    } else {
      message.timeout_height = "";
    }
    if (
      object.timeout_timestamp !== undefined &&
      object.timeout_timestamp !== null
    ) {
      message.timeout_timestamp = Number(object.timeout_timestamp);
    } else {
      message.timeout_timestamp = 0;
    }
    return message;
  },

  toJSON(message: MsgCreate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fund !== undefined && (obj.fund = message.fund);
    if (message.tokens) {
      obj.tokens = message.tokens.map((e) => (e ? Coin.toJSON(e) : undefined));
    } else {
      obj.tokens = [];
    }
    message.channel !== undefined && (obj.channel = message.channel);
    message.timeout_height !== undefined &&
      (obj.timeout_height = message.timeout_height);
    message.timeout_timestamp !== undefined &&
      (obj.timeout_timestamp = message.timeout_timestamp);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreate>): MsgCreate {
    const message = { ...baseMsgCreate } as MsgCreate;
    message.tokens = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.fund !== undefined && object.fund !== null) {
      message.fund = object.fund;
    } else {
      message.fund = "";
    }
    if (object.tokens !== undefined && object.tokens !== null) {
      for (const e of object.tokens) {
        message.tokens.push(Coin.fromPartial(e));
      }
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = object.channel;
    } else {
      message.channel = "";
    }
    if (object.timeout_height !== undefined && object.timeout_height !== null) {
      message.timeout_height = object.timeout_height;
    } else {
      message.timeout_height = "";
    }
    if (
      object.timeout_timestamp !== undefined &&
      object.timeout_timestamp !== null
    ) {
      message.timeout_timestamp = object.timeout_timestamp;
    } else {
      message.timeout_timestamp = 0;
    }
    return message;
  },
};

const baseMsgCreateResponse: object = {};

export const MsgCreateResponse = {
  encode(_: MsgCreateResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateResponse } as MsgCreateResponse;
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

  fromJSON(_: any): MsgCreateResponse {
    const message = { ...baseMsgCreateResponse } as MsgCreateResponse;
    return message;
  },

  toJSON(_: MsgCreateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreateResponse>): MsgCreateResponse {
    const message = { ...baseMsgCreateResponse } as MsgCreateResponse;
    return message;
  },
};

const baseMsgRedeem: object = { creator: "", fund: "", channel: "" };

export const MsgRedeem = {
  encode(message: MsgRedeem, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.fund !== "") {
      writer.uint32(18).string(message.fund);
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(26).fork()).ldelim();
    }
    if (message.channel !== "") {
      writer.uint32(34).string(message.channel);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRedeem {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRedeem } as MsgRedeem;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.fund = reader.string();
          break;
        case 3:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.channel = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRedeem {
    const message = { ...baseMsgRedeem } as MsgRedeem;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.fund !== undefined && object.fund !== null) {
      message.fund = String(object.fund);
    } else {
      message.fund = "";
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
    return message;
  },

  toJSON(message: MsgRedeem): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fund !== undefined && (obj.fund = message.fund);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.channel !== undefined && (obj.channel = message.channel);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRedeem>): MsgRedeem {
    const message = { ...baseMsgRedeem } as MsgRedeem;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.fund !== undefined && object.fund !== null) {
      message.fund = object.fund;
    } else {
      message.fund = "";
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
    return message;
  },
};

const baseMsgRedeemResponse: object = {};

export const MsgRedeemResponse = {
  encode(_: MsgRedeemResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRedeemResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRedeemResponse } as MsgRedeemResponse;
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

  fromJSON(_: any): MsgRedeemResponse {
    const message = { ...baseMsgRedeemResponse } as MsgRedeemResponse;
    return message;
  },

  toJSON(_: MsgRedeemResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgRedeemResponse>): MsgRedeemResponse {
    const message = { ...baseMsgRedeemResponse } as MsgRedeemResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateFund(request: MsgCreateFund): Promise<MsgCreateFundResponse>;
  Create(request: MsgCreate): Promise<MsgCreateResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Redeem(request: MsgRedeem): Promise<MsgRedeemResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateFund(request: MsgCreateFund): Promise<MsgCreateFundResponse> {
    const data = MsgCreateFund.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.etf.Msg",
      "CreateFund",
      data
    );
    return promise.then((data) =>
      MsgCreateFundResponse.decode(new Reader(data))
    );
  }

  Create(request: MsgCreate): Promise<MsgCreateResponse> {
    const data = MsgCreate.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.etf.Msg",
      "Create",
      data
    );
    return promise.then((data) => MsgCreateResponse.decode(new Reader(data)));
  }

  Redeem(request: MsgRedeem): Promise<MsgRedeemResponse> {
    const data = MsgRedeem.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.etf.Msg",
      "Redeem",
      data
    );
    return promise.then((data) => MsgRedeemResponse.decode(new Reader(data)));
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
