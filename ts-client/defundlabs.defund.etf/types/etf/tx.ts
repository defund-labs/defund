/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "defundlabs.defund.etf";

export interface MsgCreateFund {
  creator: string;
  symbol: string;
  name: string;
  description: string;
  holdings: string;
  rebalance: number;
  baseDenom: string;
  startingPrice: string;
}

export interface MsgCreateFundResponse {
}

export interface MsgCreate {
  creator: string;
  fund: string;
  tokenIn: Coin | undefined;
  channel: string;
  /**
   * Timeout height relative to the current block height.
   * The timeout is disabled when set to 0.
   */
  timeoutHeight: string;
  /**
   * Timeout timestamp in absolute nanoseconds since unix epoch.
   * The timeout is disabled when set to 0.
   */
  timeoutTimestamp: number;
}

export interface MsgCreateResponse {
}

export interface AddressMap {
  osmosisAddress: string;
}

export interface MsgRedeem {
  creator: string;
  fund: string;
  amount: Coin | undefined;
  addresses: AddressMap | undefined;
}

export interface MsgRedeemResponse {
}

function createBaseMsgCreateFund(): MsgCreateFund {
  return {
    creator: "",
    symbol: "",
    name: "",
    description: "",
    holdings: "",
    rebalance: 0,
    baseDenom: "",
    startingPrice: "",
  };
}

export const MsgCreateFund = {
  encode(message: MsgCreateFund, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.holdings !== "") {
      writer.uint32(42).string(message.holdings);
    }
    if (message.rebalance !== 0) {
      writer.uint32(48).int64(message.rebalance);
    }
    if (message.baseDenom !== "") {
      writer.uint32(58).string(message.baseDenom);
    }
    if (message.startingPrice !== "") {
      writer.uint32(66).string(message.startingPrice);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateFund {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateFund();
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
          message.holdings = reader.string();
          break;
        case 6:
          message.rebalance = longToNumber(reader.int64() as Long);
          break;
        case 7:
          message.baseDenom = reader.string();
          break;
        case 8:
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      holdings: isSet(object.holdings) ? String(object.holdings) : "",
      rebalance: isSet(object.rebalance) ? Number(object.rebalance) : 0,
      baseDenom: isSet(object.baseDenom) ? String(object.baseDenom) : "",
      startingPrice: isSet(object.startingPrice) ? String(object.startingPrice) : "",
    };
  },

  toJSON(message: MsgCreateFund): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.holdings !== undefined && (obj.holdings = message.holdings);
    message.rebalance !== undefined && (obj.rebalance = Math.round(message.rebalance));
    message.baseDenom !== undefined && (obj.baseDenom = message.baseDenom);
    message.startingPrice !== undefined && (obj.startingPrice = message.startingPrice);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateFund>, I>>(object: I): MsgCreateFund {
    const message = createBaseMsgCreateFund();
    message.creator = object.creator ?? "";
    message.symbol = object.symbol ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.holdings = object.holdings ?? "";
    message.rebalance = object.rebalance ?? 0;
    message.baseDenom = object.baseDenom ?? "";
    message.startingPrice = object.startingPrice ?? "";
    return message;
  },
};

function createBaseMsgCreateFundResponse(): MsgCreateFundResponse {
  return {};
}

export const MsgCreateFundResponse = {
  encode(_: MsgCreateFundResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateFundResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateFundResponse();
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
    return {};
  },

  toJSON(_: MsgCreateFundResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateFundResponse>, I>>(_: I): MsgCreateFundResponse {
    const message = createBaseMsgCreateFundResponse();
    return message;
  },
};

function createBaseMsgCreate(): MsgCreate {
  return { creator: "", fund: "", tokenIn: undefined, channel: "", timeoutHeight: "", timeoutTimestamp: 0 };
}

export const MsgCreate = {
  encode(message: MsgCreate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.fund !== "") {
      writer.uint32(18).string(message.fund);
    }
    if (message.tokenIn !== undefined) {
      Coin.encode(message.tokenIn, writer.uint32(26).fork()).ldelim();
    }
    if (message.channel !== "") {
      writer.uint32(34).string(message.channel);
    }
    if (message.timeoutHeight !== "") {
      writer.uint32(42).string(message.timeoutHeight);
    }
    if (message.timeoutTimestamp !== 0) {
      writer.uint32(48).uint64(message.timeoutTimestamp);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreate();
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
          message.tokenIn = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.channel = reader.string();
          break;
        case 5:
          message.timeoutHeight = reader.string();
          break;
        case 6:
          message.timeoutTimestamp = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreate {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      fund: isSet(object.fund) ? String(object.fund) : "",
      tokenIn: isSet(object.tokenIn) ? Coin.fromJSON(object.tokenIn) : undefined,
      channel: isSet(object.channel) ? String(object.channel) : "",
      timeoutHeight: isSet(object.timeoutHeight) ? String(object.timeoutHeight) : "",
      timeoutTimestamp: isSet(object.timeoutTimestamp) ? Number(object.timeoutTimestamp) : 0,
    };
  },

  toJSON(message: MsgCreate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fund !== undefined && (obj.fund = message.fund);
    message.tokenIn !== undefined && (obj.tokenIn = message.tokenIn ? Coin.toJSON(message.tokenIn) : undefined);
    message.channel !== undefined && (obj.channel = message.channel);
    message.timeoutHeight !== undefined && (obj.timeoutHeight = message.timeoutHeight);
    message.timeoutTimestamp !== undefined && (obj.timeoutTimestamp = Math.round(message.timeoutTimestamp));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreate>, I>>(object: I): MsgCreate {
    const message = createBaseMsgCreate();
    message.creator = object.creator ?? "";
    message.fund = object.fund ?? "";
    message.tokenIn = (object.tokenIn !== undefined && object.tokenIn !== null)
      ? Coin.fromPartial(object.tokenIn)
      : undefined;
    message.channel = object.channel ?? "";
    message.timeoutHeight = object.timeoutHeight ?? "";
    message.timeoutTimestamp = object.timeoutTimestamp ?? 0;
    return message;
  },
};

function createBaseMsgCreateResponse(): MsgCreateResponse {
  return {};
}

export const MsgCreateResponse = {
  encode(_: MsgCreateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateResponse();
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
    return {};
  },

  toJSON(_: MsgCreateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateResponse>, I>>(_: I): MsgCreateResponse {
    const message = createBaseMsgCreateResponse();
    return message;
  },
};

function createBaseAddressMap(): AddressMap {
  return { osmosisAddress: "" };
}

export const AddressMap = {
  encode(message: AddressMap, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.osmosisAddress !== "") {
      writer.uint32(10).string(message.osmosisAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddressMap {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddressMap();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.osmosisAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AddressMap {
    return { osmosisAddress: isSet(object.osmosisAddress) ? String(object.osmosisAddress) : "" };
  },

  toJSON(message: AddressMap): unknown {
    const obj: any = {};
    message.osmosisAddress !== undefined && (obj.osmosisAddress = message.osmosisAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AddressMap>, I>>(object: I): AddressMap {
    const message = createBaseAddressMap();
    message.osmosisAddress = object.osmosisAddress ?? "";
    return message;
  },
};

function createBaseMsgRedeem(): MsgRedeem {
  return { creator: "", fund: "", amount: undefined, addresses: undefined };
}

export const MsgRedeem = {
  encode(message: MsgRedeem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.fund !== "") {
      writer.uint32(18).string(message.fund);
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(26).fork()).ldelim();
    }
    if (message.addresses !== undefined) {
      AddressMap.encode(message.addresses, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRedeem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRedeem();
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
          message.addresses = AddressMap.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRedeem {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      fund: isSet(object.fund) ? String(object.fund) : "",
      amount: isSet(object.amount) ? Coin.fromJSON(object.amount) : undefined,
      addresses: isSet(object.addresses) ? AddressMap.fromJSON(object.addresses) : undefined,
    };
  },

  toJSON(message: MsgRedeem): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fund !== undefined && (obj.fund = message.fund);
    message.amount !== undefined && (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.addresses !== undefined
      && (obj.addresses = message.addresses ? AddressMap.toJSON(message.addresses) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRedeem>, I>>(object: I): MsgRedeem {
    const message = createBaseMsgRedeem();
    message.creator = object.creator ?? "";
    message.fund = object.fund ?? "";
    message.amount = (object.amount !== undefined && object.amount !== null)
      ? Coin.fromPartial(object.amount)
      : undefined;
    message.addresses = (object.addresses !== undefined && object.addresses !== null)
      ? AddressMap.fromPartial(object.addresses)
      : undefined;
    return message;
  },
};

function createBaseMsgRedeemResponse(): MsgRedeemResponse {
  return {};
}

export const MsgRedeemResponse = {
  encode(_: MsgRedeemResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRedeemResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRedeemResponse();
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
    return {};
  },

  toJSON(_: MsgRedeemResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRedeemResponse>, I>>(_: I): MsgRedeemResponse {
    const message = createBaseMsgRedeemResponse();
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
    this.CreateFund = this.CreateFund.bind(this);
    this.Create = this.Create.bind(this);
    this.Redeem = this.Redeem.bind(this);
  }
  CreateFund(request: MsgCreateFund): Promise<MsgCreateFundResponse> {
    const data = MsgCreateFund.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.etf.Msg", "CreateFund", data);
    return promise.then((data) => MsgCreateFundResponse.decode(new _m0.Reader(data)));
  }

  Create(request: MsgCreate): Promise<MsgCreateResponse> {
    const data = MsgCreate.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.etf.Msg", "Create", data);
    return promise.then((data) => MsgCreateResponse.decode(new _m0.Reader(data)));
  }

  Redeem(request: MsgRedeem): Promise<MsgRedeemResponse> {
    const data = MsgRedeem.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.etf.Msg", "Redeem", data);
    return promise.then((data) => MsgRedeemResponse.decode(new _m0.Reader(data)));
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
