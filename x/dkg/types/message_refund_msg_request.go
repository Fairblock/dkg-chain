package types

import (
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRefundMsgRequest = "refund_msg_request"

var _ sdk.Msg = &MsgRefundMsgRequest{}

func NewMsgRefundMsgRequest(creator string, sender github_com_cosmos_cosmos_sdk_types.AccAddress, innerMessage *types1.Any) *MsgRefundMsgRequest {
	return &MsgRefundMsgRequest{
		Creator:      creator,
		Sender:       sender,
		InnerMessage: innerMessage,
	}
}

func (msg *MsgRefundMsgRequest) Route() string {
	return RouterKey
}

func (msg *MsgRefundMsgRequest) Type() string {
	return TypeMsgRefundMsgRequest
}

func (msg *MsgRefundMsgRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRefundMsgRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRefundMsgRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
