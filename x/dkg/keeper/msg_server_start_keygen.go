package keeper

import (
	"context"
	//"os"
	"strconv"

	"dkg/x/dkg/types"

	//	"github.com/cometbft/cometbft/libs/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sirupsen/logrus"
	//	"github.com/sirupsen/logrus"
)

func (k msgServer) StartKeygen(goCtx context.Context, msg *types.MsgStartKeygen) (*types.MsgStartKeygenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	//ctx.Logger().Error("Debug: ", "start")
	//logrus.Panic("Debug: ", "start")

	// TODO: Handling the message
	k.InitCounter(ctx)
	logrus.SetLevel(logrus.DebugLevel)
	timeout, _ := strconv.ParseUint(msg.Timeout, 10, 64)
	// goCtx = context.WithValue(goCtx, "timeout", timeout)
	// goCtx = context.WithValue(goCtx, "round", 0)
	// goCtx = context.WithValue(goCtx, "start", goCtx.Value(ctx.BlockHeight()))
	// goCtx = context.WithValue(goCtx, "id", msg.KeyID)
	k.InitTimeout(ctx,0,timeout,uint64(ctx.BlockHeight()),msg.KeyID)
	//logrus.Debug("++++++++++++++++++++++++++++++++++++++++++++",goCtx.Value("id"))	
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
