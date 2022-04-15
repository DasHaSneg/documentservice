/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../documentservice/params";
import { Contract } from "../documentservice/contract";

export const protobufPackage = "cosmonaut.documentservice.documentservice";

/** GenesisState defines the documentservice module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  contractList: Contract[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  contractCount: number;
}

const baseGenesisState: object = { contractCount: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.contractList) {
      Contract.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.contractCount !== 0) {
      writer.uint32(24).uint64(message.contractCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.contractList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.contractList.push(Contract.decode(reader, reader.uint32()));
          break;
        case 3:
          message.contractCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.contractList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.contractList !== undefined && object.contractList !== null) {
      for (const e of object.contractList) {
        message.contractList.push(Contract.fromJSON(e));
      }
    }
    if (object.contractCount !== undefined && object.contractCount !== null) {
      message.contractCount = Number(object.contractCount);
    } else {
      message.contractCount = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.contractList) {
      obj.contractList = message.contractList.map((e) =>
        e ? Contract.toJSON(e) : undefined
      );
    } else {
      obj.contractList = [];
    }
    message.contractCount !== undefined &&
      (obj.contractCount = message.contractCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.contractList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.contractList !== undefined && object.contractList !== null) {
      for (const e of object.contractList) {
        message.contractList.push(Contract.fromPartial(e));
      }
    }
    if (object.contractCount !== undefined && object.contractCount !== null) {
      message.contractCount = object.contractCount;
    } else {
      message.contractCount = 0;
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
