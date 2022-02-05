/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "defundhub.defund.query";

export interface MsgCreateInterquery {
  creator: string;
  index: string;
  height: string;
  path: string;
  chainId: string;
  typeName: string;
}

export interface MsgCreateInterqueryResponse {}

export interface MsgUpdateInterquery {
  creator: string;
  index: string;
  height: string;
  path: string;
  chainId: string;
  typeName: string;
}

export interface MsgUpdateInterqueryResponse {}

export interface MsgDeleteInterquery {
  creator: string;
  index: string;
}

export interface MsgDeleteInterqueryResponse {}

const baseMsgCreateInterquery: object = {
  creator: "",
  index: "",
  height: "",
  path: "",
  chainId: "",
  typeName: "",
};

export const MsgCreateInterquery = {
  encode(
    message: MsgCreateInterquery,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.height !== "") {
      writer.uint32(26).string(message.height);
    }
    if (message.path !== "") {
      writer.uint32(34).string(message.path);
    }
    if (message.chainId !== "") {
      writer.uint32(42).string(message.chainId);
    }
    if (message.typeName !== "") {
      writer.uint32(50).string(message.typeName);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateInterquery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateInterquery } as MsgCreateInterquery;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.height = reader.string();
          break;
        case 4:
          message.path = reader.string();
          break;
        case 5:
          message.chainId = reader.string();
          break;
        case 6:
          message.typeName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateInterquery {
    const message = { ...baseMsgCreateInterquery } as MsgCreateInterquery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = String(object.height);
    } else {
      message.height = "";
    }
    if (object.path !== undefined && object.path !== null) {
      message.path = String(object.path);
    } else {
      message.path = "";
    }
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = String(object.chainId);
    } else {
      message.chainId = "";
    }
    if (object.typeName !== undefined && object.typeName !== null) {
      message.typeName = String(object.typeName);
    } else {
      message.typeName = "";
    }
    return message;
  },

  toJSON(message: MsgCreateInterquery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.height !== undefined && (obj.height = message.height);
    message.path !== undefined && (obj.path = message.path);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.typeName !== undefined && (obj.typeName = message.typeName);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateInterquery>): MsgCreateInterquery {
    const message = { ...baseMsgCreateInterquery } as MsgCreateInterquery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = "";
    }
    if (object.path !== undefined && object.path !== null) {
      message.path = object.path;
    } else {
      message.path = "";
    }
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = object.chainId;
    } else {
      message.chainId = "";
    }
    if (object.typeName !== undefined && object.typeName !== null) {
      message.typeName = object.typeName;
    } else {
      message.typeName = "";
    }
    return message;
  },
};

const baseMsgCreateInterqueryResponse: object = {};

export const MsgCreateInterqueryResponse = {
  encode(
    _: MsgCreateInterqueryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateInterqueryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateInterqueryResponse,
    } as MsgCreateInterqueryResponse;
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

  fromJSON(_: any): MsgCreateInterqueryResponse {
    const message = {
      ...baseMsgCreateInterqueryResponse,
    } as MsgCreateInterqueryResponse;
    return message;
  },

  toJSON(_: MsgCreateInterqueryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateInterqueryResponse>
  ): MsgCreateInterqueryResponse {
    const message = {
      ...baseMsgCreateInterqueryResponse,
    } as MsgCreateInterqueryResponse;
    return message;
  },
};

const baseMsgUpdateInterquery: object = {
  creator: "",
  index: "",
  height: "",
  path: "",
  chainId: "",
  typeName: "",
};

export const MsgUpdateInterquery = {
  encode(
    message: MsgUpdateInterquery,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.height !== "") {
      writer.uint32(26).string(message.height);
    }
    if (message.path !== "") {
      writer.uint32(34).string(message.path);
    }
    if (message.chainId !== "") {
      writer.uint32(42).string(message.chainId);
    }
    if (message.typeName !== "") {
      writer.uint32(50).string(message.typeName);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateInterquery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateInterquery } as MsgUpdateInterquery;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.height = reader.string();
          break;
        case 4:
          message.path = reader.string();
          break;
        case 5:
          message.chainId = reader.string();
          break;
        case 6:
          message.typeName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateInterquery {
    const message = { ...baseMsgUpdateInterquery } as MsgUpdateInterquery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = String(object.height);
    } else {
      message.height = "";
    }
    if (object.path !== undefined && object.path !== null) {
      message.path = String(object.path);
    } else {
      message.path = "";
    }
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = String(object.chainId);
    } else {
      message.chainId = "";
    }
    if (object.typeName !== undefined && object.typeName !== null) {
      message.typeName = String(object.typeName);
    } else {
      message.typeName = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateInterquery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.height !== undefined && (obj.height = message.height);
    message.path !== undefined && (obj.path = message.path);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.typeName !== undefined && (obj.typeName = message.typeName);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateInterquery>): MsgUpdateInterquery {
    const message = { ...baseMsgUpdateInterquery } as MsgUpdateInterquery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = "";
    }
    if (object.path !== undefined && object.path !== null) {
      message.path = object.path;
    } else {
      message.path = "";
    }
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = object.chainId;
    } else {
      message.chainId = "";
    }
    if (object.typeName !== undefined && object.typeName !== null) {
      message.typeName = object.typeName;
    } else {
      message.typeName = "";
    }
    return message;
  },
};

const baseMsgUpdateInterqueryResponse: object = {};

export const MsgUpdateInterqueryResponse = {
  encode(
    _: MsgUpdateInterqueryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateInterqueryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateInterqueryResponse,
    } as MsgUpdateInterqueryResponse;
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

  fromJSON(_: any): MsgUpdateInterqueryResponse {
    const message = {
      ...baseMsgUpdateInterqueryResponse,
    } as MsgUpdateInterqueryResponse;
    return message;
  },

  toJSON(_: MsgUpdateInterqueryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateInterqueryResponse>
  ): MsgUpdateInterqueryResponse {
    const message = {
      ...baseMsgUpdateInterqueryResponse,
    } as MsgUpdateInterqueryResponse;
    return message;
  },
};

const baseMsgDeleteInterquery: object = { creator: "", index: "" };

export const MsgDeleteInterquery = {
  encode(
    message: MsgDeleteInterquery,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteInterquery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteInterquery } as MsgDeleteInterquery;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteInterquery {
    const message = { ...baseMsgDeleteInterquery } as MsgDeleteInterquery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteInterquery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteInterquery>): MsgDeleteInterquery {
    const message = { ...baseMsgDeleteInterquery } as MsgDeleteInterquery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgDeleteInterqueryResponse: object = {};

export const MsgDeleteInterqueryResponse = {
  encode(
    _: MsgDeleteInterqueryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteInterqueryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteInterqueryResponse,
    } as MsgDeleteInterqueryResponse;
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

  fromJSON(_: any): MsgDeleteInterqueryResponse {
    const message = {
      ...baseMsgDeleteInterqueryResponse,
    } as MsgDeleteInterqueryResponse;
    return message;
  },

  toJSON(_: MsgDeleteInterqueryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteInterqueryResponse>
  ): MsgDeleteInterqueryResponse {
    const message = {
      ...baseMsgDeleteInterqueryResponse,
    } as MsgDeleteInterqueryResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateInterquery(
    request: MsgCreateInterquery
  ): Promise<MsgCreateInterqueryResponse>;
  UpdateInterquery(
    request: MsgUpdateInterquery
  ): Promise<MsgUpdateInterqueryResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteInterquery(
    request: MsgDeleteInterquery
  ): Promise<MsgDeleteInterqueryResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateInterquery(
    request: MsgCreateInterquery
  ): Promise<MsgCreateInterqueryResponse> {
    const data = MsgCreateInterquery.encode(request).finish();
    const promise = this.rpc.request(
      "defundhub.defund.query.Msg",
      "CreateInterquery",
      data
    );
    return promise.then((data) =>
      MsgCreateInterqueryResponse.decode(new Reader(data))
    );
  }

  UpdateInterquery(
    request: MsgUpdateInterquery
  ): Promise<MsgUpdateInterqueryResponse> {
    const data = MsgUpdateInterquery.encode(request).finish();
    const promise = this.rpc.request(
      "defundhub.defund.query.Msg",
      "UpdateInterquery",
      data
    );
    return promise.then((data) =>
      MsgUpdateInterqueryResponse.decode(new Reader(data))
    );
  }

  DeleteInterquery(
    request: MsgDeleteInterquery
  ): Promise<MsgDeleteInterqueryResponse> {
    const data = MsgDeleteInterquery.encode(request).finish();
    const promise = this.rpc.request(
      "defundhub.defund.query.Msg",
      "DeleteInterquery",
      data
    );
    return promise.then((data) =>
      MsgDeleteInterqueryResponse.decode(new Reader(data))
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
