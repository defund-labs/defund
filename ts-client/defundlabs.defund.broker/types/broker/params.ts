/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DenomTrace } from "../ibc/applications/transfer/v1/transfer";

export const protobufPackage = "defundlabs.defund.broker";

export interface BaseDenoms {
  AtomTrace: DenomTrace | undefined;
  OsmoTrace: DenomTrace | undefined;
}

/** Params defines the parameters for the broker module. */
export interface Params {
  /** set the base denoms */
  baseDenoms: BaseDenoms | undefined;
}

function createBaseBaseDenoms(): BaseDenoms {
  return { AtomTrace: undefined, OsmoTrace: undefined };
}

export const BaseDenoms = {
  encode(message: BaseDenoms, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.AtomTrace !== undefined) {
      DenomTrace.encode(message.AtomTrace, writer.uint32(10).fork()).ldelim();
    }
    if (message.OsmoTrace !== undefined) {
      DenomTrace.encode(message.OsmoTrace, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BaseDenoms {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBaseDenoms();
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
    return {
      AtomTrace: isSet(object.AtomTrace) ? DenomTrace.fromJSON(object.AtomTrace) : undefined,
      OsmoTrace: isSet(object.OsmoTrace) ? DenomTrace.fromJSON(object.OsmoTrace) : undefined,
    };
  },

  toJSON(message: BaseDenoms): unknown {
    const obj: any = {};
    message.AtomTrace !== undefined
      && (obj.AtomTrace = message.AtomTrace ? DenomTrace.toJSON(message.AtomTrace) : undefined);
    message.OsmoTrace !== undefined
      && (obj.OsmoTrace = message.OsmoTrace ? DenomTrace.toJSON(message.OsmoTrace) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BaseDenoms>, I>>(object: I): BaseDenoms {
    const message = createBaseBaseDenoms();
    message.AtomTrace = (object.AtomTrace !== undefined && object.AtomTrace !== null)
      ? DenomTrace.fromPartial(object.AtomTrace)
      : undefined;
    message.OsmoTrace = (object.OsmoTrace !== undefined && object.OsmoTrace !== null)
      ? DenomTrace.fromPartial(object.OsmoTrace)
      : undefined;
    return message;
  },
};

function createBaseParams(): Params {
  return { baseDenoms: undefined };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.baseDenoms !== undefined) {
      BaseDenoms.encode(message.baseDenoms, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.baseDenoms = BaseDenoms.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    return { baseDenoms: isSet(object.baseDenoms) ? BaseDenoms.fromJSON(object.baseDenoms) : undefined };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.baseDenoms !== undefined
      && (obj.baseDenoms = message.baseDenoms ? BaseDenoms.toJSON(message.baseDenoms) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.baseDenoms = (object.baseDenoms !== undefined && object.baseDenoms !== null)
      ? BaseDenoms.fromPartial(object.baseDenoms)
      : undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
