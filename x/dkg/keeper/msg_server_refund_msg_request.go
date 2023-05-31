package keeper

import (
	"context"

	"dkg/x/dkg/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RefundMsgRequest(goCtx context.Context, msg *types.MsgRefundMsgRequest) (*types.MsgRefundMsgRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	//s := string(msg.Sender)
	b,_ :=msg.Marshal()

//panic(b)
	// TODO: Handling the message
	_ = ctx
	event := sdk.NewEvent(
		types.EventTypeKeygen,
		sdk.NewAttribute(types.AttributeValueMsg,string(b)),
		sdk.NewAttribute("module", "dkg"),
	)
	ctx.EventManager().EmitEvent(event)
	return &types.MsgRefundMsgRequestResponse{}, nil
}
