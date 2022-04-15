/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "cosmonaut.documentservice.documentservice";

export interface Contract {
  id: number;
  contractHash: string;
  state: string;
  seller: string;
  buyer: string;
  createDate: string;
}

const baseContract: object = {
  id: 0,
  contractHash: "",
  state: "",
  seller: "",
  buyer: "",
  createDate: "",
};

export const Contract = {
  encode(message: Contract, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.contractHash !== "") {
      writer.uint32(18).string(message.contractHash);
    }
    if (message.state !== "") {
      writer.uint32(26).string(message.state);
    }
    if (message.seller !== "") {
      writer.uint32(34).string(message.seller);
    }
    if (message.buyer !== "") {
      writer.uint32(42).string(message.buyer);
    }
    if (message.createDate !== "") {
      writer.uint32(50).string(message.createDate);
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
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.contractHash = reader.string();
          break;
        case 3:
          message.state = reader.string();
          break;
        case 4:
          message.seller = reader.string();
          break;
        case 5:
          message.buyer = reader.string();
          break;
        case 6:
          message.createDate = reader.string();
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
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.contractHash !== undefined && object.contractHash !== null) {
      message.contractHash = String(object.contractHash);
    } else {
      message.contractHash = "";
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
    if (object.createDate !== undefined && object.createDate !== null) {
      message.createDate = String(object.createDate);
    } else {
      message.createDate = "";
    }
    return message;
  },

  toJSON(message: Contract): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.contractHash !== undefined &&
      (obj.contractHash = message.contractHash);
    message.state !== undefined && (obj.state = message.state);
    message.seller !== undefined && (obj.seller = message.seller);
    message.buyer !== undefined && (obj.buyer = message.buyer);
    message.createDate !== undefined && (obj.createDate = message.createDate);
    return obj;
  },

  fromPartial(object: DeepPartial<Contract>): Contract {
    const message = { ...baseContract } as Contract;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.contractHash !== undefined && object.contractHash !== null) {
      message.contractHash = object.contractHash;
    } else {
      message.contractHash = "";
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
    if (object.createDate !== undefined && object.createDate !== null) {
      message.createDate = object.createDate;
    } else {
      message.createDate = "";
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
