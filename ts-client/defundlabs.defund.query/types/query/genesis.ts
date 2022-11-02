/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Interquery, InterqueryResult, InterqueryTimeoutResult } from "./interquery";

export const protobufPackage = "defundlabs.defund.query";

/** GenesisState defines the query module's genesis state. */
export interface GenesisState {
  interqueryList: Interquery[];
  interqueryResultList: InterqueryResult[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  interqueryTimeoutResultList: InterqueryTimeoutResult[];
}

function createBaseGenesisState(): GenesisState {
  return { interqueryList: [], interqueryResultList: [], interqueryTimeoutResultList: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.interqueryList) {
      Interquery.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.interqueryResultList) {
      InterqueryResult.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.interqueryTimeoutResultList) {
      InterqueryTimeoutResult.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interqueryList.push(Interquery.decode(reader, reader.uint32()));
          break;
        case 2:
          message.interqueryResultList.push(InterqueryResult.decode(reader, reader.uint32()));
          break;
        case 3:
          message.interqueryTimeoutResultList.push(InterqueryTimeoutResult.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      interqueryList: Array.isArray(object?.interqueryList)
        ? object.interqueryList.map((e: any) => Interquery.fromJSON(e))
        : [],
      interqueryResultList: Array.isArray(object?.interqueryResultList)
        ? object.interqueryResultList.map((e: any) => InterqueryResult.fromJSON(e))
        : [],
      interqueryTimeoutResultList: Array.isArray(object?.interqueryTimeoutResultList)
        ? object.interqueryTimeoutResultList.map((e: any) => InterqueryTimeoutResult.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.interqueryList) {
      obj.interqueryList = message.interqueryList.map((e) => e ? Interquery.toJSON(e) : undefined);
    } else {
      obj.interqueryList = [];
    }
    if (message.interqueryResultList) {
      obj.interqueryResultList = message.interqueryResultList.map((e) => e ? InterqueryResult.toJSON(e) : undefined);
    } else {
      obj.interqueryResultList = [];
    }
    if (message.interqueryTimeoutResultList) {
      obj.interqueryTimeoutResultList = message.interqueryTimeoutResultList.map((e) =>
        e ? InterqueryTimeoutResult.toJSON(e) : undefined
      );
    } else {
      obj.interqueryTimeoutResultList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.interqueryList = object.interqueryList?.map((e) => Interquery.fromPartial(e)) || [];
    message.interqueryResultList = object.interqueryResultList?.map((e) => InterqueryResult.fromPartial(e)) || [];
    message.interqueryTimeoutResultList =
      object.interqueryTimeoutResultList?.map((e) => InterqueryTimeoutResult.fromPartial(e)) || [];
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
