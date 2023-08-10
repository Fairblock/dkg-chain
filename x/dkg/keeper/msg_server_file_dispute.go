package keeper

import (
	"context"
	"strconv"
	"strings"

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

	//"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sirupsen/logrus"

	//bls12381 "github.com/kilic/bls12-381"

	distIBE "github.com/FairBlock/DistributedIBE"

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
	share           distIBE.Share
	commit          distIBE.Commitments
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

type P struct {
	Concat  string
	G       []byte
	Pkj     []byte
	Pki     []byte
	Ek      []byte
	C       []byte
	R       []byte
	cNew    []byte
	X1t     []byte
	Y1t     []byte
	T1P     []byte
	commits [][]byte
}

func VerifyProof(pointG, publicKeyI, publicKeyJ, encryptionKeyIJ kyber.Point, c []byte, r kyber.Scalar, cReal []byte) bool {
	reverseBytes(c)
	c2KyberScalar := bls.NewKyberScalar()
	c2KyberScalar.UnmarshalBinary(c)
	//c2KyberScalar := bls.NewKyberScalar().SetBytes(c)

	//	x1ToR := bls.NewBLS12381Suite().G1().Point().Base()
	//	y1ToC := bls.NewBLS12381Suite().G1().Point().Base()
	//	t1Prime := bls.NewBLS12381Suite().G1().Point().Base()
	suite := bls.NewBLS12381Suite()
	x1ToR := suite.G1().Point()

	x1ToR.Mul(r, pointG)
	g, _ := pointG.MarshalBinary()
	//cb , _ := c2KyberScalar.MarshalBinary()
	pkj, _ := publicKeyJ.MarshalBinary()
	pki, _ := publicKeyI.MarshalBinary()
	ek, _ := encryptionKeyIJ.MarshalBinary()
	// rr,_:= r.MarshalBinary()

	// x1t,_ := x1ToR.MarshalBinary()
	// panic(P{G:g,R:rr,X1t: x1t})
	//x1ToR := bls.NewBLS12381Suite().G1().Point().Mul(r,pointG)
	y1ToC := bls.NewBLS12381Suite().G1().Point().Mul(c2KyberScalar, publicKeyJ)
	t1Prime := bls.NewBLS12381Suite().G1().Point().Add(x1ToR, y1ToC)
	//x1t,_ := y1ToC.MarshalBinary()

	x2ToR := bls.NewBLS12381Suite().G1().Point().Base()
	y2ToC := bls.NewBLS12381Suite().G1().Point().Base()
	t2Prime := bls.NewBLS12381Suite().G1().Point().Base()

	x2ToR.Mul(r, publicKeyI)
	y2ToC.Mul(c2KyberScalar, encryptionKeyIJ)
	t2Prime.Add(x2ToR, y2ToC)
	t1p, _ := t1Prime.MarshalBinary()
	t2p, _ := t2Prime.MarshalBinary()
	//concat := bytes.Join([][]byte{g, pkj,pki, ek, t1p, t2p}, []byte{})
	//s := concat

	//y1t,_ := y1ToC.MarshalBinary()

	gstr := make([]string, len(g))
	for i, b := range g {
		gstr[i] = strconv.Itoa(int(b))
	}
	pkjstr := make([]string, len(pkj))
	for i, b := range pkj {
		pkjstr[i] = strconv.Itoa(int(b))
	}
	pkistr := make([]string, len(pki))
	for i, b := range pki {
		pkistr[i] = strconv.Itoa(int(b))
	}
	ekstr := make([]string, len(ek))
	for i, b := range ek {
		ekstr[i] = strconv.Itoa(int(b))
	}
	t1pstr := make([]string, len(t1p))
	for i, b := range t1p {
		t1pstr[i] = strconv.Itoa(int(b))
	}
	t2pstr := make([]string, len(t2p))
	for i, b := range t2p {
		t2pstr[i] = strconv.Itoa(int(b))
	}

	s := "[" + strings.Join(gstr, ", ") + "]" + "[" + strings.Join(pkjstr, ", ") + "]" + "[" + strings.Join(pkistr, ", ") + "]" + "[" + strings.Join(ekstr, ", ") + "]" + "[" + strings.Join(t1pstr, ", ") + "]" + "[" + strings.Join(t2pstr, ", ") + "]"
	//panic(P{C:c,cNew: t1p})
	//	concat := []string{string(g), string(pkj), string(pki), string(ek), string(t1p), string(t2p)}
	//	s := strings.Join(concat, "")
	h := sha256.New()
	h.Write([]byte(s))
	cPrime := h.Sum(nil)
	//panic(P{C: cReal ,cNew:cPrime})
	return bytes.Equal(cPrime, cReal)

}

// func handleDispute(listOfComplaints []listOfComplaints) {

// 	for i := 0; i < len(listOfComplaints); i++ {
// 		verifyResult := VerifyProof(PointG, listOfComplaints[i].pubKeyOfAccuser, listOfComplaints[i].pubKeyOfAccusee, listOfComplaints[i].kij, listOfComplaints[i].proof.hash, listOfComplaints[i].proof.scalar)
// 		if !verifyResult {
// 			fmt.Println("Accuser is malicous")
// 			// slash the accuser goes here
// 		}
// 		if distIBE.VerifyShare(listOfComplaints[i].share, listOfComplaints[i].commit) {
// 			fmt.Println("Accusee is malicious")
// 			//slash the accusee goes here
// 		} else {
// 			fmt.Println("Accuser is malicious")
// 			// slash the accuser
// 		}
// 	}

// 	//the result can be a column added to the complaint list that indicates the result of the complaint
// }
func reverseBytes(data []byte) {
	length := len(data)
	for i := 0; i < length/2; i++ {
		data[i], data[length-i-1] = data[length-i-1], data[i]
	}
}
func (k msgServer) FileDispute(goCtx context.Context, msg *types.MsgFileDispute) (*types.MsgFileDisputeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//var r cipher.Stream
	// dst := make([]byte, len(msg.Dispute.AddressOfAccusee))
	// src := make([]byte, len(msg.Dispute.AddressOfAccusee))
	// r.XORKeyStream(dst, src)

	var slashed string
	var dispute = types.Dispute{
		AddressOfAccuser: msg.Dispute.AddressOfAccuser,
		AddressOfAccusee: msg.Dispute.AddressOfAccusee,
		Share:            msg.Dispute.Share,
		Commit:           msg.Dispute.Commit,
		Kij:              msg.Dispute.Kij,
		CZkProof:         msg.Dispute.CZkProof,
		RZkProof:         msg.Dispute.RZkProof,
		Id:               msg.Dispute.Id,
		AccuserId:        msg.Dispute.AccuserId,
		FaulterId:        msg.Dispute.FaulterId,
		CReal:            msg.Dispute.CReal,
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

	reverseBytes(dispute.RZkProof)
	rScalar := bls.NewKyberScalar()
	rScalar.UnmarshalBinary(dispute.RZkProof)
	//rs, _ := rScalar.MarshalBinary()

	res := VerifyProof(PointG, addOfAccusee, addOfAccuser, secretKeyIJ, dispute.CZkProof, rScalar, dispute.CReal)

	if res {
		//slash the accusee
		commits := make([]kyber.Point, 0)
		for i := 0; i < len(dispute.Commit.Commitments); i++ {

			c := suite.G1().Point()
			c.UnmarshalBinary(dispute.Commit.Commitments[i])
			commits = append(commits, c)
		}

		//reverseBytes(dispute.Share.Index)
		reverseBytes(dispute.Share.Value)
		sharei := bls.NewKyberScalar()
		sharei.SetInt64(int64(binary.BigEndian.Uint32(dispute.Share.Index) + 1))
		sharev := bls.NewKyberScalar()
		sharev.UnmarshalBinary(dispute.Share.Value)
		//panic(P{Concat: sharev.String(),C:dispute.Share.Index,R:dispute.Share.Value, commits: dispute.Commit.Commitments})
		verify := distIBE.VerifyVSSShare(distIBE.Share{Index: sharei, Value: sharev}, commits)
		if !verify {
			//panic("accusee")
			slashed = string(rune(dispute.FaulterId))
		}
		if verify {
			//panic("accuser 1")
			//slash the accuser
			slashed = string(rune(dispute.AccuserId))
		}
	}

	if !res {
		//panic("accuser 2")
		//slash the accuser
		slashed = string(rune(dispute.AccuserId))
	}
	counting := k.IncreaseCounter(ctx, 1)
	str_count := strconv.FormatUint(counting, 10)
	event := sdk.NewEvent(
		types.EventTypeKeygen,
		sdk.NewAttribute(types.AttributeValueDispute, slashed),
		sdk.NewAttribute("keyID", msg.KeyId),
		sdk.NewAttribute("from", msg.Creator),
		sdk.NewAttribute("index", str_count),
		sdk.NewAttribute("module", "dkg"),
	)
	ctx.EventManager().EmitEvent(event)
	logrus.Info("------------ indexDispute1: ", str_count)
	
	faulter, _ := strconv.Atoi(slashed)
	k.AddFaulter(ctx,uint64(faulter),msg.KeyId)
	logrus.Info("------------ indexDispute2: ", str_count)
	
	return &types.MsgFileDisputeResponse{Verdict: res, IdOfSlashedValidator: []byte(slashed)}, nil
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
