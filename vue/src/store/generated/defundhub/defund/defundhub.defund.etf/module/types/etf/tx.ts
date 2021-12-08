/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "defundhub.defund.etf";

export interface MsgCreateFund {
  creator: string;
  id: string;
  address: string;
  symbol: string;
  name: string;
  description: string;
  shares: Coin | undefined;
}

export interface MsgCreateFundResponse {}

export interface MsgUpdateFund {
  creator: string;
  id: string;
  address: string;
  symbol: string;
  name: string;
  description: string;
  shares: Coin | undefined;
}

export interface MsgUpdateFundResponse {}

const baseMsgCreateFund: object = {
  creator: "",
  id: "",
  address: "",
  symbol: "",
  name: "",
  description: "",
};

export const MsgCreateFund = {
  encode(message: MsgCreateFund, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.symbol !== "") {
      writer.uint32(34).string(message.symbol);
    }
    if (message.name !== "") {
      writer.uint32(42).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(50).string(message.description);
    }
    if (message.shares !== undefined) {
      Coin.encode(message.shares, writer.uint32(58).fork()).ldelim();
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
          message.id = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
          message.symbol = reader.string();
          break;
        case 5:
          message.name = reader.string();
          break;
        case 6:
          message.description = reader.string();
          break;
        case 7:
          message.shares = Coin.decode(reader, reader.uint32());
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
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
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
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = Coin.fromJSON(object.shares);
    } else {
      message.shares = undefined;
    }
    return message;
  },

  toJSON(message: MsgCreateFund): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.address !== undefined && (obj.address = message.address);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    message.shares !== undefined &&
      (obj.shares = message.shares ? Coin.toJSON(message.shares) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateFund>): MsgCreateFund {
    const message = { ...baseMsgCreateFund } as MsgCreateFund;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
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
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = Coin.fromPartial(object.shares);
    } else {
      message.shares = undefined;
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

const baseMsgUpdateFund: object = {
  creator: "",
  id: "",
  address: "",
  symbol: "",
  name: "",
  description: "",
};

export const MsgUpdateFund = {
  encode(message: MsgUpdateFund, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.symbol !== "") {
      writer.uint32(34).string(message.symbol);
    }
    if (message.name !== "") {
      writer.uint32(42).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(50).string(message.description);
    }
    if (message.shares !== undefined) {
      Coin.encode(message.shares, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateFund {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateFund } as MsgUpdateFund;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
          message.symbol = reader.string();
          break;
        case 5:
          message.name = reader.string();
          break;
        case 6:
          message.description = reader.string();
          break;
        case 7:
          message.shares = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateFund {
    const message = { ...baseMsgUpdateFund } as MsgUpdateFund;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
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
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = Coin.fromJSON(object.shares);
    } else {
      message.shares = undefined;
    }
    return message;
  },

  toJSON(message: MsgUpdateFund): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.address !== undefined && (obj.address = message.address);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    message.shares !== undefined &&
      (obj.shares = message.shares ? Coin.toJSON(message.shares) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateFund>): MsgUpdateFund {
    const message = { ...baseMsgUpdateFund } as MsgUpdateFund;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
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
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = Coin.fromPartial(object.shares);
    } else {
      message.shares = undefined;
    }
    return message;
  },
};

const baseMsgUpdateFundResponse: object = {};

export const MsgUpdateFundResponse = {
  encode(_: MsgUpdateFundResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateFundResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateFundResponse } as MsgUpdateFundResponse;
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

  fromJSON(_: any): MsgUpdateFundResponse {
    const message = { ...baseMsgUpdateFundResponse } as MsgUpdateFundResponse;
    return message;
  },

  toJSON(_: MsgUpdateFundResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgUpdateFundResponse>): MsgUpdateFundResponse {
    const message = { ...baseMsgUpdateFundResponse } as MsgUpdateFundResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateFund(request: MsgCreateFund): Promise<MsgCreateFundResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdateFund(request: MsgUpdateFund): Promise<MsgUpdateFundResponse>;
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

  UpdateFund(request: MsgUpdateFund): Promise<MsgUpdateFundResponse> {
    const data = MsgUpdateFund.encode(request).finish();
    const promise = this.rpc.request(
      "defundhub.defund.etf.Msg",
      "UpdateFund",
      data
    );
    return promise.then((data) =>
      MsgUpdateFundResponse.decode(new Reader(data))
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
