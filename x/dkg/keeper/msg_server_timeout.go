package keeper

import (
	"context"

	"dkg/x/dkg/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Timeout(goCtx context.Context, msg *types.MsgTimeout) (*types.MsgTimeoutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	event := sdk.NewEvent(
		types.EventTypeKeygen,
		sdk.NewAttribute("dkg-id", msg.Id),
		sdk.NewAttribute("timeout-round", msg.Round),
		sdk.NewAttribute("creator", msg.Creator),
	)
	ctx.EventManager().EmitEvent(event)
	return &types.MsgTimeoutResponse{}, nil
}
