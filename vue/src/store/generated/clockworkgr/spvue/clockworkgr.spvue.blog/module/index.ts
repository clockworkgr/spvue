import {  StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreatePost } from "./types/blog/tx";
import { MsgDeletePost } from "./types/blog/tx";
import { MsgUpdatePost } from "./types/blog/tx";


const types = [
  ["/clockworkgr.spvue.blog.MsgCreatePost", MsgCreatePost],
  ["/clockworkgr.spvue.blog.MsgDeletePost", MsgDeletePost],
  ["/clockworkgr.spvue.blog.MsgUpdatePost", MsgUpdatePost],
  
];

const registry = new Registry(<any>types);

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
  if (!wallet) throw new Error("wallet is required");

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee=defaultFee, memo=null }: SignAndBroadcastOptions) => memo?client.signAndBroadcast(address, msgs, fee,memo):client.signAndBroadcast(address, msgs, fee),
    msgCreatePost: (data: MsgCreatePost): EncodeObject => ({ typeUrl: "/clockworkgr.spvue.blog.MsgCreatePost", value: data }),
    msgDeletePost: (data: MsgDeletePost): EncodeObject => ({ typeUrl: "/clockworkgr.spvue.blog.MsgDeletePost", value: data }),
    msgUpdatePost: (data: MsgUpdatePost): EncodeObject => ({ typeUrl: "/clockworkgr.spvue.blog.MsgUpdatePost", value: data }),
    
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
