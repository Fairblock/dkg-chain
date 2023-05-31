package keeper

import (
	"context"
	//"crypto/cipher"

	//"crypto/cipher"
	"encoding/binary"

	"dkg/x/dkg/types"

	//	dsp "github.com/FairBlock/eth-dkg-go"

	//	bls "github.com/drand/kyber-bls12381"

	//"github.com/cosmos/cosmos-sdk/store/prefix"

	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//bls12381 "github.com/kilic/bls12-381"

	vsskyber "github.com/FairBlock/vsskyber"

	"github.com/drand/kyber"
	bls "github.com/drand/kyber-bls12381"
)

type pubAndPrivKey struct {
	publicKey  kyber.Point
	privateKey kyber.Scalar
}

type zkProof struct {
	hash   []byte
	scalar kyber.Scalar
}

type listOfComplaints struct {
	pubKeyOfAccuser kyber.Point
	pubKeyOfAccusee kyber.Point
	share           vsskyber.Share
	commit          vsskyber.Commitments
	kij             kyber.Point
	proof           zkProof
}

var PointG = bls.NewBLS12381Suite().G1().Point().Base()

const groupOrderInHex = "73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001"

var groupOrder, _ = new(big.Int).SetString(groupOrderInHex, 16)

func initInput(numOfUsers int, threshInPercent float64) (threshold int, PublicAndPrivateKeyArray []pubAndPrivKey, err error) {

	PublicAndPrivateKeyArray = make([]pubAndPrivKey, numOfUsers)

	for i := 0; i < numOfUsers; i++ {
		one := big.NewInt(int64(1))
		max := big.NewInt(int64(0))
		max.Sub(groupOrder, one)

		w, err := rand.Int(rand.Reader, max)
		if err != nil {
			return threshold, PublicAndPrivateKeyArray, fmt.Errorf("unable to generate secret key for users")
		}
		w.Add(w, one)

		privateKey := bls.NewKyberScalar()
		privateKey = kyber.Scalar.SetInt64(privateKey, w.Int64())

		publicKey := PointG.Mul(privateKey, PointG)

		PublicAndPrivateKeyArray[i].publicKey = publicKey
		PublicAndPrivateKeyArray[i].privateKey = privateKey

	}

	threshold = int(float64(numOfUsers) * threshInPercent)

	return threshold, PublicAndPrivateKeyArray, nil
}

func VerifyProof(pointG, publicKeyJ, publicKeyI, encryptionKeyIJ kyber.Point, c []byte, r kyber.Scalar) bool {

	c2KyberScalar := bls.NewKyberScalar().SetBytes(c)

	x1ToR := bls.NewBLS12381Suite().G1().Point().Base()
	y1ToC := bls.NewBLS12381Suite().G1().Point().Base()
	t1Prime := bls.NewBLS12381Suite().G1().Point().Base()

	x1ToR.Mul(r, pointG)
	y1ToC.Mul(c2KyberScalar, publicKeyJ)
	t1Prime.Add(x1ToR, y1ToC)

	x2ToR := bls.NewBLS12381Suite().G1().Point().Base()
	y2ToC := bls.NewBLS12381Suite().G1().Point().Base()
	t2Prime := bls.NewBLS12381Suite().G1().Point().Base()

	x2ToR.Mul(r, publicKeyI)
	y2ToC.Mul(c2KyberScalar, encryptionKeyIJ)
	t2Prime.Add(x2ToR, y2ToC)

	concat := []string{pointG.String(), publicKeyJ.String(), publicKeyI.String(), encryptionKeyIJ.String(), t1Prime.String(), t2Prime.String()}
	s := strings.Join(concat, "")
	h := sha256.New()
	h.Write([]byte(s))
	cPrime := h.Sum(nil)

	return bytes.Equal(cPrime, c)

}

func handleDispute(listOfComplaints []listOfComplaints) {

	for i := 0; i < len(listOfComplaints); i++ {
		verifyResult := VerifyProof(PointG, listOfComplaints[i].pubKeyOfAccuser, listOfComplaints[i].pubKeyOfAccusee, listOfComplaints[i].kij, listOfComplaints[i].proof.hash, listOfComplaints[i].proof.scalar)
		if !verifyResult {
			fmt.Println("Accuser is malicous")
			// slash the accuser goes here
		}
		if vsskyber.VerifyShare(listOfComplaints[i].share, listOfComplaints[i].commit) {
			fmt.Println("Accusee is malicious")
			//slash the accusee goes here
		} else {
			fmt.Println("Accuser is malicious")
			// slash the accuser
		}
	}

	//the result can be a column added to the complaint list that indicates the result of the complaint
}
func (k msgServer) FileDispute(goCtx context.Context, msg *types.MsgFileDispute) (*types.MsgFileDisputeResponse, error) {
	 ctx := sdk.UnwrapSDKContext(goCtx)

	//var r cipher.Stream
	// dst := make([]byte, len(msg.Dispute.AddressOfAccusee))
	// src := make([]byte, len(msg.Dispute.AddressOfAccusee))
	// r.XORKeyStream(dst, src)

	var slashed []byte
	var dispute = types.Dispute{
		AddressOfAccuser: msg.Dispute.AddressOfAccuser,
		AddressOfAccusee: msg.Dispute.AddressOfAccusee,
		Share:            msg.Dispute.Share,
		Commit:           msg.Dispute.Commit,
		Kij:              msg.Dispute.Kij,
		CZkProof:         msg.Dispute.CZkProof,
		RZkProof:         msg.Dispute.RZkProof,
		Id:               msg.Dispute.Id,
	}

	count := k.GetDisputeCount(ctx)
	dispute.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisputeKey))
	appendedValue := k.cdc.MustMarshal(&dispute)
	store.Set(GetDisputeIDBytes(dispute.Id), appendedValue)
	k.SetDisputeCount(ctx, count+1)
	suite := bls.NewBLS12381Suite()

	
	 addOfAccusee := suite.G1().Point()
	 addOfAccuser := suite.G1().Point()
	 secretKeyIJ := suite.G1().Point()
//panic(dispute.AddressOfAccusee)
	secretKeyIJ.UnmarshalBinary(dispute.Kij)
	addOfAccusee.UnmarshalBinary(dispute.AddressOfAccusee)
	addOfAccuser.UnmarshalBinary(dispute.AddressOfAccuser)
	
	

	rScalar := bls.NewKyberScalar().SetBytes(dispute.RZkProof)

	res := VerifyProof(PointG, addOfAccusee, addOfAccuser, secretKeyIJ, dispute.CZkProof, rScalar)

	if res {
		//slash the accusee
		commits := new([]kyber.Point)
		for i := 0; i < len(dispute.Commit.Commitments); i++ {
			(*commits)[i] = suite.G1().Point()
			(*commits)[i].UnmarshalBinary(dispute.Commit.Commitments[i])
		}
		index := new(kyber.Scalar)
		(*index).SetBytes(dispute.Share.Index)
		value := new(kyber.Scalar)
		(*value).SetBytes(dispute.Share.Value)
		verify := vsskyber.VerifyShare(vsskyber.Share{Index: *index, Value: *value },*commits)
		if !verify{
		slashed = dispute.AddressOfAccusee}
		if verify {
			//slash the accuser
			slashed = dispute.AddressOfAccuser
		}
	}

	if !res {
		//slash the accuser
		slashed = dispute.AddressOfAccuser
	}
	event := sdk.NewEvent(
		types.EventTypeKeygen,
		sdk.NewAttribute(types.AttributeValueDispute, string(slashed)),
		sdk.NewAttribute("keyID", msg.KeyId),
		sdk.NewAttribute("from", msg.Creator),
		sdk.NewAttribute("module", "dkg"),
	)
	ctx.EventManager().EmitEvent(event)
	return &types.MsgFileDisputeResponse{Verdict: res, IdOfSlashedValidator: slashed}, nil
	//return nil, nil
}

func (k Keeper) GetDisputeCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DisputeCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
//	return 0
	return binary.BigEndian.Uint64(bz)
}

func GetDisputeIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetDisputeCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DisputeCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}
