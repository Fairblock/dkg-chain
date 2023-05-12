package keeper_test

import (
	"context"
	"testing"

	keepertest "dkg/testutil/keeper"
	"dkg/x/dkg/keeper"
	"dkg/x/dkg/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DkgKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
