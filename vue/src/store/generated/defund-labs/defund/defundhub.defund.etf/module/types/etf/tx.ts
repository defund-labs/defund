/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "defundhub.defund.etf";

export interface MsgCreateFund {
  creator: string;
  symbol: string;
  name: string;
  description: string;
  broker: string;
  holdings: string;
  rebalance: number;
  baseDenom: string;
  connectionId: string;
}

export interface MsgCreateFundResponse {}

export interface MsgInvest {
  creator: string;
  fund: string;
  amount: Coin | undefined;
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

export interface MsgInvestResponse {}

export interface MsgUninvest {
  creator: string;
  fund: string;
  amount: Coin | undefined;
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

export interface MsgUninvestResponse {}

const baseMsgCreateFund: object = {
  creator: "",
  symbol: "",
  name: "",
  description: "",
  broker: "",
  holdings: "",
  rebalance: 0,
  baseDenom: "",
  connectionId: "",
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
    if (message.connectionId !== "") {
      writer.uint32(74).string(message.connectionId);
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
          message.connectionId = reader.string();
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
    if (object.connectionId !== undefined && object.connectionId !== null) {
      message.connectionId = String(object.connectionId);
    } else {
      message.connectionId = "";
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
    message.connectionId !== undefined &&
      (obj.connectionId = message.connectionId);
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
    if (object.connectionId !== undefined && object.connectionId !== null) {
      message.connectionId = object.connectionId;
    } else {
      message.connectionId = "";
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

const baseMsgInvest: object = {
  creator: "",
  fund: "",
  channel: "",
  timeout_height: "",
  timeout_timestamp: 0,
};

export const MsgInvest = {
  encode(message: MsgInvest, writer: Writer = Writer.create()): Writer {
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
    if (message.timeout_height !== "") {
      writer.uint32(42).string(message.timeout_height);
    }
    if (message.timeout_timestamp !== 0) {
      writer.uint32(48).uint64(message.timeout_timestamp);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInvest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInvest } as MsgInvest;
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

  fromJSON(object: any): MsgInvest {
    const message = { ...baseMsgInvest } as MsgInvest;
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

  toJSON(message: MsgInvest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fund !== undefined && (obj.fund = message.fund);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.channel !== undefined && (obj.channel = message.channel);
    message.timeout_height !== undefined &&
      (obj.timeout_height = message.timeout_height);
    message.timeout_timestamp !== undefined &&
      (obj.timeout_timestamp = message.timeout_timestamp);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgInvest>): MsgInvest {
    const message = { ...baseMsgInvest } as MsgInvest;
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

const baseMsgInvestResponse: object = {};

export const MsgInvestResponse = {
  encode(_: MsgInvestResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInvestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInvestResponse } as MsgInvestResponse;
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

  fromJSON(_: any): MsgInvestResponse {
    const message = { ...baseMsgInvestResponse } as MsgInvestResponse;
    return message;
  },

  toJSON(_: MsgInvestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgInvestResponse>): MsgInvestResponse {
    const message = { ...baseMsgInvestResponse } as MsgInvestResponse;
    return message;
  },
};

const baseMsgUninvest: object = {
  creator: "",
  fund: "",
  channel: "",
  timeout_height: "",
  timeout_timestamp: 0,
};

export const MsgUninvest = {
  encode(message: MsgUninvest, writer: Writer = Writer.create()): Writer {
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
    if (message.timeout_height !== "") {
      writer.uint32(42).string(message.timeout_height);
    }
    if (message.timeout_timestamp !== 0) {
      writer.uint32(48).uint64(message.timeout_timestamp);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUninvest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUninvest } as MsgUninvest;
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

  fromJSON(object: any): MsgUninvest {
    const message = { ...baseMsgUninvest } as MsgUninvest;
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

  toJSON(message: MsgUninvest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fund !== undefined && (obj.fund = message.fund);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.channel !== undefined && (obj.channel = message.channel);
    message.timeout_height !== undefined &&
      (obj.timeout_height = message.timeout_height);
    message.timeout_timestamp !== undefined &&
      (obj.timeout_timestamp = message.timeout_timestamp);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUninvest>): MsgUninvest {
    const message = { ...baseMsgUninvest } as MsgUninvest;
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

const baseMsgUninvestResponse: object = {};

export const MsgUninvestResponse = {
  encode(_: MsgUninvestResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUninvestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUninvestResponse } as MsgUninvestResponse;
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

  fromJSON(_: any): MsgUninvestResponse {
    const message = { ...baseMsgUninvestResponse } as MsgUninvestResponse;
    return message;
  },

  toJSON(_: MsgUninvestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgUninvestResponse>): MsgUninvestResponse {
    const message = { ...baseMsgUninvestResponse } as MsgUninvestResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateFund(request: MsgCreateFund): Promise<MsgCreateFundResponse>;
  Invest(request: MsgInvest): Promise<MsgInvestResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Uninvest(request: MsgUninvest): Promise<MsgUninvestResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateFund(request: MsgCreateFund): Promise<MsgCreateFundResponse> {
    const data = MsgCreateFund.encode(request).finish();
    const promise = this.rpc.request(
      "defundhub.defund.etf.Msg",
      "CreateFund",
      data
    );
    return promise.then((data) =>
      MsgCreateFundResponse.decode(new Reader(data))
    );
  }

  Invest(request: MsgInvest): Promise<MsgInvestResponse> {
    const data = MsgInvest.encode(request).finish();
    const promise = this.rpc.request(
      "defundhub.defund.etf.Msg",
      "Invest",
      data
    );
    return promise.then((data) => MsgInvestResponse.decode(new Reader(data)));
  }

  Uninvest(request: MsgUninvest): Promise<MsgUninvestResponse> {
    const data = MsgUninvest.encode(request).finish();
    const promise = this.rpc.request(
      "defundhub.defund.etf.Msg",
      "Uninvest",
      data
    );
    return promise.then((data) => MsgUninvestResponse.decode(new Reader(data)));
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
