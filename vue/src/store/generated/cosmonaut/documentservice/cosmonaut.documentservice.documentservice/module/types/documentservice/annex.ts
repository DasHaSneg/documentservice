/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "cosmonaut.documentservice.documentservice";

export interface Annex {
  creator: string;
  id: number;
  annexHash: string;
  contractId: string;
  state: string;
  seller: string;
  buyer: string;
  createDate: string;
}

const baseAnnex: object = {
  creator: "",
  id: 0,
  annexHash: "",
  contractId: "",
  state: "",
  seller: "",
  buyer: "",
  createDate: "",
};

export const Annex = {
  encode(message: Annex, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.annexHash !== "") {
      writer.uint32(26).string(message.annexHash);
    }
    if (message.contractId !== "") {
      writer.uint32(34).string(message.contractId);
    }
    if (message.state !== "") {
      writer.uint32(42).string(message.state);
    }
    if (message.seller !== "") {
      writer.uint32(50).string(message.seller);
    }
    if (message.buyer !== "") {
      writer.uint32(58).string(message.buyer);
    }
    if (message.createDate !== "") {
      writer.uint32(66).string(message.createDate);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Annex {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAnnex } as Annex;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.annexHash = reader.string();
          break;
        case 4:
          message.contractId = reader.string();
          break;
        case 5:
          message.state = reader.string();
          break;
        case 6:
          message.seller = reader.string();
          break;
        case 7:
          message.buyer = reader.string();
          break;
        case 8:
          message.createDate = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Annex {
    const message = { ...baseAnnex } as Annex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.annexHash !== undefined && object.annexHash !== null) {
      message.annexHash = String(object.annexHash);
    } else {
      message.annexHash = "";
    }
    if (object.contractId !== undefined && object.contractId !== null) {
      message.contractId = String(object.contractId);
    } else {
      message.contractId = "";
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

  toJSON(message: Annex): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.annexHash !== undefined && (obj.annexHash = message.annexHash);
    message.contractId !== undefined && (obj.contractId = message.contractId);
    message.state !== undefined && (obj.state = message.state);
    message.seller !== undefined && (obj.seller = message.seller);
    message.buyer !== undefined && (obj.buyer = message.buyer);
    message.createDate !== undefined && (obj.createDate = message.createDate);
    return obj;
  },

  fromPartial(object: DeepPartial<Annex>): Annex {
    const message = { ...baseAnnex } as Annex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.annexHash !== undefined && object.annexHash !== null) {
      message.annexHash = object.annexHash;
    } else {
      message.annexHash = "";
    }
    if (object.contractId !== undefined && object.contractId !== null) {
      message.contractId = object.contractId;
    } else {
      message.contractId = "";
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
