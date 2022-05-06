/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Any } from "../google/protobuf/any";

export const protobufPackage = "defundlabs.defund.broker";

/** MsgRegisterBrokerAccount defines the payload for Msg/RegisterBrokerAccount */
export interface MsgRegisterBrokerAccount {
  owner: string;
  connection_id: string;
}

/** MsgRegisterBrokerAccountResponse defines the response for Msg/RegisterBrokerAccount */
export interface MsgRegisterBrokerAccountResponse {}

/** MsgCosmosSwap defines the payload for Msg/CosmosSwap */
export interface MsgCosmosSwap {
  owner: string;
  connection_id: string;
  msg: Any | undefined;
}

/** MsgCosmosSwapResponse defines the response for Msg/CosmosSwap */
export interface MsgCosmosSwapResponse {}

const baseMsgRegisterBrokerAccount: object = { owner: "", connection_id: "" };

export const MsgRegisterBrokerAccount = {
  encode(
    message: MsgRegisterBrokerAccount,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.connection_id !== "") {
      writer.uint32(18).string(message.connection_id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRegisterBrokerAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRegisterBrokerAccount,
    } as MsgRegisterBrokerAccount;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.connection_id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterBrokerAccount {
    const message = {
      ...baseMsgRegisterBrokerAccount,
    } as MsgRegisterBrokerAccount;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.connection_id !== undefined && object.connection_id !== null) {
      message.connection_id = String(object.connection_id);
    } else {
      message.connection_id = "";
    }
    return message;
  },

  toJSON(message: MsgRegisterBrokerAccount): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.connection_id !== undefined &&
      (obj.connection_id = message.connection_id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRegisterBrokerAccount>
  ): MsgRegisterBrokerAccount {
    const message = {
      ...baseMsgRegisterBrokerAccount,
    } as MsgRegisterBrokerAccount;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.connection_id !== undefined && object.connection_id !== null) {
      message.connection_id = object.connection_id;
    } else {
      message.connection_id = "";
    }
    return message;
  },
};

const baseMsgRegisterBrokerAccountResponse: object = {};

export const MsgRegisterBrokerAccountResponse = {
  encode(
    _: MsgRegisterBrokerAccountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRegisterBrokerAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRegisterBrokerAccountResponse,
    } as MsgRegisterBrokerAccountResponse;
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

  fromJSON(_: any): MsgRegisterBrokerAccountResponse {
    const message = {
      ...baseMsgRegisterBrokerAccountResponse,
    } as MsgRegisterBrokerAccountResponse;
    return message;
  },

  toJSON(_: MsgRegisterBrokerAccountResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRegisterBrokerAccountResponse>
  ): MsgRegisterBrokerAccountResponse {
    const message = {
      ...baseMsgRegisterBrokerAccountResponse,
    } as MsgRegisterBrokerAccountResponse;
    return message;
  },
};

const baseMsgCosmosSwap: object = { owner: "", connection_id: "" };

export const MsgCosmosSwap = {
  encode(message: MsgCosmosSwap, writer: Writer = Writer.create()): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.connection_id !== "") {
      writer.uint32(18).string(message.connection_id);
    }
    if (message.msg !== undefined) {
      Any.encode(message.msg, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCosmosSwap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCosmosSwap } as MsgCosmosSwap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.connection_id = reader.string();
          break;
        case 3:
          message.msg = Any.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCosmosSwap {
    const message = { ...baseMsgCosmosSwap } as MsgCosmosSwap;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.connection_id !== undefined && object.connection_id !== null) {
      message.connection_id = String(object.connection_id);
    } else {
      message.connection_id = "";
    }
    if (object.msg !== undefined && object.msg !== null) {
      message.msg = Any.fromJSON(object.msg);
    } else {
      message.msg = undefined;
    }
    return message;
  },

  toJSON(message: MsgCosmosSwap): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.connection_id !== undefined &&
      (obj.connection_id = message.connection_id);
    message.msg !== undefined &&
      (obj.msg = message.msg ? Any.toJSON(message.msg) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCosmosSwap>): MsgCosmosSwap {
    const message = { ...baseMsgCosmosSwap } as MsgCosmosSwap;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.connection_id !== undefined && object.connection_id !== null) {
      message.connection_id = object.connection_id;
    } else {
      message.connection_id = "";
    }
    if (object.msg !== undefined && object.msg !== null) {
      message.msg = Any.fromPartial(object.msg);
    } else {
      message.msg = undefined;
    }
    return message;
  },
};

const baseMsgCosmosSwapResponse: object = {};

export const MsgCosmosSwapResponse = {
  encode(_: MsgCosmosSwapResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCosmosSwapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCosmosSwapResponse } as MsgCosmosSwapResponse;
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

  fromJSON(_: any): MsgCosmosSwapResponse {
    const message = { ...baseMsgCosmosSwapResponse } as MsgCosmosSwapResponse;
    return message;
  },

  toJSON(_: MsgCosmosSwapResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCosmosSwapResponse>): MsgCosmosSwapResponse {
    const message = { ...baseMsgCosmosSwapResponse } as MsgCosmosSwapResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** Register defines a rpc handler for MsgRegisterBrokerAccount */
  RegisterBrokerAccount(
    request: MsgRegisterBrokerAccount
  ): Promise<MsgRegisterBrokerAccountResponse>;
  /** CosmosSwap defines a rpc handler for MsgCosmosSwap */
  CosmosSwap(request: MsgCosmosSwap): Promise<MsgCosmosSwapResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  RegisterBrokerAccount(
    request: MsgRegisterBrokerAccount
  ): Promise<MsgRegisterBrokerAccountResponse> {
    const data = MsgRegisterBrokerAccount.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.broker.Msg",
      "RegisterBrokerAccount",
      data
    );
    return promise.then((data) =>
      MsgRegisterBrokerAccountResponse.decode(new Reader(data))
    );
  }

  CosmosSwap(request: MsgCosmosSwap): Promise<MsgCosmosSwapResponse> {
    const data = MsgCosmosSwap.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.broker.Msg",
      "CosmosSwap",
      data
    );
    return promise.then((data) =>
      MsgCosmosSwapResponse.decode(new Reader(data))
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
