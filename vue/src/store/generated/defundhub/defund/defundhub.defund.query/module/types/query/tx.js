/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "defundhub.defund.query";
const baseMsgCreateInterquery = {
    creator: "",
    index: "",
    height: "",
    path: "",
    chainId: "",
    typeName: "",
};
export const MsgCreateInterquery = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateInterquery };
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
    fromJSON(object) {
        const message = { ...baseMsgCreateInterquery };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = "";
        }
        if (object.height !== undefined && object.height !== null) {
            message.height = String(object.height);
        }
        else {
            message.height = "";
        }
        if (object.path !== undefined && object.path !== null) {
            message.path = String(object.path);
        }
        else {
            message.path = "";
        }
        if (object.chainId !== undefined && object.chainId !== null) {
            message.chainId = String(object.chainId);
        }
        else {
            message.chainId = "";
        }
        if (object.typeName !== undefined && object.typeName !== null) {
            message.typeName = String(object.typeName);
        }
        else {
            message.typeName = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.index !== undefined && (obj.index = message.index);
        message.height !== undefined && (obj.height = message.height);
        message.path !== undefined && (obj.path = message.path);
        message.chainId !== undefined && (obj.chainId = message.chainId);
        message.typeName !== undefined && (obj.typeName = message.typeName);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateInterquery };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = "";
        }
        if (object.height !== undefined && object.height !== null) {
            message.height = object.height;
        }
        else {
            message.height = "";
        }
        if (object.path !== undefined && object.path !== null) {
            message.path = object.path;
        }
        else {
            message.path = "";
        }
        if (object.chainId !== undefined && object.chainId !== null) {
            message.chainId = object.chainId;
        }
        else {
            message.chainId = "";
        }
        if (object.typeName !== undefined && object.typeName !== null) {
            message.typeName = object.typeName;
        }
        else {
            message.typeName = "";
        }
        return message;
    },
};
const baseMsgCreateInterqueryResponse = {};
export const MsgCreateInterqueryResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateInterqueryResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgCreateInterqueryResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgCreateInterqueryResponse,
        };
        return message;
    },
};
const baseMsgUpdateInterquery = {
    creator: "",
    index: "",
    height: "",
    path: "",
    chainId: "",
    typeName: "",
};
export const MsgUpdateInterquery = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateInterquery };
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
    fromJSON(object) {
        const message = { ...baseMsgUpdateInterquery };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = "";
        }
        if (object.height !== undefined && object.height !== null) {
            message.height = String(object.height);
        }
        else {
            message.height = "";
        }
        if (object.path !== undefined && object.path !== null) {
            message.path = String(object.path);
        }
        else {
            message.path = "";
        }
        if (object.chainId !== undefined && object.chainId !== null) {
            message.chainId = String(object.chainId);
        }
        else {
            message.chainId = "";
        }
        if (object.typeName !== undefined && object.typeName !== null) {
            message.typeName = String(object.typeName);
        }
        else {
            message.typeName = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.index !== undefined && (obj.index = message.index);
        message.height !== undefined && (obj.height = message.height);
        message.path !== undefined && (obj.path = message.path);
        message.chainId !== undefined && (obj.chainId = message.chainId);
        message.typeName !== undefined && (obj.typeName = message.typeName);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateInterquery };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = "";
        }
        if (object.height !== undefined && object.height !== null) {
            message.height = object.height;
        }
        else {
            message.height = "";
        }
        if (object.path !== undefined && object.path !== null) {
            message.path = object.path;
        }
        else {
            message.path = "";
        }
        if (object.chainId !== undefined && object.chainId !== null) {
            message.chainId = object.chainId;
        }
        else {
            message.chainId = "";
        }
        if (object.typeName !== undefined && object.typeName !== null) {
            message.typeName = object.typeName;
        }
        else {
            message.typeName = "";
        }
        return message;
    },
};
const baseMsgUpdateInterqueryResponse = {};
export const MsgUpdateInterqueryResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgUpdateInterqueryResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgUpdateInterqueryResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgUpdateInterqueryResponse,
        };
        return message;
    },
};
const baseMsgDeleteInterquery = { creator: "", index: "" };
export const MsgDeleteInterquery = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.index !== "") {
            writer.uint32(18).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteInterquery };
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
    fromJSON(object) {
        const message = { ...baseMsgDeleteInterquery };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.index !== undefined && (obj.index = message.index);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDeleteInterquery };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = "";
        }
        return message;
    },
};
const baseMsgDeleteInterqueryResponse = {};
export const MsgDeleteInterqueryResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDeleteInterqueryResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgDeleteInterqueryResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgDeleteInterqueryResponse,
        };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateInterquery(request) {
        const data = MsgCreateInterquery.encode(request).finish();
        const promise = this.rpc.request("defundhub.defund.query.Msg", "CreateInterquery", data);
        return promise.then((data) => MsgCreateInterqueryResponse.decode(new Reader(data)));
    }
    UpdateInterquery(request) {
        const data = MsgUpdateInterquery.encode(request).finish();
        const promise = this.rpc.request("defundhub.defund.query.Msg", "UpdateInterquery", data);
        return promise.then((data) => MsgUpdateInterqueryResponse.decode(new Reader(data)));
    }
    DeleteInterquery(request) {
        const data = MsgDeleteInterquery.encode(request).finish();
        const promise = this.rpc.request("defundhub.defund.query.Msg", "DeleteInterquery", data);
        return promise.then((data) => MsgDeleteInterqueryResponse.decode(new Reader(data)));
    }
}
