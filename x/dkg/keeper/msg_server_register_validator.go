package keeper

import (
	"context"
	"strconv"

	"dkg/x/dkg/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sirupsen/logrus"
)

func (k msgServer) RegisterValidator(goCtx context.Context, msg *types.MsgRegisterValidator) (*types.MsgRegisterValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// TODO: Handling the message
	_ = ctx
	if msg.Participation {
		k.AddAddress(ctx,msg.Creator)
	}
	if !msg.Participation {
		k.RemoveAddress(ctx,msg.Creator)
	}
	logrus.Info("registeration: ***********************************", msg.Creator, msg.Participation)
	event := sdk.NewEvent(
		"dkg-registeration",
		sdk.NewAttribute("address", msg.Creator),
		sdk.NewAttribute("participation", strconv.FormatBool(msg.Participation)),
	)
	ctx.EventManager().EmitEvent(event)
	return &types.MsgRegisterValidatorResponse{}, nil
}
