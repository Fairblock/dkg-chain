package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTimeout = "timeout"

var _ sdk.Msg = &MsgTimeout{}

func NewMsgTimeout(creator string, round string, id string) *MsgTimeout {
	return &MsgTimeout{
		Creator: creator,
		Round:   round,
		Id:      id,
	}
}

func (msg *MsgTimeout) Route() string {
	return RouterKey
}

func (msg *MsgTimeout) Type() string {
	return TypeMsgTimeout
}

func (msg *MsgTimeout) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTimeout) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTimeout) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
