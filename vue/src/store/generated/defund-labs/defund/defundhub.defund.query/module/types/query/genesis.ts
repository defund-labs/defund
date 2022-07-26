/* eslint-disable */
import {
  Interquery,
  InterqueryResult,
  InterqueryTimeoutResult,
} from "../query/interquery";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "defundhub.defund.query";

/** GenesisState defines the query module's genesis state. */
export interface GenesisState {
  interqueryList: Interquery[];
  interqueryResultList: InterqueryResult[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  interqueryTimeoutResultList: InterqueryTimeoutResult[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.interqueryList = [];
    message.interqueryResultList = [];
    message.interqueryTimeoutResultList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interqueryList.push(
            Interquery.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.interqueryResultList.push(
            InterqueryResult.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.interqueryTimeoutResultList.push(
            InterqueryTimeoutResult.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.interqueryList = [];
    message.interqueryResultList = [];
    message.interqueryTimeoutResultList = [];
    if (object.interqueryList !== undefined && object.interqueryList !== null) {
      for (const e of object.interqueryList) {
        message.interqueryList.push(Interquery.fromJSON(e));
      }
    }
    if (
      object.interqueryResultList !== undefined &&
      object.interqueryResultList !== null
    ) {
      for (const e of object.interqueryResultList) {
        message.interqueryResultList.push(InterqueryResult.fromJSON(e));
      }
    }
    if (
      object.interqueryTimeoutResultList !== undefined &&
      object.interqueryTimeoutResultList !== null
    ) {
      for (const e of object.interqueryTimeoutResultList) {
        message.interqueryTimeoutResultList.push(
          InterqueryTimeoutResult.fromJSON(e)
        );
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.interqueryList) {
      obj.interqueryList = message.interqueryList.map((e) =>
        e ? Interquery.toJSON(e) : undefined
      );
    } else {
      obj.interqueryList = [];
    }
    if (message.interqueryResultList) {
      obj.interqueryResultList = message.interqueryResultList.map((e) =>
        e ? InterqueryResult.toJSON(e) : undefined
      );
    } else {
      obj.interqueryResultList = [];
    }
    if (message.interqueryTimeoutResultList) {
      obj.interqueryTimeoutResultList = message.interqueryTimeoutResultList.map(
        (e) => (e ? InterqueryTimeoutResult.toJSON(e) : undefined)
      );
    } else {
      obj.interqueryTimeoutResultList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.interqueryList = [];
    message.interqueryResultList = [];
    message.interqueryTimeoutResultList = [];
    if (object.interqueryList !== undefined && object.interqueryList !== null) {
      for (const e of object.interqueryList) {
        message.interqueryList.push(Interquery.fromPartial(e));
      }
    }
    if (
      object.interqueryResultList !== undefined &&
      object.interqueryResultList !== null
    ) {
      for (const e of object.interqueryResultList) {
        message.interqueryResultList.push(InterqueryResult.fromPartial(e));
      }
    }
    if (
      object.interqueryTimeoutResultList !== undefined &&
      object.interqueryTimeoutResultList !== null
    ) {
      for (const e of object.interqueryTimeoutResultList) {
        message.interqueryTimeoutResultList.push(
          InterqueryTimeoutResult.fromPartial(e)
        );
      }
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
