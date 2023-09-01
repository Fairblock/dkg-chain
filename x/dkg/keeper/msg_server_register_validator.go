package keeper

import (
	"context"
	"strconv"

	"dkg/x/dkg/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterValidator(goCtx context.Context, msg *types.MsgRegisterValidator) (*types.MsgRegisterValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	if msg.Participation {
		k.AddAddress(ctx,msg.Address)
	}
	if !msg.Participation {
		k.RemoveAddress(ctx,msg.Address)
	}
	event := sdk.NewEvent(
		"dkg-registeration",
		sdk.NewAttribute("address", msg.Address),
		sdk.NewAttribute("participation", strconv.FormatBool(msg.Participation)),
	)
	ctx.EventManager().EmitEvent(event)
	return &types.MsgRegisterValidatorResponse{}, nil
}
