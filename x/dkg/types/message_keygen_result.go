package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgKeygenResult = "keygen_result"

var _ sdk.Msg = &MsgKeygenResult{}

func NewMsgKeygenResult(creator string, myIndex string, commitment string) *MsgKeygenResult {
	return &MsgKeygenResult{
		Creator:    creator,
		MyIndex:    myIndex,
		Commitment: commitment,
	}
}

func (msg *MsgKeygenResult) Route() string {
	return RouterKey
}

func (msg *MsgKeygenResult) Type() string {
	return TypeMsgKeygenResult
}

func (msg *MsgKeygenResult) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgKeygenResult) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgKeygenResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
