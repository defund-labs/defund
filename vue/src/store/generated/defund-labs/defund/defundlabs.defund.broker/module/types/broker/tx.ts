/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "defundlabs.defund.broker";

export interface MsgAddLiquiditySource {
  creator: string;
  brokerId: string;
  poolId: string;
}

export interface MsgAddLiquiditySourceResponse {}

export interface MsgAddConnectionBroker {
  creator: string;
  brokerId: string;
  connectionId: string;
}

export interface MsgAddConnectionBrokerResponse {}

const baseMsgAddLiquiditySource: object = {
  creator: "",
  brokerId: "",
  poolId: "",
};

export const MsgAddLiquiditySource = {
  encode(
    message: MsgAddLiquiditySource,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.brokerId !== "") {
      writer.uint32(18).string(message.brokerId);
    }
    if (message.poolId !== "") {
      writer.uint32(26).string(message.poolId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddLiquiditySource {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddLiquiditySource } as MsgAddLiquiditySource;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.brokerId = reader.string();
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

  fromJSON(object: any): MsgAddLiquiditySource {
    const message = { ...baseMsgAddLiquiditySource } as MsgAddLiquiditySource;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.brokerId !== undefined && object.brokerId !== null) {
      message.brokerId = String(object.brokerId);
    } else {
      message.brokerId = "";
    }
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = String(object.poolId);
    } else {
      message.poolId = "";
    }
    return message;
  },

  toJSON(message: MsgAddLiquiditySource): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.brokerId !== undefined && (obj.brokerId = message.brokerId);
    message.poolId !== undefined && (obj.poolId = message.poolId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgAddLiquiditySource>
  ): MsgAddLiquiditySource {
    const message = { ...baseMsgAddLiquiditySource } as MsgAddLiquiditySource;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.brokerId !== undefined && object.brokerId !== null) {
      message.brokerId = object.brokerId;
    } else {
      message.brokerId = "";
    }
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = object.poolId;
    } else {
      message.poolId = "";
    }
    return message;
  },
};

const baseMsgAddLiquiditySourceResponse: object = {};

export const MsgAddLiquiditySourceResponse = {
  encode(
    _: MsgAddLiquiditySourceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddLiquiditySourceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddLiquiditySourceResponse,
    } as MsgAddLiquiditySourceResponse;
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

  fromJSON(_: any): MsgAddLiquiditySourceResponse {
    const message = {
      ...baseMsgAddLiquiditySourceResponse,
    } as MsgAddLiquiditySourceResponse;
    return message;
  },

  toJSON(_: MsgAddLiquiditySourceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddLiquiditySourceResponse>
  ): MsgAddLiquiditySourceResponse {
    const message = {
      ...baseMsgAddLiquiditySourceResponse,
    } as MsgAddLiquiditySourceResponse;
    return message;
  },
};

const baseMsgAddConnectionBroker: object = {
  creator: "",
  brokerId: "",
  connectionId: "",
};

export const MsgAddConnectionBroker = {
  encode(
    message: MsgAddConnectionBroker,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.brokerId !== "") {
      writer.uint32(18).string(message.brokerId);
    }
    if (message.connectionId !== "") {
      writer.uint32(26).string(message.connectionId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddConnectionBroker {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddConnectionBroker } as MsgAddConnectionBroker;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.brokerId = reader.string();
          break;
        case 3:
          message.connectionId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddConnectionBroker {
    const message = { ...baseMsgAddConnectionBroker } as MsgAddConnectionBroker;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.brokerId !== undefined && object.brokerId !== null) {
      message.brokerId = String(object.brokerId);
    } else {
      message.brokerId = "";
    }
    if (object.connectionId !== undefined && object.connectionId !== null) {
      message.connectionId = String(object.connectionId);
    } else {
      message.connectionId = "";
    }
    return message;
  },

  toJSON(message: MsgAddConnectionBroker): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.brokerId !== undefined && (obj.brokerId = message.brokerId);
    message.connectionId !== undefined &&
      (obj.connectionId = message.connectionId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgAddConnectionBroker>
  ): MsgAddConnectionBroker {
    const message = { ...baseMsgAddConnectionBroker } as MsgAddConnectionBroker;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.brokerId !== undefined && object.brokerId !== null) {
      message.brokerId = object.brokerId;
    } else {
      message.brokerId = "";
    }
    if (object.connectionId !== undefined && object.connectionId !== null) {
      message.connectionId = object.connectionId;
    } else {
      message.connectionId = "";
    }
    return message;
  },
};

const baseMsgAddConnectionBrokerResponse: object = {};

export const MsgAddConnectionBrokerResponse = {
  encode(
    _: MsgAddConnectionBrokerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddConnectionBrokerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddConnectionBrokerResponse,
    } as MsgAddConnectionBrokerResponse;
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

  fromJSON(_: any): MsgAddConnectionBrokerResponse {
    const message = {
      ...baseMsgAddConnectionBrokerResponse,
    } as MsgAddConnectionBrokerResponse;
    return message;
  },

  toJSON(_: MsgAddConnectionBrokerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddConnectionBrokerResponse>
  ): MsgAddConnectionBrokerResponse {
    const message = {
      ...baseMsgAddConnectionBrokerResponse,
    } as MsgAddConnectionBrokerResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** Register defines a rpc handler for MsgAddLiquiditySource */
  AddLiquiditySource(
    request: MsgAddLiquiditySource
  ): Promise<MsgAddLiquiditySourceResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddConnectionBroker(
    request: MsgAddConnectionBroker
  ): Promise<MsgAddConnectionBrokerResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  AddLiquiditySource(
    request: MsgAddLiquiditySource
  ): Promise<MsgAddLiquiditySourceResponse> {
    const data = MsgAddLiquiditySource.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.broker.Msg",
      "AddLiquiditySource",
      data
    );
    return promise.then((data) =>
      MsgAddLiquiditySourceResponse.decode(new Reader(data))
    );
  }

  AddConnectionBroker(
    request: MsgAddConnectionBroker
  ): Promise<MsgAddConnectionBrokerResponse> {
    const data = MsgAddConnectionBroker.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.broker.Msg",
      "AddConnectionBroker",
      data
    );
    return promise.then((data) =>
      MsgAddConnectionBrokerResponse.decode(new Reader(data))
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
