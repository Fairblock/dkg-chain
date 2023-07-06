package simulation

import (
	"math/rand"

	"dkg/x/dkg/keeper"
	"dkg/x/dkg/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgKeygenResult(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgKeygenResult{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the KeygenResult simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "KeygenResult simulation not implemented"), nil, nil
	}
}
