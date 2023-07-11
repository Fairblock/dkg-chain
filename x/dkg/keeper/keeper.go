package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"dkg/x/dkg/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) InitCounter(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	counter := types.Counter{Count: 0}
	store.Set([]byte("counter"), counter.MustMarshalBinaryBare())
}

func (k Keeper) IncreaseCounter(ctx sdk.Context, amount uint64) uint64 {
	store := ctx.KVStore(k.storeKey)
	var counter types.Counter
	bz := store.Get([]byte("counter"))
	counter.MustUnmarshalBinaryBare(bz)
	counter.Count += amount
	store.Set([]byte("counter"), counter.MustMarshalBinaryBare())
	return counter.Count
}

type HandleMsgInitCounter struct {
	// Add necessary fields, if any
}

func (k Keeper) InitTimeout(ctx sdk.Context, round uint64, timeout uint64, start uint64, id string) {
	store := ctx.KVStore(k.storeKey)
	timeoutData := types.TimeoutData{Round: round,Start: start,Timeout: timeout,Id: id}
	store.Set([]byte("timeoutData"), timeoutData.MustMarshalBinaryBare())
}

func (k Keeper) GetTimeout(ctx sdk.Context) types.TimeoutData{
	store := ctx.KVStore(k.storeKey)
	var timeoutData types.TimeoutData
	bz := store.Get([]byte("timeoutData"))
	if bz == nil{
		return types.TimeoutData{}
	}
	timeoutData.MustUnmarshalBinaryBare(bz)
	return timeoutData
}

func (k Keeper) NextRound(ctx sdk.Context){
	store := ctx.KVStore(k.storeKey)
	var timeoutData types.TimeoutData
	bz := store.Get([]byte("timeoutData"))
	timeoutData.MustUnmarshalBinaryBare(bz)
	timeoutData.Round = timeoutData.Round + 1
	store.Set([]byte("timeoutData"), timeoutData.MustMarshalBinaryBare())

}

// func (msg HandleMsgInitCounter) HandleMsg(ctx sdk.Context, k CounterKeeper) sdk.Result {
//     k.InitCounter(ctx)
//     return sdk.Result{Events: ctx.EventManager().ABCIEvents()}
// }

// type HandleMsgIncreaseCounter struct {
//     Amount uint64 `json:"amount"`
// }

// func (msg HandleMsgIncreaseCounter) HandleMsg(ctx sdk.Context, k CounterKeeper) sdk.Result {
//     k.IncreaseCounter(ctx, msg.Amount)
//     return sdk.Result{Events: ctx.EventManager().ABCIEvents()}
// }
