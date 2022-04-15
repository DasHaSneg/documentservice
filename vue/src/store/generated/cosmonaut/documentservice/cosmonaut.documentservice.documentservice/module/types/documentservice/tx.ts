/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "cosmonaut.documentservice.documentservice";

export interface MsgCreateContract {
  creator: string;
  contractHash: string;
  buyer: string;
}

export interface MsgCreateContractResponse {
  id: number;
  createDate: string;
}

export interface MsgCreateAnnex {
  creator: string;
  annexHash: string;
  contractId: number;
  buyer: string;
}

export interface MsgCreateAnnexResponse {
  id: number;
  createDate: string;
}

export interface MsgSignAnnex {
  creator: string;
  annexId: number;
}

export interface MsgSignAnnexResponse {}

const baseMsgCreateContract: object = {
  creator: "",
  contractHash: "",
  buyer: "",
};

export const MsgCreateContract = {
  encode(message: MsgCreateContract, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.contractHash !== "") {
      writer.uint32(18).string(message.contractHash);
    }
    if (message.buyer !== "") {
      writer.uint32(26).string(message.buyer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateContract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateContract } as MsgCreateContract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.contractHash = reader.string();
          break;
        case 3:
          message.buyer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateContract {
    const message = { ...baseMsgCreateContract } as MsgCreateContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.contractHash !== undefined && object.contractHash !== null) {
      message.contractHash = String(object.contractHash);
    } else {
      message.contractHash = "";
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = String(object.buyer);
    } else {
      message.buyer = "";
    }
    return message;
  },

  toJSON(message: MsgCreateContract): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.contractHash !== undefined &&
      (obj.contractHash = message.contractHash);
    message.buyer !== undefined && (obj.buyer = message.buyer);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateContract>): MsgCreateContract {
    const message = { ...baseMsgCreateContract } as MsgCreateContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.contractHash !== undefined && object.contractHash !== null) {
      message.contractHash = object.contractHash;
    } else {
      message.contractHash = "";
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = object.buyer;
    } else {
      message.buyer = "";
    }
    return message;
  },
};

const baseMsgCreateContractResponse: object = { id: 0, createDate: "" };

export const MsgCreateContractResponse = {
  encode(
    message: MsgCreateContractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.createDate !== "") {
      writer.uint32(18).string(message.createDate);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateContractResponse,
    } as MsgCreateContractResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.createDate = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateContractResponse {
    const message = {
      ...baseMsgCreateContractResponse,
    } as MsgCreateContractResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.createDate !== undefined && object.createDate !== null) {
      message.createDate = String(object.createDate);
    } else {
      message.createDate = "";
    }
    return message;
  },

  toJSON(message: MsgCreateContractResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.createDate !== undefined && (obj.createDate = message.createDate);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateContractResponse>
  ): MsgCreateContractResponse {
    const message = {
      ...baseMsgCreateContractResponse,
    } as MsgCreateContractResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.createDate !== undefined && object.createDate !== null) {
      message.createDate = object.createDate;
    } else {
      message.createDate = "";
    }
    return message;
  },
};

const baseMsgCreateAnnex: object = {
  creator: "",
  annexHash: "",
  contractId: 0,
  buyer: "",
};

export const MsgCreateAnnex = {
  encode(message: MsgCreateAnnex, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.annexHash !== "") {
      writer.uint32(18).string(message.annexHash);
    }
    if (message.contractId !== 0) {
      writer.uint32(24).uint64(message.contractId);
    }
    if (message.buyer !== "") {
      writer.uint32(34).string(message.buyer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateAnnex {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateAnnex } as MsgCreateAnnex;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.annexHash = reader.string();
          break;
        case 3:
          message.contractId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.buyer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAnnex {
    const message = { ...baseMsgCreateAnnex } as MsgCreateAnnex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.annexHash !== undefined && object.annexHash !== null) {
      message.annexHash = String(object.annexHash);
    } else {
      message.annexHash = "";
    }
    if (object.contractId !== undefined && object.contractId !== null) {
      message.contractId = Number(object.contractId);
    } else {
      message.contractId = 0;
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = String(object.buyer);
    } else {
      message.buyer = "";
    }
    return message;
  },

  toJSON(message: MsgCreateAnnex): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.annexHash !== undefined && (obj.annexHash = message.annexHash);
    message.contractId !== undefined && (obj.contractId = message.contractId);
    message.buyer !== undefined && (obj.buyer = message.buyer);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateAnnex>): MsgCreateAnnex {
    const message = { ...baseMsgCreateAnnex } as MsgCreateAnnex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.annexHash !== undefined && object.annexHash !== null) {
      message.annexHash = object.annexHash;
    } else {
      message.annexHash = "";
    }
    if (object.contractId !== undefined && object.contractId !== null) {
      message.contractId = object.contractId;
    } else {
      message.contractId = 0;
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = object.buyer;
    } else {
      message.buyer = "";
    }
    return message;
  },
};

const baseMsgCreateAnnexResponse: object = { id: 0, createDate: "" };

export const MsgCreateAnnexResponse = {
  encode(
    message: MsgCreateAnnexResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.createDate !== "") {
      writer.uint32(18).string(message.createDate);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateAnnexResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateAnnexResponse } as MsgCreateAnnexResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.createDate = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAnnexResponse {
    const message = { ...baseMsgCreateAnnexResponse } as MsgCreateAnnexResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.createDate !== undefined && object.createDate !== null) {
      message.createDate = String(object.createDate);
    } else {
      message.createDate = "";
    }
    return message;
  },

  toJSON(message: MsgCreateAnnexResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.createDate !== undefined && (obj.createDate = message.createDate);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateAnnexResponse>
  ): MsgCreateAnnexResponse {
    const message = { ...baseMsgCreateAnnexResponse } as MsgCreateAnnexResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.createDate !== undefined && object.createDate !== null) {
      message.createDate = object.createDate;
    } else {
      message.createDate = "";
    }
    return message;
  },
};

const baseMsgSignAnnex: object = { creator: "", annexId: 0 };

export const MsgSignAnnex = {
  encode(message: MsgSignAnnex, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.annexId !== 0) {
      writer.uint32(16).uint64(message.annexId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSignAnnex {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSignAnnex } as MsgSignAnnex;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.annexId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSignAnnex {
    const message = { ...baseMsgSignAnnex } as MsgSignAnnex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.annexId !== undefined && object.annexId !== null) {
      message.annexId = Number(object.annexId);
    } else {
      message.annexId = 0;
    }
    return message;
  },

  toJSON(message: MsgSignAnnex): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.annexId !== undefined && (obj.annexId = message.annexId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSignAnnex>): MsgSignAnnex {
    const message = { ...baseMsgSignAnnex } as MsgSignAnnex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.annexId !== undefined && object.annexId !== null) {
      message.annexId = object.annexId;
    } else {
      message.annexId = 0;
    }
    return message;
  },
};

const baseMsgSignAnnexResponse: object = {};

export const MsgSignAnnexResponse = {
  encode(_: MsgSignAnnexResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSignAnnexResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSignAnnexResponse } as MsgSignAnnexResponse;
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

  fromJSON(_: any): MsgSignAnnexResponse {
    const message = { ...baseMsgSignAnnexResponse } as MsgSignAnnexResponse;
    return message;
  },

  toJSON(_: MsgSignAnnexResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSignAnnexResponse>): MsgSignAnnexResponse {
    const message = { ...baseMsgSignAnnexResponse } as MsgSignAnnexResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateContract(
    request: MsgCreateContract
  ): Promise<MsgCreateContractResponse>;
  CreateAnnex(request: MsgCreateAnnex): Promise<MsgCreateAnnexResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SignAnnex(request: MsgSignAnnex): Promise<MsgSignAnnexResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateContract(
    request: MsgCreateContract
  ): Promise<MsgCreateContractResponse> {
    const data = MsgCreateContract.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.documentservice.documentservice.Msg",
      "CreateContract",
      data
    );
    return promise.then((data) =>
      MsgCreateContractResponse.decode(new Reader(data))
    );
  }

  CreateAnnex(request: MsgCreateAnnex): Promise<MsgCreateAnnexResponse> {
    const data = MsgCreateAnnex.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.documentservice.documentservice.Msg",
      "CreateAnnex",
      data
    );
    return promise.then((data) =>
      MsgCreateAnnexResponse.decode(new Reader(data))
    );
  }

  SignAnnex(request: MsgSignAnnex): Promise<MsgSignAnnexResponse> {
    const data = MsgSignAnnex.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.documentservice.documentservice.Msg",
      "SignAnnex",
      data
    );
    return promise.then((data) =>
      MsgSignAnnexResponse.decode(new Reader(data))
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
