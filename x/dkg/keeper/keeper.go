package keeper

import (
	"fmt"
	"sync"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/sirupsen/logrus"

	//	"github.com/sirupsen/logrus"

	//"github.com/sirupsen/logrus"
	"github.com/tendermint/tendermint/libs/log"

	"dkg/x/dkg/types"
)

type (
	Keeper struct {
		mu sync.Mutex
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
	k.mu.Lock()
    defer k.mu.Unlock()
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

func (k Keeper) InitMPK(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	pks := make(map[uint64][]byte)
	mpkData := types.MPKData{Pks: pks,Id: id}
	store.Set([]byte("mpkData"), mpkData.MustMarshalBinaryBare())
	store.Set([]byte("mpkData2"), mpkData.MustMarshalBinaryBare())
	store.Set([]byte("mpkData3"), mpkData.MustMarshalBinaryBare())
	k.InitializeList(ctx)
}

// func (k Keeper) AddFaulter(ctx sdk.Context, faulterId uint64, dkgId string){

	// k.mu.Lock()
    // defer k.mu.Unlock()
	// store := ctx.KVStore(k.storeKey)
	// var mpkData types.MPKData
	// var mpkData2 types.MPKData
	// var mpkData3 types.MPKData
	// bz := store.Get([]byte("mpkData3"))
	// mpkData3.MustUnmarshalBinaryBare(bz)
	// logrus.Info("one---------------------------", faulterId)
	// if (mpkData3.Id == dkgId){
		
	// 	if mpkData3.Pks[faulterId] == nil {
			
	// 		bz = store.Get([]byte("mpkData2"))
	// 		mpkData2.MustUnmarshalBinaryBare(bz)
			
	// 		if mpkData2.Pks[faulterId] == nil {
	// 			bz = store.Get([]byte("mpkData"))
	// 			mpkData.MustUnmarshalBinaryBare(bz)
	// 			logrus.Info("three---------------------------", mpkData)
	// 			mpkData.Pks[faulterId] = make([]byte, 48)
	// 			logrus.Info("four---------------------------", mpkData)
	// 			b:= mpkData.MustMarshalBinaryBare()
				
				
	// 			logrus.Info("byte---------------------------", b)
				
	// 			store.Set([]byte("mpkData"), b) 
	// 			logrus.Info("five---------------------------")
	// 			return
	// 		}

	// 		mpkData2.Pks[faulterId] = make([]byte, 48)
			
	// 			store.Set([]byte("mpkData2"), mpkData2.MustMarshalBinaryBare())
	// 			return
	// 	}
	// mpkData3.Pks[faulterId] = make([]byte, 48)
	// logrus.Info("six---------------------------", mpkData3)
	// store.Set([]byte("mpkData3"), mpkData3.MustMarshalBinaryBare())}

// }

func (k Keeper) InitializeList(ctx sdk.Context) {
	list := types.Faulters{
		FaultyList: []uint64{},
		Lookup: map[uint64]bool{},
	}
	store := ctx.KVStore(k.storeKey)
	bz := list.MustMarshalBinaryBare()
	store.Set([]byte("faulters"), bz)
}

func (k Keeper) SetList(ctx sdk.Context, list types.Faulters) {
	store := ctx.KVStore(k.storeKey)
	bz := list.MustMarshalBinaryBare()
	store.Set([]byte("faulters"), bz)
}

func (k Keeper) GetList(ctx sdk.Context) (list types.Faulters, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte("faulters"))
	if bz == nil {
		return list, false
	}
	list.MustUnmarshalBinaryBare(bz)
	return list, true
}

func (k Keeper) AddFaulter(ctx sdk.Context, number uint64) {
	list, found := k.GetList(ctx)
	if !found {
		list = types.Faulters{}
	}
	if !list.Lookup[number] { 
	list.FaultyList = append(list.FaultyList, number)
	list.Lookup[number] = true
	k.SetList(ctx, list)}
}

func (k Keeper) AddPk(ctx sdk.Context, pk []byte, id uint64){
	k.mu.Lock()
    defer k.mu.Unlock()
	store := ctx.KVStore(k.storeKey)
	var mpkData types.MPKData
	var mpkData2 types.MPKData
	var mpkData3 types.MPKData
	bz := store.Get([]byte("mpkData"))
	mpkData.MustUnmarshalBinaryBare(bz)
	logrus.Info("len---------------------------", len(mpkData.Pks) ,id)
	if len(mpkData.Pks) == 60 {
		
		bz = store.Get([]byte("mpkData2"))
		mpkData2.MustUnmarshalBinaryBare(bz)
		if len(mpkData2.Pks) == 60 {
			bz = store.Get([]byte("mpkData3"))
			mpkData3.MustUnmarshalBinaryBare(bz)
			mpkData3.Pks[id] = pk
			store.Set([]byte("mpkData3"), mpkData3.MustMarshalBinaryBare())
			return
		}
		mpkData2.Pks[id] = pk
		logrus.Info("before set ((((((((((()))))))))))")
			store.Set([]byte("mpkData2"), mpkData2.MustMarshalBinaryBare())
			logrus.Info("after set ((((((((((()))))))))))")
			return
	}
	//logrus.Info("here(((((((((((((((((((((((((((((((((((((((((((())))))))))))))))))))))))))))))))))))))))))))", id, mpkData.Pks)
	mpkData.Pks[id] = pk
	//logrus.Info("here after(((((((((((((((((((((((((((((((((((((((((((())))))))))))))))))))))))))))))))))))))))))))", id, mpkData.Pks)
	//logrus.Info("before set ((((((((((()))))))))))")
	store.Set([]byte("mpkData"), mpkData.MustMarshalBinaryBare())
	//logrus.Info("before set ((((((((((()))))))))))")

}


func (k Keeper) GetMPKData(ctx sdk.Context) types.MPKData{
	store := ctx.KVStore(k.storeKey)
	var mpkData types.MPKData
	var mpkData2 types.MPKData
	var mpkData3 types.MPKData
	bz := store.Get([]byte("mpkData"))
	bz2 := store.Get([]byte("mpkData2"))
	bz3 := store.Get([]byte("mpkData3"))
	mpkData.MustUnmarshalBinaryBare(bz)
	mpkData2.MustUnmarshalBinaryBare(bz2)
	mpkData3.MustUnmarshalBinaryBare(bz3)
	for k, v := range mpkData2.Pks {
        mpkData.Pks[k] = v
    }
	for k, v := range mpkData3.Pks {
        mpkData.Pks[k] = v
    }
	return mpkData

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
