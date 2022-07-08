/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Broker } from "../broker/broker";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "defundlabs.defund.broker";

/** QueryInterchainAccountFromAddressRequest is the request type for the Query/InterchainAccountAddress RPC */
export interface QueryInterchainAccountFromAddressRequest {
  owner: string;
  connection_id: string;
}

/** QueryInterchainAccountFromAddressResponse the response type for the Query/InterchainAccountAddress RPC */
export interface QueryInterchainAccountFromAddressResponse {
  interchain_account_address: string;
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

const baseQueryInterchainAccountFromAddressRequest: object = {
  owner: "",
  connection_id: "",
};

export const QueryInterchainAccountFromAddressRequest = {
  encode(
    message: QueryInterchainAccountFromAddressRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.connection_id !== "") {
      writer.uint32(18).string(message.connection_id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryInterchainAccountFromAddressRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryInterchainAccountFromAddressRequest,
    } as QueryInterchainAccountFromAddressRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.connection_id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryInterchainAccountFromAddressRequest {
    const message = {
      ...baseQueryInterchainAccountFromAddressRequest,
    } as QueryInterchainAccountFromAddressRequest;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.connection_id !== undefined && object.connection_id !== null) {
      message.connection_id = String(object.connection_id);
    } else {
      message.connection_id = "";
    }
    return message;
  },

  toJSON(message: QueryInterchainAccountFromAddressRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.connection_id !== undefined &&
      (obj.connection_id = message.connection_id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryInterchainAccountFromAddressRequest>
  ): QueryInterchainAccountFromAddressRequest {
    const message = {
      ...baseQueryInterchainAccountFromAddressRequest,
    } as QueryInterchainAccountFromAddressRequest;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.connection_id !== undefined && object.connection_id !== null) {
      message.connection_id = object.connection_id;
    } else {
      message.connection_id = "";
    }
    return message;
  },
};

const baseQueryInterchainAccountFromAddressResponse: object = {
  interchain_account_address: "",
};

export const QueryInterchainAccountFromAddressResponse = {
  encode(
    message: QueryInterchainAccountFromAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.interchain_account_address !== "") {
      writer.uint32(10).string(message.interchain_account_address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryInterchainAccountFromAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryInterchainAccountFromAddressResponse,
    } as QueryInterchainAccountFromAddressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interchain_account_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryInterchainAccountFromAddressResponse {
    const message = {
      ...baseQueryInterchainAccountFromAddressResponse,
    } as QueryInterchainAccountFromAddressResponse;
    if (
      object.interchain_account_address !== undefined &&
      object.interchain_account_address !== null
    ) {
      message.interchain_account_address = String(
        object.interchain_account_address
      );
    } else {
      message.interchain_account_address = "";
    }
    return message;
  },

  toJSON(message: QueryInterchainAccountFromAddressResponse): unknown {
    const obj: any = {};
    message.interchain_account_address !== undefined &&
      (obj.interchain_account_address = message.interchain_account_address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryInterchainAccountFromAddressResponse>
  ): QueryInterchainAccountFromAddressResponse {
    const message = {
      ...baseQueryInterchainAccountFromAddressResponse,
    } as QueryInterchainAccountFromAddressResponse;
    if (
      object.interchain_account_address !== undefined &&
      object.interchain_account_address !== null
    ) {
      message.interchain_account_address = object.interchain_account_address;
    } else {
      message.interchain_account_address = "";
    }
    return message;
  },
};

const baseQueryBrokerRequest: object = { broker: "" };

export const QueryBrokerRequest = {
  encode(
    message: QueryBrokerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.broker !== "") {
      writer.uint32(10).string(message.broker);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBrokerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBrokerRequest } as QueryBrokerRequest;
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
    const message = { ...baseQueryBrokerRequest } as QueryBrokerRequest;
    if (object.broker !== undefined && object.broker !== null) {
      message.broker = String(object.broker);
    } else {
      message.broker = "";
    }
    return message;
  },

  toJSON(message: QueryBrokerRequest): unknown {
    const obj: any = {};
    message.broker !== undefined && (obj.broker = message.broker);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBrokerRequest>): QueryBrokerRequest {
    const message = { ...baseQueryBrokerRequest } as QueryBrokerRequest;
    if (object.broker !== undefined && object.broker !== null) {
      message.broker = object.broker;
    } else {
      message.broker = "";
    }
    return message;
  },
};

const baseQueryBrokerResponse: object = {};

export const QueryBrokerResponse = {
  encode(
    message: QueryBrokerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.broker !== undefined) {
      Broker.encode(message.broker, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBrokerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBrokerResponse } as QueryBrokerResponse;
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
    const message = { ...baseQueryBrokerResponse } as QueryBrokerResponse;
    if (object.broker !== undefined && object.broker !== null) {
      message.broker = Broker.fromJSON(object.broker);
    } else {
      message.broker = undefined;
    }
    return message;
  },

  toJSON(message: QueryBrokerResponse): unknown {
    const obj: any = {};
    message.broker !== undefined &&
      (obj.broker = message.broker ? Broker.toJSON(message.broker) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBrokerResponse>): QueryBrokerResponse {
    const message = { ...baseQueryBrokerResponse } as QueryBrokerResponse;
    if (object.broker !== undefined && object.broker !== null) {
      message.broker = Broker.fromPartial(object.broker);
    } else {
      message.broker = undefined;
    }
    return message;
  },
};

const baseQueryBrokersRequest: object = {};

export const QueryBrokersRequest = {
  encode(
    message: QueryBrokersRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBrokersRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBrokersRequest } as QueryBrokersRequest;
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
    const message = { ...baseQueryBrokersRequest } as QueryBrokersRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryBrokersRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBrokersRequest>): QueryBrokersRequest {
    const message = { ...baseQueryBrokersRequest } as QueryBrokersRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryBrokersResponse: object = {};

export const QueryBrokersResponse = {
  encode(
    message: QueryBrokersResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.brokers) {
      Broker.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBrokersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBrokersResponse } as QueryBrokersResponse;
    message.brokers = [];
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
    const message = { ...baseQueryBrokersResponse } as QueryBrokersResponse;
    message.brokers = [];
    if (object.brokers !== undefined && object.brokers !== null) {
      for (const e of object.brokers) {
        message.brokers.push(Broker.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryBrokersResponse): unknown {
    const obj: any = {};
    if (message.brokers) {
      obj.brokers = message.brokers.map((e) =>
        e ? Broker.toJSON(e) : undefined
      );
    } else {
      obj.brokers = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBrokersResponse>): QueryBrokersResponse {
    const message = { ...baseQueryBrokersResponse } as QueryBrokersResponse;
    message.brokers = [];
    if (object.brokers !== undefined && object.brokers !== null) {
      for (const e of object.brokers) {
        message.brokers.push(Broker.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** QueryInterchainAccountFromAddress returns the interchain account for given owner address on a given connection pair */
  InterchainAccountFromAddress(
    request: QueryInterchainAccountFromAddressRequest
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
  }
  InterchainAccountFromAddress(
    request: QueryInterchainAccountFromAddressRequest
  ): Promise<QueryInterchainAccountFromAddressResponse> {
    const data = QueryInterchainAccountFromAddressRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.broker.Query",
      "InterchainAccountFromAddress",
      data
    );
    return promise.then((data) =>
      QueryInterchainAccountFromAddressResponse.decode(new Reader(data))
    );
  }

  Broker(request: QueryBrokerRequest): Promise<QueryBrokerResponse> {
    const data = QueryBrokerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.broker.Query",
      "Broker",
      data
    );
    return promise.then((data) => QueryBrokerResponse.decode(new Reader(data)));
  }

  Brokers(request: QueryBrokersRequest): Promise<QueryBrokersResponse> {
    const data = QueryBrokersRequest.encode(request).finish();
    const promise = this.rpc.request(
      "defundlabs.defund.broker.Query",
      "Brokers",
      data
    );
    return promise.then((data) =>
      QueryBrokersResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
