/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

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

/** Query defines the gRPC querier service. */
export interface Query {
  /** QueryInterchainAccountFromAddress returns the interchain account for given owner address on a given connection pair */
  InterchainAccountFromAddress(
    request: QueryInterchainAccountFromAddressRequest
  ): Promise<QueryInterchainAccountFromAddressResponse>;
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
