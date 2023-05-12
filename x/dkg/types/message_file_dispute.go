package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFileDispute = "file_dispute"

var _ sdk.Msg = &MsgFileDispute{}

func NewMsgFileDispute(creator string) *MsgFileDispute {
	return &MsgFileDispute{
		Creator: creator,
	}
}

func (msg *MsgFileDispute) Route() string {
	return RouterKey
}

func (msg *MsgFileDispute) Type() string {
	return TypeMsgFileDispute
}

func (msg *MsgFileDispute) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFileDispute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFileDispute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
