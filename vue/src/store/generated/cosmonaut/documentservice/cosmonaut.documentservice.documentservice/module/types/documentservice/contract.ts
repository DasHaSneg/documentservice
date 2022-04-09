/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "cosmonaut.documentservice.documentservice";

export interface Contract {
  creator: string;
  Id: number;
  contract_hash: string;
  state: string;
  seller: string;
  buyer: string;
  create_date: string;
  createdAt: number;
}

const baseContract: object = {
  creator: "",
  Id: 0,
  contract_hash: "",
  state: "",
  seller: "",
  buyer: "",
  create_date: "",
  createdAt: 0,
};

export const Contract = {
  encode(message: Contract, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.Id !== 0) {
      writer.uint32(16).uint64(message.Id);
    }
    if (message.contract_hash !== "") {
      writer.uint32(26).string(message.contract_hash);
    }
    if (message.state !== "") {
      writer.uint32(34).string(message.state);
    }
    if (message.seller !== "") {
      writer.uint32(42).string(message.seller);
    }
    if (message.buyer !== "") {
      writer.uint32(50).string(message.buyer);
    }
    if (message.create_date !== "") {
      writer.uint32(58).string(message.create_date);
    }
    if (message.createdAt !== 0) {
      writer.uint32(64).int64(message.createdAt);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Contract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseContract } as Contract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.Id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.contract_hash = reader.string();
          break;
        case 4:
          message.state = reader.string();
          break;
        case 5:
          message.seller = reader.string();
          break;
        case 6:
          message.buyer = reader.string();
          break;
        case 7:
          message.create_date = reader.string();
          break;
        case 8:
          message.createdAt = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Contract {
    const message = { ...baseContract } as Contract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.Id !== undefined && object.Id !== null) {
      message.Id = Number(object.Id);
    } else {
      message.Id = 0;
    }
    if (object.contract_hash !== undefined && object.contract_hash !== null) {
      message.contract_hash = String(object.contract_hash);
    } else {
      message.contract_hash = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = String(object.state);
    } else {
      message.state = "";
    }
    if (object.seller !== undefined && object.seller !== null) {
      message.seller = String(object.seller);
    } else {
      message.seller = "";
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = String(object.buyer);
    } else {
      message.buyer = "";
    }
    if (object.create_date !== undefined && object.create_date !== null) {
      message.create_date = String(object.create_date);
    } else {
      message.create_date = "";
    }
    if (object.createdAt !== undefined && object.createdAt !== null) {
      message.createdAt = Number(object.createdAt);
    } else {
      message.createdAt = 0;
    }
    return message;
  },

  toJSON(message: Contract): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.Id !== undefined && (obj.Id = message.Id);
    message.contract_hash !== undefined &&
      (obj.contract_hash = message.contract_hash);
    message.state !== undefined && (obj.state = message.state);
    message.seller !== undefined && (obj.seller = message.seller);
    message.buyer !== undefined && (obj.buyer = message.buyer);
    message.create_date !== undefined &&
      (obj.create_date = message.create_date);
    message.createdAt !== undefined && (obj.createdAt = message.createdAt);
    return obj;
  },

  fromPartial(object: DeepPartial<Contract>): Contract {
    const message = { ...baseContract } as Contract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.Id !== undefined && object.Id !== null) {
      message.Id = object.Id;
    } else {
      message.Id = 0;
    }
    if (object.contract_hash !== undefined && object.contract_hash !== null) {
      message.contract_hash = object.contract_hash;
    } else {
      message.contract_hash = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = object.state;
    } else {
      message.state = "";
    }
    if (object.seller !== undefined && object.seller !== null) {
      message.seller = object.seller;
    } else {
      message.seller = "";
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = object.buyer;
    } else {
      message.buyer = "";
    }
    if (object.create_date !== undefined && object.create_date !== null) {
      message.create_date = object.create_date;
    } else {
      message.create_date = "";
    }
    if (object.createdAt !== undefined && object.createdAt !== null) {
      message.createdAt = object.createdAt;
    } else {
      message.createdAt = 0;
    }
    return message;
  },
};

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
