/* eslint-disable */
import { Interquery } from "../query/interquery";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "defundhub.defund.query";
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.interqueryList) {
            Interquery.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.interqueryList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.interqueryList.push(Interquery.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.interqueryList = [];
        if (object.interqueryList !== undefined && object.interqueryList !== null) {
            for (const e of object.interqueryList) {
                message.interqueryList.push(Interquery.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.interqueryList) {
            obj.interqueryList = message.interqueryList.map((e) => e ? Interquery.toJSON(e) : undefined);
        }
        else {
            obj.interqueryList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.interqueryList = [];
        if (object.interqueryList !== undefined && object.interqueryList !== null) {
            for (const e of object.interqueryList) {
                message.interqueryList.push(Interquery.fromPartial(e));
            }
        }
        return message;
    },
};
