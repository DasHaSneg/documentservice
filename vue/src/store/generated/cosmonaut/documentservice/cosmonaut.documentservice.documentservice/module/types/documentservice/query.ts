/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../documentservice/params";
import { Contract } from "../documentservice/contract";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Annex } from "../documentservice/annex";

export const protobufPackage = "cosmonaut.documentservice.documentservice";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetContractRequest {
  id: number;
}

export interface QueryGetContractResponse {
  Contract: Contract | undefined;
}

export interface QueryAllContractRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllContractResponse {
  Contract: Contract[];
  pagination: PageResponse | undefined;
}

export interface QueryGetAnnexRequest {
  id: number;
}

export interface QueryGetAnnexResponse {
  Annex: Annex | undefined;
}

export interface QueryAllAnnexRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllAnnexResponse {
  Annex: Annex[];
  pagination: PageResponse | undefined;
}

export interface QueryContractsByInnRequest {
  inn: string;
  pagination: PageRequest | undefined;
}

export interface QueryContractsByInnResponse {
  Contract: Contract[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetContractRequest: object = { id: 0 };

export const QueryGetContractRequest = {
  encode(
    message: QueryGetContractRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetContractRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetContractRequest,
    } as QueryGetContractRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetContractRequest {
    const message = {
      ...baseQueryGetContractRequest,
    } as QueryGetContractRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetContractRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetContractRequest>
  ): QueryGetContractRequest {
    const message = {
      ...baseQueryGetContractRequest,
    } as QueryGetContractRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetContractResponse: object = {};

export const QueryGetContractResponse = {
  encode(
    message: QueryGetContractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Contract !== undefined) {
      Contract.encode(message.Contract, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetContractResponse,
    } as QueryGetContractResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Contract = Contract.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetContractResponse {
    const message = {
      ...baseQueryGetContractResponse,
    } as QueryGetContractResponse;
    if (object.Contract !== undefined && object.Contract !== null) {
      message.Contract = Contract.fromJSON(object.Contract);
    } else {
      message.Contract = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetContractResponse): unknown {
    const obj: any = {};
    message.Contract !== undefined &&
      (obj.Contract = message.Contract
        ? Contract.toJSON(message.Contract)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetContractResponse>
  ): QueryGetContractResponse {
    const message = {
      ...baseQueryGetContractResponse,
    } as QueryGetContractResponse;
    if (object.Contract !== undefined && object.Contract !== null) {
      message.Contract = Contract.fromPartial(object.Contract);
    } else {
      message.Contract = undefined;
    }
    return message;
  },
};

const baseQueryAllContractRequest: object = {};

export const QueryAllContractRequest = {
  encode(
    message: QueryAllContractRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllContractRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllContractRequest,
    } as QueryAllContractRequest;
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

  fromJSON(object: any): QueryAllContractRequest {
    const message = {
      ...baseQueryAllContractRequest,
    } as QueryAllContractRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllContractRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllContractRequest>
  ): QueryAllContractRequest {
    const message = {
      ...baseQueryAllContractRequest,
    } as QueryAllContractRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllContractResponse: object = {};

export const QueryAllContractResponse = {
  encode(
    message: QueryAllContractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Contract) {
      Contract.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllContractResponse,
    } as QueryAllContractResponse;
    message.Contract = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Contract.push(Contract.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllContractResponse {
    const message = {
      ...baseQueryAllContractResponse,
    } as QueryAllContractResponse;
    message.Contract = [];
    if (object.Contract !== undefined && object.Contract !== null) {
      for (const e of object.Contract) {
        message.Contract.push(Contract.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllContractResponse): unknown {
    const obj: any = {};
    if (message.Contract) {
      obj.Contract = message.Contract.map((e) =>
        e ? Contract.toJSON(e) : undefined
      );
    } else {
      obj.Contract = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllContractResponse>
  ): QueryAllContractResponse {
    const message = {
      ...baseQueryAllContractResponse,
    } as QueryAllContractResponse;
    message.Contract = [];
    if (object.Contract !== undefined && object.Contract !== null) {
      for (const e of object.Contract) {
        message.Contract.push(Contract.fromPartial(e));
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

const baseQueryGetAnnexRequest: object = { id: 0 };

export const QueryGetAnnexRequest = {
  encode(
    message: QueryGetAnnexRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetAnnexRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetAnnexRequest } as QueryGetAnnexRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAnnexRequest {
    const message = { ...baseQueryGetAnnexRequest } as QueryGetAnnexRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetAnnexRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetAnnexRequest>): QueryGetAnnexRequest {
    const message = { ...baseQueryGetAnnexRequest } as QueryGetAnnexRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetAnnexResponse: object = {};

export const QueryGetAnnexResponse = {
  encode(
    message: QueryGetAnnexResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Annex !== undefined) {
      Annex.encode(message.Annex, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetAnnexResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetAnnexResponse } as QueryGetAnnexResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Annex = Annex.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAnnexResponse {
    const message = { ...baseQueryGetAnnexResponse } as QueryGetAnnexResponse;
    if (object.Annex !== undefined && object.Annex !== null) {
      message.Annex = Annex.fromJSON(object.Annex);
    } else {
      message.Annex = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetAnnexResponse): unknown {
    const obj: any = {};
    message.Annex !== undefined &&
      (obj.Annex = message.Annex ? Annex.toJSON(message.Annex) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAnnexResponse>
  ): QueryGetAnnexResponse {
    const message = { ...baseQueryGetAnnexResponse } as QueryGetAnnexResponse;
    if (object.Annex !== undefined && object.Annex !== null) {
      message.Annex = Annex.fromPartial(object.Annex);
    } else {
      message.Annex = undefined;
    }
    return message;
  },
};

const baseQueryAllAnnexRequest: object = {};

export const QueryAllAnnexRequest = {
  encode(
    message: QueryAllAnnexRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllAnnexRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllAnnexRequest } as QueryAllAnnexRequest;
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

  fromJSON(object: any): QueryAllAnnexRequest {
    const message = { ...baseQueryAllAnnexRequest } as QueryAllAnnexRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAnnexRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllAnnexRequest>): QueryAllAnnexRequest {
    const message = { ...baseQueryAllAnnexRequest } as QueryAllAnnexRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllAnnexResponse: object = {};

export const QueryAllAnnexResponse = {
  encode(
    message: QueryAllAnnexResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Annex) {
      Annex.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllAnnexResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllAnnexResponse } as QueryAllAnnexResponse;
    message.Annex = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Annex.push(Annex.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllAnnexResponse {
    const message = { ...baseQueryAllAnnexResponse } as QueryAllAnnexResponse;
    message.Annex = [];
    if (object.Annex !== undefined && object.Annex !== null) {
      for (const e of object.Annex) {
        message.Annex.push(Annex.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAnnexResponse): unknown {
    const obj: any = {};
    if (message.Annex) {
      obj.Annex = message.Annex.map((e) => (e ? Annex.toJSON(e) : undefined));
    } else {
      obj.Annex = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAnnexResponse>
  ): QueryAllAnnexResponse {
    const message = { ...baseQueryAllAnnexResponse } as QueryAllAnnexResponse;
    message.Annex = [];
    if (object.Annex !== undefined && object.Annex !== null) {
      for (const e of object.Annex) {
        message.Annex.push(Annex.fromPartial(e));
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

const baseQueryContractsByInnRequest: object = { inn: "" };

export const QueryContractsByInnRequest = {
  encode(
    message: QueryContractsByInnRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.inn !== "") {
      writer.uint32(10).string(message.inn);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryContractsByInnRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryContractsByInnRequest,
    } as QueryContractsByInnRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.inn = reader.string();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryContractsByInnRequest {
    const message = {
      ...baseQueryContractsByInnRequest,
    } as QueryContractsByInnRequest;
    if (object.inn !== undefined && object.inn !== null) {
      message.inn = String(object.inn);
    } else {
      message.inn = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryContractsByInnRequest): unknown {
    const obj: any = {};
    message.inn !== undefined && (obj.inn = message.inn);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryContractsByInnRequest>
  ): QueryContractsByInnRequest {
    const message = {
      ...baseQueryContractsByInnRequest,
    } as QueryContractsByInnRequest;
    if (object.inn !== undefined && object.inn !== null) {
      message.inn = object.inn;
    } else {
      message.inn = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryContractsByInnResponse: object = {};

export const QueryContractsByInnResponse = {
  encode(
    message: QueryContractsByInnResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Contract) {
      Contract.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryContractsByInnResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryContractsByInnResponse,
    } as QueryContractsByInnResponse;
    message.Contract = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Contract.push(Contract.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryContractsByInnResponse {
    const message = {
      ...baseQueryContractsByInnResponse,
    } as QueryContractsByInnResponse;
    message.Contract = [];
    if (object.Contract !== undefined && object.Contract !== null) {
      for (const e of object.Contract) {
        message.Contract.push(Contract.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryContractsByInnResponse): unknown {
    const obj: any = {};
    if (message.Contract) {
      obj.Contract = message.Contract.map((e) =>
        e ? Contract.toJSON(e) : undefined
      );
    } else {
      obj.Contract = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryContractsByInnResponse>
  ): QueryContractsByInnResponse {
    const message = {
      ...baseQueryContractsByInnResponse,
    } as QueryContractsByInnResponse;
    message.Contract = [];
    if (object.Contract !== undefined && object.Contract !== null) {
      for (const e of object.Contract) {
        message.Contract.push(Contract.fromPartial(e));
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
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Contract by id. */
  Contract(request: QueryGetContractRequest): Promise<QueryGetContractResponse>;
  /** Queries a list of Contract items. */
  ContractAll(
    request: QueryAllContractRequest
  ): Promise<QueryAllContractResponse>;
  /** Queries a Annex by id. */
  Annex(request: QueryGetAnnexRequest): Promise<QueryGetAnnexResponse>;
  /** Queries a list of Annex items. */
  AnnexAll(request: QueryAllAnnexRequest): Promise<QueryAllAnnexResponse>;
  /** Queries a list of ContractsByInn items. */
  ContractsByInn(
    request: QueryContractsByInnRequest
  ): Promise<QueryContractsByInnResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.documentservice.documentservice.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Contract(
    request: QueryGetContractRequest
  ): Promise<QueryGetContractResponse> {
    const data = QueryGetContractRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.documentservice.documentservice.Query",
      "Contract",
      data
    );
    return promise.then((data) =>
      QueryGetContractResponse.decode(new Reader(data))
    );
  }

  ContractAll(
    request: QueryAllContractRequest
  ): Promise<QueryAllContractResponse> {
    const data = QueryAllContractRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.documentservice.documentservice.Query",
      "ContractAll",
      data
    );
    return promise.then((data) =>
      QueryAllContractResponse.decode(new Reader(data))
    );
  }

  Annex(request: QueryGetAnnexRequest): Promise<QueryGetAnnexResponse> {
    const data = QueryGetAnnexRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.documentservice.documentservice.Query",
      "Annex",
      data
    );
    return promise.then((data) =>
      QueryGetAnnexResponse.decode(new Reader(data))
    );
  }

  AnnexAll(request: QueryAllAnnexRequest): Promise<QueryAllAnnexResponse> {
    const data = QueryAllAnnexRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.documentservice.documentservice.Query",
      "AnnexAll",
      data
    );
    return promise.then((data) =>
      QueryAllAnnexResponse.decode(new Reader(data))
    );
  }

  ContractsByInn(
    request: QueryContractsByInnRequest
  ): Promise<QueryContractsByInnResponse> {
    const data = QueryContractsByInnRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.documentservice.documentservice.Query",
      "ContractsByInn",
      data
    );
    return promise.then((data) =>
      QueryContractsByInnResponse.decode(new Reader(data))
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

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
