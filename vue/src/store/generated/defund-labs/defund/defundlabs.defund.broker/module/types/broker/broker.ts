/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "defundlabs.defund.broker";

export interface Source {
  pool_id: number;
  interquery_id: string;
  status: string;
  append: Uint8Array;
}

export interface Broker {
  id: string;
  connection_id: string;
  pools: Source[];
  baseDenom: string;
  status: string;
}

export interface Transfer {
  id: string;
  channel: string;
  sequence: number;
  status: string;
  token: Coin | undefined;
  sender: string;
  receiver: string;
}

const baseSource: object = { pool_id: 0, interquery_id: "", status: "" };

export const Source = {
  encode(message: Source, writer: Writer = Writer.create()): Writer {
    if (message.pool_id !== 0) {
      writer.uint32(8).uint64(message.pool_id);
    }
    if (message.interquery_id !== "") {
      writer.uint32(18).string(message.interquery_id);
    }
    if (message.status !== "") {
      writer.uint32(26).string(message.status);
    }
    if (message.append.length !== 0) {
      writer.uint32(34).bytes(message.append);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Source {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSource } as Source;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pool_id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.interquery_id = reader.string();
          break;
        case 3:
          message.status = reader.string();
          break;
        case 4:
          message.append = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Source {
    const message = { ...baseSource } as Source;
    if (object.pool_id !== undefined && object.pool_id !== null) {
      message.pool_id = Number(object.pool_id);
    } else {
      message.pool_id = 0;
    }
    if (object.interquery_id !== undefined && object.interquery_id !== null) {
      message.interquery_id = String(object.interquery_id);
    } else {
      message.interquery_id = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.append !== undefined && object.append !== null) {
      message.append = bytesFromBase64(object.append);
    }
    return message;
  },

  toJSON(message: Source): unknown {
    const obj: any = {};
    message.pool_id !== undefined && (obj.pool_id = message.pool_id);
    message.interquery_id !== undefined &&
      (obj.interquery_id = message.interquery_id);
    message.status !== undefined && (obj.status = message.status);
    message.append !== undefined &&
      (obj.append = base64FromBytes(
        message.append !== undefined ? message.append : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<Source>): Source {
    const message = { ...baseSource } as Source;
    if (object.pool_id !== undefined && object.pool_id !== null) {
      message.pool_id = object.pool_id;
    } else {
      message.pool_id = 0;
    }
    if (object.interquery_id !== undefined && object.interquery_id !== null) {
      message.interquery_id = object.interquery_id;
    } else {
      message.interquery_id = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.append !== undefined && object.append !== null) {
      message.append = object.append;
    } else {
      message.append = new Uint8Array();
    }
    return message;
  },
};

const baseBroker: object = {
  id: "",
  connection_id: "",
  baseDenom: "",
  status: "",
};

export const Broker = {
  encode(message: Broker, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.connection_id !== "") {
      writer.uint32(18).string(message.connection_id);
    }
    for (const v of message.pools) {
      Source.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.baseDenom !== "") {
      writer.uint32(34).string(message.baseDenom);
    }
    if (message.status !== "") {
      writer.uint32(42).string(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Broker {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBroker } as Broker;
    message.pools = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.connection_id = reader.string();
          break;
        case 3:
          message.pools.push(Source.decode(reader, reader.uint32()));
          break;
        case 4:
          message.baseDenom = reader.string();
          break;
        case 5:
          message.status = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Broker {
    const message = { ...baseBroker } as Broker;
    message.pools = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.connection_id !== undefined && object.connection_id !== null) {
      message.connection_id = String(object.connection_id);
    } else {
      message.connection_id = "";
    }
    if (object.pools !== undefined && object.pools !== null) {
      for (const e of object.pools) {
        message.pools.push(Source.fromJSON(e));
      }
    }
    if (object.baseDenom !== undefined && object.baseDenom !== null) {
      message.baseDenom = String(object.baseDenom);
    } else {
      message.baseDenom = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    return message;
  },

  toJSON(message: Broker): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.connection_id !== undefined &&
      (obj.connection_id = message.connection_id);
    if (message.pools) {
      obj.pools = message.pools.map((e) => (e ? Source.toJSON(e) : undefined));
    } else {
      obj.pools = [];
    }
    message.baseDenom !== undefined && (obj.baseDenom = message.baseDenom);
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(object: DeepPartial<Broker>): Broker {
    const message = { ...baseBroker } as Broker;
    message.pools = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.connection_id !== undefined && object.connection_id !== null) {
      message.connection_id = object.connection_id;
    } else {
      message.connection_id = "";
    }
    if (object.pools !== undefined && object.pools !== null) {
      for (const e of object.pools) {
        message.pools.push(Source.fromPartial(e));
      }
    }
    if (object.baseDenom !== undefined && object.baseDenom !== null) {
      message.baseDenom = object.baseDenom;
    } else {
      message.baseDenom = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    return message;
  },
};

const baseTransfer: object = {
  id: "",
  channel: "",
  sequence: 0,
  status: "",
  sender: "",
  receiver: "",
};

export const Transfer = {
  encode(message: Transfer, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.channel !== "") {
      writer.uint32(18).string(message.channel);
    }
    if (message.sequence !== 0) {
      writer.uint32(24).uint64(message.sequence);
    }
    if (message.status !== "") {
      writer.uint32(34).string(message.status);
    }
    if (message.token !== undefined) {
      Coin.encode(message.token, writer.uint32(42).fork()).ldelim();
    }
    if (message.sender !== "") {
      writer.uint32(50).string(message.sender);
    }
    if (message.receiver !== "") {
      writer.uint32(58).string(message.receiver);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Transfer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTransfer } as Transfer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.channel = reader.string();
          break;
        case 3:
          message.sequence = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.status = reader.string();
          break;
        case 5:
          message.token = Coin.decode(reader, reader.uint32());
          break;
        case 6:
          message.sender = reader.string();
          break;
        case 7:
          message.receiver = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Transfer {
    const message = { ...baseTransfer } as Transfer;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = String(object.channel);
    } else {
      message.channel = "";
    }
    if (object.sequence !== undefined && object.sequence !== null) {
      message.sequence = Number(object.sequence);
    } else {
      message.sequence = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = Coin.fromJSON(object.token);
    } else {
      message.token = undefined;
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    return message;
  },

  toJSON(message: Transfer): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.channel !== undefined && (obj.channel = message.channel);
    message.sequence !== undefined && (obj.sequence = message.sequence);
    message.status !== undefined && (obj.status = message.status);
    message.token !== undefined &&
      (obj.token = message.token ? Coin.toJSON(message.token) : undefined);
    message.sender !== undefined && (obj.sender = message.sender);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    return obj;
  },

  fromPartial(object: DeepPartial<Transfer>): Transfer {
    const message = { ...baseTransfer } as Transfer;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = object.channel;
    } else {
      message.channel = "";
    }
    if (object.sequence !== undefined && object.sequence !== null) {
      message.sequence = object.sequence;
    } else {
      message.sequence = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = Coin.fromPartial(object.token);
    } else {
      message.token = undefined;
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
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

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
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
