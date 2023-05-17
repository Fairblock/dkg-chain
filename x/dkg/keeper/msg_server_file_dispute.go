package keeper

import (
	"context"
	//"crypto/cipher"
	"encoding/binary"

	"dkg/x/dkg/types"

	//	dsp "github.com/FairBlock/eth-dkg-go"

	//	bls "github.com/drand/kyber-bls12381"

	//"github.com/cosmos/cosmos-sdk/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) FileDispute(goCtx context.Context, msg *types.MsgFileDispute) (*types.MsgFileDisputeResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// var r cipher.Stream
	// dst := make([]byte, len(msg.Dispute.AddressOfAccusee))
	// src := make([]byte, len(msg.Dispute.AddressOfAccusee))
	// r.XORKeyStream(dst, src)

	// var slashed []byte
	// var dispute = types.Dispute{
	// 	AddressOfAccuser: msg.Dispute.AddressOfAccuser,
	// 	AddressOfAccusee: msg.Dispute.AddressOfAccusee,
	// 	Share:            msg.Dispute.Share,
	// 	Commit:           msg.Dispute.Commit,
	// 	Kij:              msg.Dispute.Kij,
	// 	CZkProof:         msg.Dispute.CZkProof,
	// 	RZkProof:         msg.Dispute.RZkProof,
	// 	Id:               msg.Dispute.Id,
	// }

	// count := k.GetDisputeCount(ctx)
	// dispute.Id = count
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisputeKey))
	// appendedValue := k.cdc.MustMarshal(&dispute)
	// store.Set(GetDisputeIDBytes(dispute.Id), appendedValue)
	// k.SetDisputeCount(ctx, count+1)

	// addOfAccusee := bls.NewBLS12381Suite().G1().Point().Base().Embed(dispute.AddressOfAccusee, r)
	// addOfAccuser := bls.NewBLS12381Suite().G1().Point().Base().Embed(dispute.AddressOfAccuser, r)
	// secretKeyIJ := bls.NewBLS12381Suite().G1().Point().Base().Embed(dispute.Kij, r)

	// rScalar := bls.NewKyberScalar().SetBytes(dispute.RZkProof)

	// res := dsp.VerifyProof(dsp.PointG, addOfAccusee, addOfAccuser, secretKeyIJ, dispute.CZkProof, rScalar)

	// if res {
	// 	//slash the accusee
	// 	slashed = dispute.AddressOfAccusee
	// }

	// if !res {
	// 	//slash the accuser
	// 	slashed = dispute.AddressOfAccuser
	// }

	// return &types.MsgFileDisputeResponse{Verdict: res, IdOfSlashedValidator: slashed}, nil
	return nil, nil
}

func (k Keeper) GetDisputeCount(ctx sdk.Context) uint64 {
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	// byteKey := types.KeyPrefix(types.DisputeCountKey)
	// bz := store.Get(byteKey)
	// if bz == nil {
	// 	return 0
	// }
	return 0
	// return binary.BigEndian.Uint64(bz)
}

func GetDisputeIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetDisputeCount(ctx sdk.Context, count uint64) {
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	// byteKey := types.KeyPrefix(types.DisputeCountKey)
	// bz := make([]byte, 8)
	// binary.BigEndian.PutUint64(bz, count)
	// store.Set(byteKey, bz)
}
