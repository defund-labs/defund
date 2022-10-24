/* eslint-disable */
import { DenomTrace } from "../ibc/applications/transfer/v1/transfer";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "defundlabs.defund.broker";

export interface BaseDenoms {
  AtomTrace: DenomTrace | undefined;
  OsmoTrace: DenomTrace | undefined;
}

/** Params defines the parameters for the broker module. */
export interface Params {
  /** set the base denoms */
  base_denoms: BaseDenoms | undefined;
}

const baseBaseDenoms: object = {};

export const BaseDenoms = {
  encode(message: BaseDenoms, writer: Writer = Writer.create()): Writer {
    if (message.AtomTrace !== undefined) {
      DenomTrace.encode(message.AtomTrace, writer.uint32(10).fork()).ldelim();
    }
    if (message.OsmoTrace !== undefined) {
      DenomTrace.encode(message.OsmoTrace, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): BaseDenoms {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBaseDenoms } as BaseDenoms;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.AtomTrace = DenomTrace.decode(reader, reader.uint32());
          break;
        case 2:
          message.OsmoTrace = DenomTrace.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BaseDenoms {
    const message = { ...baseBaseDenoms } as BaseDenoms;
    if (object.AtomTrace !== undefined && object.AtomTrace !== null) {
      message.AtomTrace = DenomTrace.fromJSON(object.AtomTrace);
    } else {
      message.AtomTrace = undefined;
    }
    if (object.OsmoTrace !== undefined && object.OsmoTrace !== null) {
      message.OsmoTrace = DenomTrace.fromJSON(object.OsmoTrace);
    } else {
      message.OsmoTrace = undefined;
    }
    return message;
  },

  toJSON(message: BaseDenoms): unknown {
    const obj: any = {};
    message.AtomTrace !== undefined &&
      (obj.AtomTrace = message.AtomTrace
        ? DenomTrace.toJSON(message.AtomTrace)
        : undefined);
    message.OsmoTrace !== undefined &&
      (obj.OsmoTrace = message.OsmoTrace
        ? DenomTrace.toJSON(message.OsmoTrace)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<BaseDenoms>): BaseDenoms {
    const message = { ...baseBaseDenoms } as BaseDenoms;
    if (object.AtomTrace !== undefined && object.AtomTrace !== null) {
      message.AtomTrace = DenomTrace.fromPartial(object.AtomTrace);
    } else {
      message.AtomTrace = undefined;
    }
    if (object.OsmoTrace !== undefined && object.OsmoTrace !== null) {
      message.OsmoTrace = DenomTrace.fromPartial(object.OsmoTrace);
    } else {
      message.OsmoTrace = undefined;
    }
    return message;
  },
};

const baseParams: object = {};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.base_denoms !== undefined) {
      BaseDenoms.encode(message.base_denoms, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.base_denoms = BaseDenoms.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (object.base_denoms !== undefined && object.base_denoms !== null) {
      message.base_denoms = BaseDenoms.fromJSON(object.base_denoms);
    } else {
      message.base_denoms = undefined;
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.base_denoms !== undefined &&
      (obj.base_denoms = message.base_denoms
        ? BaseDenoms.toJSON(message.base_denoms)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.base_denoms !== undefined && object.base_denoms !== null) {
      message.base_denoms = BaseDenoms.fromPartial(object.base_denoms);
    } else {
      message.base_denoms = undefined;
    }
    return message;
  },
};

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
