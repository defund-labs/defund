/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Broker } from "./broker";

export const protobufPackage = "defundlabs.defund.broker";

/** QueryInterchainAccountFromAddressRequest is the request type for the Query/InterchainAccountAddress RPC */
export interface QueryInterchainAccountFromAddressRequest {
  owner: string;
  connectionId: string;
}

/** QueryInterchainAccountFromAddressResponse the response type for the Query/InterchainAccountAddress RPC */
export interface QueryInterchainAccountFromAddressResponse {
  interchainAccountAddress: string;
}

/** QueryBrokerRequest is the request type for the Query/Broker RPC */
export interface QueryBrokerRequest {
  broker: string;
}

/** QueryBrokerResponse the response type for the Query/Broker RPC */
export interface QueryBrokerResponse {
  broker: Broker | undefined;
}

/** QueryBrokersRequest is the request type for the Query/Brokers RPC */
export interface QueryBrokersRequest {
  pagination: PageRequest | undefined;
}

/** QueryBrokersResponse the response type for the Query/Brokers RPC */
export interface QueryBrokersResponse {
  brokers: Broker[];
  pagination: PageResponse | undefined;
}

function createBaseQueryInterchainAccountFromAddressRequest(): QueryInterchainAccountFromAddressRequest {
  return { owner: "", connectionId: "" };
}

export const QueryInterchainAccountFromAddressRequest = {
  encode(message: QueryInterchainAccountFromAddressRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.connectionId !== "") {
      writer.uint32(18).string(message.connectionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryInterchainAccountFromAddressRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryInterchainAccountFromAddressRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.connectionId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryInterchainAccountFromAddressRequest {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      connectionId: isSet(object.connectionId) ? String(object.connectionId) : "",
    };
  },

  toJSON(message: QueryInterchainAccountFromAddressRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.connectionId !== undefined && (obj.connectionId = message.connectionId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryInterchainAccountFromAddressRequest>, I>>(
    object: I,
  ): QueryInterchainAccountFromAddressRequest {
    const message = createBaseQueryInterchainAccountFromAddressRequest();
    message.owner = object.owner ?? "";
    message.connectionId = object.connectionId ?? "";
    return message;
  },
};

function createBaseQueryInterchainAccountFromAddressResponse(): QueryInterchainAccountFromAddressResponse {
  return { interchainAccountAddress: "" };
}

export const QueryInterchainAccountFromAddressResponse = {
  encode(message: QueryInterchainAccountFromAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.interchainAccountAddress !== "") {
      writer.uint32(10).string(message.interchainAccountAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryInterchainAccountFromAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryInterchainAccountFromAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interchainAccountAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryInterchainAccountFromAddressResponse {
    return {
      interchainAccountAddress: isSet(object.interchainAccountAddress) ? String(object.interchainAccountAddress) : "",
    };
  },

  toJSON(message: QueryInterchainAccountFromAddressResponse): unknown {
    const obj: any = {};
    message.interchainAccountAddress !== undefined && (obj.interchainAccountAddress = message.interchainAccountAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryInterchainAccountFromAddressResponse>, I>>(
    object: I,
  ): QueryInterchainAccountFromAddressResponse {
    const message = createBaseQueryInterchainAccountFromAddressResponse();
    message.interchainAccountAddress = object.interchainAccountAddress ?? "";
    return message;
  },
};

function createBaseQueryBrokerRequest(): QueryBrokerRequest {
  return { broker: "" };
}

export const QueryBrokerRequest = {
  encode(message: QueryBrokerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.broker !== "") {
      writer.uint32(10).string(message.broker);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryBrokerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryBrokerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.broker = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBrokerRequest {
    return { broker: isSet(object.broker) ? String(object.broker) : "" };
  },

  toJSON(message: QueryBrokerRequest): unknown {
    const obj: any = {};
    message.broker !== undefined && (obj.broker = message.broker);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryBrokerRequest>, I>>(object: I): QueryBrokerRequest {
    const message = createBaseQueryBrokerRequest();
    message.broker = object.broker ?? "";
    return message;
  },
};

function createBaseQueryBrokerResponse(): QueryBrokerResponse {
  return { broker: undefined };
}

export const QueryBrokerResponse = {
  encode(message: QueryBrokerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.broker !== undefined) {
      Broker.encode(message.broker, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryBrokerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryBrokerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.broker = Broker.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBrokerResponse {
    return { broker: isSet(object.broker) ? Broker.fromJSON(object.broker) : undefined };
  },

  toJSON(message: QueryBrokerResponse): unknown {
    const obj: any = {};
    message.broker !== undefined && (obj.broker = message.broker ? Broker.toJSON(message.broker) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryBrokerResponse>, I>>(object: I): QueryBrokerResponse {
    const message = createBaseQueryBrokerResponse();
    message.broker = (object.broker !== undefined && object.broker !== null)
      ? Broker.fromPartial(object.broker)
      : undefined;
    return message;
  },
};

function createBaseQueryBrokersRequest(): QueryBrokersRequest {
  return { pagination: undefined };
}

export const QueryBrokersRequest = {
  encode(message: QueryBrokersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryBrokersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryBrokersRequest();
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

  fromJSON(object: any): QueryBrokersRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryBrokersRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryBrokersRequest>, I>>(object: I): QueryBrokersRequest {
    const message = createBaseQueryBrokersRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryBrokersResponse(): QueryBrokersResponse {
  return { brokers: [], pagination: undefined };
}

export const QueryBrokersResponse = {
  encode(message: QueryBrokersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.brokers) {
      Broker.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryBrokersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryBrokersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.brokers.push(Broker.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryBrokersResponse {
    return {
      brokers: Array.isArray(object?.brokers) ? object.brokers.map((e: any) => Broker.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryBrokersResponse): unknown {
    const obj: any = {};
    if (message.brokers) {
      obj.brokers = message.brokers.map((e) => e ? Broker.toJSON(e) : undefined);
    } else {
      obj.brokers = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryBrokersResponse>, I>>(object: I): QueryBrokersResponse {
    const message = createBaseQueryBrokersResponse();
    message.brokers = object.brokers?.map((e) => Broker.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** QueryInterchainAccountFromAddress returns the interchain account for given owner address on a given connection pair */
  InterchainAccountFromAddress(
    request: QueryInterchainAccountFromAddressRequest,
  ): Promise<QueryInterchainAccountFromAddressResponse>;
  /** QueryBrokerRequest returns the broker based on the broker id requested */
  Broker(request: QueryBrokerRequest): Promise<QueryBrokerResponse>;
  /** QueryBrokersRequest returns all brokers */
  Brokers(request: QueryBrokersRequest): Promise<QueryBrokersResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.InterchainAccountFromAddress = this.InterchainAccountFromAddress.bind(this);
    this.Broker = this.Broker.bind(this);
    this.Brokers = this.Brokers.bind(this);
  }
  InterchainAccountFromAddress(
    request: QueryInterchainAccountFromAddressRequest,
  ): Promise<QueryInterchainAccountFromAddressResponse> {
    const data = QueryInterchainAccountFromAddressRequest.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.broker.Query", "InterchainAccountFromAddress", data);
    return promise.then((data) => QueryInterchainAccountFromAddressResponse.decode(new _m0.Reader(data)));
  }

  Broker(request: QueryBrokerRequest): Promise<QueryBrokerResponse> {
    const data = QueryBrokerRequest.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.broker.Query", "Broker", data);
    return promise.then((data) => QueryBrokerResponse.decode(new _m0.Reader(data)));
  }

  Brokers(request: QueryBrokersRequest): Promise<QueryBrokersResponse> {
    const data = QueryBrokersRequest.encode(request).finish();
    const promise = this.rpc.request("defundlabs.defund.broker.Query", "Brokers", data);
    return promise.then((data) => QueryBrokersResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
