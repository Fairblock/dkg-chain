package keeper

import (
	"context"

	"dkg/x/dkg/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) KeygenResult(goCtx context.Context, msg *types.MsgKeygenResult) (*types.MsgKeygenResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	
	_ = ctx
	event := sdk.NewEvent(
		"dkg-result",
		sdk.NewAttribute("commitment", msg.Commitment),
		sdk.NewAttribute("myIndex", msg.MyIndex),
		
	)
	ctx.EventManager().EmitEvent(event)
	return &types.MsgKeygenResultResponse{}, nil
}
