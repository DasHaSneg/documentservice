// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSignContract } from "./types/documentservice/tx";
import { MsgSignAnnex } from "./types/documentservice/tx";
import { MsgCompleteContract } from "./types/documentservice/tx";
import { MsgCreateContract } from "./types/documentservice/tx";
import { MsgCreateAnnex } from "./types/documentservice/tx";


const types = [
  ["/cosmonaut.documentservice.documentservice.MsgSignContract", MsgSignContract],
  ["/cosmonaut.documentservice.documentservice.MsgSignAnnex", MsgSignAnnex],
  ["/cosmonaut.documentservice.documentservice.MsgCompleteContract", MsgCompleteContract],
  ["/cosmonaut.documentservice.documentservice.MsgCreateContract", MsgCreateContract],
  ["/cosmonaut.documentservice.documentservice.MsgCreateAnnex", MsgCreateAnnex],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgSignContract: (data: MsgSignContract): EncodeObject => ({ typeUrl: "/cosmonaut.documentservice.documentservice.MsgSignContract", value: MsgSignContract.fromPartial( data ) }),
    msgSignAnnex: (data: MsgSignAnnex): EncodeObject => ({ typeUrl: "/cosmonaut.documentservice.documentservice.MsgSignAnnex", value: MsgSignAnnex.fromPartial( data ) }),
    msgCompleteContract: (data: MsgCompleteContract): EncodeObject => ({ typeUrl: "/cosmonaut.documentservice.documentservice.MsgCompleteContract", value: MsgCompleteContract.fromPartial( data ) }),
    msgCreateContract: (data: MsgCreateContract): EncodeObject => ({ typeUrl: "/cosmonaut.documentservice.documentservice.MsgCreateContract", value: MsgCreateContract.fromPartial( data ) }),
    msgCreateAnnex: (data: MsgCreateAnnex): EncodeObject => ({ typeUrl: "/cosmonaut.documentservice.documentservice.MsgCreateAnnex", value: MsgCreateAnnex.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
