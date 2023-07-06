package keeper

import (
	"context"
	"strconv"

	"dkg/x/dkg/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RefundMsgRequest(goCtx context.Context, msg *types.MsgRefundMsgRequest) (*types.MsgRefundMsgRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	b, _ := msg.Marshal()

	count := k.IncreaseCounter(ctx, 1)
	str_count := strconv.FormatUint(count, 10)

	_ = ctx
	event := sdk.NewEvent(
		types.EventTypeKeygen,
		sdk.NewAttribute(types.AttributeValueMsg, string(b)),
		sdk.NewAttribute("module", "dkg"),
		sdk.NewAttribute("index", str_count),
	)
	ctx.EventManager().EmitEvent(event)
	return &types.MsgRefundMsgRequestResponse{}, nil
}
