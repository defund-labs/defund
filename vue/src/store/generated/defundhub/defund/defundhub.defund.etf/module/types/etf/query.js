/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Fund } from "../etf/fund";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
export const protobufPackage = "defundhub.defund.etf";
const baseQueryGetFundRequest = { index: "" };
export const QueryGetFundRequest = {
    encode(message, writer = Writer.create()) {
        if (message.index !== "") {
            writer.uint32(10).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetFundRequest };
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
        const message = { ...baseQueryGetFundRequest };
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
        const message = { ...baseQueryGetFundRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = "";
        }
        return message;
    },
};
const baseQueryGetFundResponse = {};
export const QueryGetFundResponse = {
    encode(message, writer = Writer.create()) {
        if (message.fund !== undefined) {
            Fund.encode(message.fund, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetFundResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.fund = Fund.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetFundResponse };
        if (object.fund !== undefined && object.fund !== null) {
            message.fund = Fund.fromJSON(object.fund);
        }
        else {
            message.fund = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.fund !== undefined &&
            (obj.fund = message.fund ? Fund.toJSON(message.fund) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetFundResponse };
        if (object.fund !== undefined && object.fund !== null) {
            message.fund = Fund.fromPartial(object.fund);
        }
        else {
            message.fund = undefined;
        }
        return message;
    },
};
const baseQueryAllFundRequest = {};
export const QueryAllFundRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllFundRequest };
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
        const message = { ...baseQueryAllFundRequest };
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
        const message = { ...baseQueryAllFundRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllFundResponse = {};
export const QueryAllFundResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.fund) {
            Fund.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllFundResponse };
        message.fund = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.fund.push(Fund.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllFundResponse };
        message.fund = [];
        if (object.fund !== undefined && object.fund !== null) {
            for (const e of object.fund) {
                message.fund.push(Fund.fromJSON(e));
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
        if (message.fund) {
            obj.fund = message.fund.map((e) => (e ? Fund.toJSON(e) : undefined));
        }
        else {
            obj.fund = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllFundResponse };
        message.fund = [];
        if (object.fund !== undefined && object.fund !== null) {
            for (const e of object.fund) {
                message.fund.push(Fund.fromPartial(e));
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
const baseQueryFundPriceRequest = { ticker: "" };
export const QueryFundPriceRequest = {
    encode(message, writer = Writer.create()) {
        if (message.ticker !== "") {
            writer.uint32(10).string(message.ticker);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryFundPriceRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.ticker = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryFundPriceRequest };
        if (object.ticker !== undefined && object.ticker !== null) {
            message.ticker = String(object.ticker);
        }
        else {
            message.ticker = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.ticker !== undefined && (obj.ticker = message.ticker);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryFundPriceRequest };
        if (object.ticker !== undefined && object.ticker !== null) {
            message.ticker = object.ticker;
        }
        else {
            message.ticker = "";
        }
        return message;
    },
};
const baseQueryFundPriceResponse = {};
export const QueryFundPriceResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryFundPriceResponse };
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
        const message = { ...baseQueryFundPriceResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryFundPriceResponse };
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Fund(request) {
        const data = QueryGetFundRequest.encode(request).finish();
        const promise = this.rpc.request("defundhub.defund.etf.Query", "Fund", data);
        return promise.then((data) => QueryGetFundResponse.decode(new Reader(data)));
    }
    FundAll(request) {
        const data = QueryAllFundRequest.encode(request).finish();
        const promise = this.rpc.request("defundhub.defund.etf.Query", "FundAll", data);
        return promise.then((data) => QueryAllFundResponse.decode(new Reader(data)));
    }
    FundPrice(request) {
        const data = QueryFundPriceRequest.encode(request).finish();
        const promise = this.rpc.request("defundhub.defund.etf.Query", "FundPrice", data);
        return promise.then((data) => QueryFundPriceResponse.decode(new Reader(data)));
    }
}
