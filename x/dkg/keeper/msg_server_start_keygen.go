package keeper

import (
	"context"

	"dkg/x/dkg/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) StartKeygen(goCtx context.Context, msg *types.MsgStartKeygen) (*types.MsgStartKeygenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	k.InitCounter(ctx)
	_ = ctx
	event := sdk.NewEvent(
		types.EventTypeKeygen,
		sdk.NewAttribute(types.AttributeValueStart, msg.KeyID),
		sdk.NewAttribute("threshold", msg.Threshold),
		sdk.NewAttribute("participants", msg.Participants),
		sdk.NewAttribute("timeout", msg.Timeout),
		sdk.NewAttribute("module", "dkg"),
	)
	ctx.EventManager().EmitEvent(event)
	return &types.MsgStartKeygenResponse{}, nil
}
