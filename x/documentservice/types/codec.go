package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateContract{}, "documentservice/CreateContract", nil)
	cdc.RegisterConcrete(&MsgCreateAnnex{}, "documentservice/CreateAnnex", nil)
	cdc.RegisterConcrete(&MsgSignAnnex{}, "documentservice/SignAnnex", nil)
	cdc.RegisterConcrete(&MsgSignContract{}, "documentservice/SignContract", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAnnex{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSignAnnex{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSignContract{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
