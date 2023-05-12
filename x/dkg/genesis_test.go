package dkg_test

import (
	"testing"

	keepertest "dkg/testutil/keeper"
	"dkg/testutil/nullify"
	"dkg/x/dkg"
	"dkg/x/dkg/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DkgKeeper(t)
	dkg.InitGenesis(ctx, *k, genesisState)
	got := dkg.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
