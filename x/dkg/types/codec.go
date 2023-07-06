package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRefundMsgRequest{}, "dkg/RefundMsgRequest", nil)
	cdc.RegisterConcrete(&MsgFileDispute{}, "dkg/FileDispute", nil)
	cdc.RegisterConcrete(&MsgStartKeygen{}, "dkg/StartKeygen", nil)
	cdc.RegisterConcrete(&MsgKeygenResult{}, "dkg/KeygenResult", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRefundMsgRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFileDispute{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStartKeygen{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgKeygenResult{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
