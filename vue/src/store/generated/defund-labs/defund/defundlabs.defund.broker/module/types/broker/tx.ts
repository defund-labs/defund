/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Any } from "../google/protobuf/any";

export const protobufPackage = "defundlabs.defund.broker";

/** MsgRegisterAccount defines the payload for Msg/RegisterAccount */
export interface MsgRegisterAccount {
  owner: string;
  connection_id: string;
}

/** MsgRegisterAccountResponse defines the response for Msg/RegisterAccount */
export interface MsgRegisterAccountResponse {}

/** MsgSubmitTx defines the payload for Msg/SubmitTx */
export interface MsgSubmitTx {
  owner: string;
  connection_id: string;
  msg: Any | undefined;
}

/** MsgSubmitTxResponse defines the response for Msg/SubmitTx */
export interface MsgSubmitTxResponse {}

const baseMsgRegisterAccount: object = { owner: "", connection_id: "" };

export const MsgRegisterAccount = {
  encode(
    message: MsgRegisterAccount,
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

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRegisterAccount } as MsgRegisterAccount;
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

  fromJSON(object: any): MsgRegisterAccount {
    const message = { ...baseMsgRegisterAccount } as MsgRegisterAccount;
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

  toJSON(message: MsgRegisterAccount): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.connection_id !== undefined &&
      (obj.connection_id = message.connection_id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRegisterAccount>): MsgRegisterAccount {
    const message = { ...baseMsgRegisterAccount } as MsgRegisterAccount;
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

const baseMsgRegisterAccountResponse: object = {};

export const MsgRegisterAccountResponse = {
  encode(
    _: MsgRegisterAccountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRegisterAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRegisterAccountResponse,
    } as MsgRegisterAccountResponse;
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

  fromJSON(_: any): MsgRegisterAccountResponse {
    const message = {
      ...baseMsgRegisterAccountResponse,
    } as MsgRegisterAccountResponse;
    return message;
  },

  toJSON(_: MsgRegisterAccountResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRegisterAccountResponse>
  ): MsgRegisterAccountResponse {
    const message = {
      ...baseMsgRegisterAccountResponse,
    } as MsgRegisterAccountResponse;
    return message;
  },
};

const baseMsgSubmitTx: object = { owner: "", connection_id: "" };

export const MsgSubmitTx = {
  encode(message: MsgSubmitTx, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): MsgSubmitTx {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSubmitTx } as MsgSubmitTx;
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

  fromJSON(object: any): MsgSubmitTx {
    const message = { ...baseMsgSubmitTx } as MsgSubmitTx;
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

  toJSON(message: MsgSubmitTx): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.connection_id !== undefined &&
      (obj.connection_id = message.connection_id);
    message.msg !== undefined &&
      (obj.msg = message.msg ? Any.toJSON(message.msg) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSubmitTx>): MsgSubmitTx {
    const message = { ...baseMsgSubmitTx } as MsgSubmitTx;
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

const baseMsgSubmitTxResponse: object = {};

export const MsgSubmitTxResponse = {
  encode(_: MsgSubmitTxResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSubmitTxResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSubmitTxResponse } as MsgSubmitTxResponse;
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

  fromJSON(_: any): MsgSubmitTxResponse {
    const message = { ...baseMsgSubmitTxResponse } as MsgSubmitTxResponse;
    return message;
  },

  toJSON(_: MsgSubmitTxResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSubmitTxResponse>): MsgSubmitTxResponse {
    const message = { ...baseMsgSubmitTxResponse } as MsgSubmitTxResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** Register defines a rpc handler for MsgRegisterAccount */
  RegisterAccount(
    request: MsgRegisterAccount
  ): Promise<MsgRegisterAccountResponse>;
  /** SubmitTx defines a rpc handler for MsgSubmitTx */
  SubmitTx(request: MsgSubmitTx): Promise<MsgSubmitTxResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  RegisterAccount(
    request: MsgRegisterAccount
  ): Promise<MsgRegisterAccountResponse> {
    const data = MsgRegisterAccount.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.broker.Msg",
      "RegisterAccount",
      data
    );
    return promise.then((data) =>
      MsgRegisterAccountResponse.decode(new Reader(data))
    );
  }

  SubmitTx(request: MsgSubmitTx): Promise<MsgSubmitTxResponse> {
    const data = MsgSubmitTx.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.broker.Msg",
      "SubmitTx",
      data
    );
    return promise.then((data) => MsgSubmitTxResponse.decode(new Reader(data)));
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
