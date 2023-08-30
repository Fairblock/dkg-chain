package keeper

import (
	"context"
	//"os"
	"strconv"

	"dkg/x/dkg/types"

	//	"github.com/cometbft/cometbft/libs/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sirupsen/logrus"
	"google.golang.org/go/pkg/mod/github.com/sirupsen/logrus@v1.9.2"
	//	"github.com/sirupsen/logrus"
)

func (k msgServer) StartKeygen(goCtx context.Context, msg *types.MsgStartKeygen) (*types.MsgStartKeygenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.InitCounter(ctx)
	logrus.SetLevel(logrus.DebugLevel)
	timeout, err := strconv.ParseUint(msg.Timeout, 10, 64)
	if err != nil {
		logrus.Error("start keygen error:", err)
	}
	k.InitTimeout(ctx,0,timeout,uint64(ctx.BlockHeight()),msg.KeyID)
	k.InitMPK(ctx,msg.KeyID)
	
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
