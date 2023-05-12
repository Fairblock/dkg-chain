package keeper

import (
	"dkg/x/dkg/types"
)

var _ types.QueryServer = Keeper{}
