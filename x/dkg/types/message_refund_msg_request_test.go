package types

import (
	"testing"

	"dkg/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgRefundMsgRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRefundMsgRequest
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRefundMsgRequest{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRefundMsgRequest{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			//panic(tt.msg.InnerMessage)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
