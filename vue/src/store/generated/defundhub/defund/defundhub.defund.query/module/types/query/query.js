/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Interquery } from "../query/interquery";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
export const protobufPackage = "defundhub.defund.query";
const baseQueryGetInterqueryRequest = { index: "" };
export const QueryGetInterqueryRequest = {
    encode(message, writer = Writer.create()) {
        if (message.index !== "") {
            writer.uint32(10).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetInterqueryRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
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
        const message = {
            ...baseQueryGetInterqueryRequest,
        };
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
        message.index !== undefined && (obj.index = message.index);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetInterqueryRequest,
        };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = "";
        }
        return message;
    },
};
const baseQueryGetInterqueryResponse = {};
export const QueryGetInterqueryResponse = {
    encode(message, writer = Writer.create()) {
        if (message.interquery !== undefined) {
            Interquery.encode(message.interquery, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetInterqueryResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.interquery = Interquery.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryGetInterqueryResponse,
        };
        if (object.interquery !== undefined && object.interquery !== null) {
            message.interquery = Interquery.fromJSON(object.interquery);
        }
        else {
            message.interquery = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.interquery !== undefined &&
            (obj.interquery = message.interquery
                ? Interquery.toJSON(message.interquery)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetInterqueryResponse,
        };
        if (object.interquery !== undefined && object.interquery !== null) {
            message.interquery = Interquery.fromPartial(object.interquery);
        }
        else {
            message.interquery = undefined;
        }
        return message;
    },
};
const baseQueryAllInterqueryRequest = {};
export const QueryAllInterqueryRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllInterqueryRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryAllInterqueryRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllInterqueryRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllInterqueryResponse = {};
export const QueryAllInterqueryResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.interquery) {
            Interquery.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllInterqueryResponse,
        };
        message.interquery = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.interquery.push(Interquery.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryAllInterqueryResponse,
        };
        message.interquery = [];
        if (object.interquery !== undefined && object.interquery !== null) {
            for (const e of object.interquery) {
                message.interquery.push(Interquery.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.interquery) {
            obj.interquery = message.interquery.map((e) => e ? Interquery.toJSON(e) : undefined);
        }
        else {
            obj.interquery = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllInterqueryResponse,
        };
        message.interquery = [];
        if (object.interquery !== undefined && object.interquery !== null) {
            for (const e of object.interquery) {
                message.interquery.push(Interquery.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Interquery(request) {
        const data = QueryGetInterqueryRequest.encode(request).finish();
        const promise = this.rpc.request("defundhub.defund.query.Query", "Interquery", data);
        return promise.then((data) => QueryGetInterqueryResponse.decode(new Reader(data)));
    }
    InterqueryAll(request) {
        const data = QueryAllInterqueryRequest.encode(request).finish();
        const promise = this.rpc.request("defundhub.defund.query.Query", "InterqueryAll", data);
        return promise.then((data) => QueryAllInterqueryResponse.decode(new Reader(data)));
    }
}
