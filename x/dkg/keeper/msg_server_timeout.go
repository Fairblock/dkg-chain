package keeper

import (
	"context"
	"encoding/binary"
	"encoding/json"

	//"strconv"

	"dkg/x/dkg/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bls "github.com/drand/kyber-bls12381"
	"github.com/sirupsen/logrus"
)
type Bcast struct {
	UIVssCommit vssCommit `json:"u_i_vss_commit"`
	ID          uint      `json:"id"`
}

type vssCommit struct {
	CoeffCommits [][]byte `json:"coeff_commits"`
}
func (k msgServer) Timeout(goCtx context.Context, msg *types.MsgTimeout) (*types.MsgTimeoutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	
	// c := ctx.Context()
		
	
	round := msg.Round 
	// r, _ := strconv.Atoi(round)
	// r = r+ 1
	id := msg.Id
		// Construct your event with attributes
		event := sdk.NewEvent(
			"dkg-timeout",
			sdk.NewAttribute("round", round),
			sdk.NewAttribute("id", msg.Id),
			// Add more attributes as needed
		)

		// Emit the event
		ctx.EventManager().EmitEvent(event)
		
	
	if round == "2" {
		CalculateMPK(ctx,id)
		
	}
	_ = ctx
	
	return &types.MsgTimeoutResponse{}, nil
}
func CalculateMPK(ctx sdk.Context, id string) {
	events := ctx.EventManager().Events()
	suite := bls.NewBLS12381Suite()

	mpk := suite.G1().Point()
	first := true
	faulters := []uint64{}
	for _, event := range events {
		eventType := event.Type
		attributes := event.Attributes
		
		if eventType == "keygen" {

			for _, attribute := range attributes {

				if string(attribute.Key) == "dispute" {

					faulters = append(faulters,  binary.BigEndian.Uint64(attribute.Value))
					logrus.Info(faulters)

				}}}}
	for _, event := range events {
		eventType := event.Type
		attributes := event.Attributes
		message := new(types.TrafficOut)
		if eventType == "keygen" {

			for _, attribute := range attributes {

				if string(attribute.Key) == "message" {
					message.Unmarshal(attribute.Value)
					if message.RoundNum == "2" {
						if message.IsBroadcast {
							var bcast Bcast
							err := json.Unmarshal(message.Payload, &bcast)
							if err != nil {
								logrus.Error("Error:", err)
							}
							for i := 0; i < len(faulters); i++ {
								if faulters[i] == uint64(bcast.ID){
									break
								}
							}
							if first {
								mpk.UnmarshalBinary(bcast.UIVssCommit.CoeffCommits[0])
								first = false
								break
							}
							mpkPrime := suite.G1().Point()
							mpkPrime.UnmarshalBinary(bcast.UIVssCommit.CoeffCommits[0])
							mpk = suite.G1().Point().Add(mpk,mpkPrime)
						}
					}
				}
			}
		}

	}
	event := sdk.NewEvent(
		"dkg-mpk",
		sdk.NewAttribute("mpk", mpk.String()),
		sdk.NewAttribute("id", id),
		// Add more attributes as needed
	)

	// Emit the event
	ctx.EventManager().EmitEvent(event)	
}