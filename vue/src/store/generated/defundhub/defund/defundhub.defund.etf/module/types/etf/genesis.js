/* eslint-disable */
import { Fund } from "../etf/fund";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "defundhub.defund.etf";
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.fundList) {
            Fund.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.fundList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.fundList.push(Fund.decode(reader, reader.uint32()));
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
        message.fundList = [];
        if (object.fundList !== undefined && object.fundList !== null) {
            for (const e of object.fundList) {
                message.fundList.push(Fund.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.fundList) {
            obj.fundList = message.fundList.map((e) => e ? Fund.toJSON(e) : undefined);
        }
        else {
            obj.fundList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.fundList = [];
        if (object.fundList !== undefined && object.fundList !== null) {
            for (const e of object.fundList) {
                message.fundList.push(Fund.fromPartial(e));
            }
        }
        return message;
    },
};
