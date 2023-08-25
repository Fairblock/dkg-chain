package keeper

import (
	"context"

	"dkg/x/dkg/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) KeygenResult(goCtx context.Context, msg *types.MsgKeygenResult) (*types.MsgKeygenResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	
	_ = ctx

	return &types.MsgKeygenResultResponse{}, nil
}
