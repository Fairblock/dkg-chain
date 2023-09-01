package dkg

import (
	"context"
//	"math/rand"

	//"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"

	// this line is used by starport scaffolding # 1

	bls "github.com/drand/kyber-bls12381"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"dkg/x/dkg/client/cli"
	"dkg/x/dkg/keeper"
	"dkg/x/dkg/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface that defines the independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the name of the module as a string
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the amino codec for the module, which is used to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns a default GenesisState for the module, marshalled to json.RawMessage. The default GenesisState need to be defined by the module developer and is primarily used for testing
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root Tx command for the module. The subcommands of this root command are used by end-users to generate new transactions containing messages defined in the module
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the root query command for the module. The subcommands of this root command are used by end-users to generate new queries to the subset of the state defined by the module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

// Deprecated: use RegisterServices
func (am AppModule) Route() sdk.Route { return sdk.Route{} }

// Deprecated: use RegisterServices
func (AppModule) QuerierRoute() string { return types.RouterKey }

// Deprecated: use RegisterServices
func (am AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the invariants of the module. If an invariant deviates from its predicted value, the InvariantRegistry triggers appropriate logic (most often the chain will be halted)
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the module's genesis initialization. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion is a sequence number for state-breaking change of the module. It should be incremented on each consensus-breaking change introduced by the module. To avoid wrong/empty versions, the initial version should be set to 1
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock contains the logic that is automatically triggered at the beginning of each block
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

type Bcast struct {
	UIVssCommit vssCommit `json:"u_i_vss_commit"`
	ID          uint      `json:"id"`
}

type vssCommit struct {
	CoeffCommits [][]byte `json:"coeff_commits"`
}

// EndBlock contains the logic that is automatically triggered at the end of each block
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {

	timeoutData := am.keeper.GetTimeout(ctx)
	if ctx.BlockHeight() == int64(timeoutData.Start){
		
		participants := am.keeper.GetAddressList(ctx)
		participantsString, _ := json.Marshal(participants.Addresses)
		
		t:= len(participants.Addresses) * 1/2 
		am.keeper.InitCounter(ctx)
		am.keeper.InitMPK(ctx, timeoutData.Id)
		event := sdk.NewEvent(
			"keygen",
			sdk.NewAttribute("start", timeoutData.Id),
			sdk.NewAttribute("threshold", strconv.FormatInt(int64(t),10)),
			sdk.NewAttribute("participants", string(participantsString)),
			sdk.NewAttribute("timeout",strconv.FormatInt(int64(timeoutData.Timeout),10)),
			sdk.NewAttribute("module", "dkg"),
		)
		ctx.EventManager().EmitEvent(event)
	
		logrus.Info("start++++++++++++++++_____________________________________________________________: ", participants.Addresses)
		am.keeper.InitializeAddressList(ctx)
		return []abci.ValidatorUpdate{}
	}
	faultyList, _ := am.keeper.GetList(ctx)
	if timeoutData.Id != "" {
		desiredHeight := timeoutData.Timeout
		round := timeoutData.Round
		start := timeoutData.Start
		id := timeoutData.Id
		if round == 2 {
			if ctx.BlockHeight() == int64(uint64(start)+desiredHeight+120) {
				// Construct your event with attributes
				logrus.Info("End of round:", round)
				event := sdk.NewEvent(
					"dkg-timeout",
					sdk.NewAttribute("round", strconv.FormatUint(round, 10)),
					sdk.NewAttribute("id", id),
					// Add more attributes as needed
				)

				// Emit the event
				ctx.EventManager().EmitEvent(event)
				am.keeper.NextRound(ctx)
			}
		}
		if round == 0 {
			if ctx.BlockHeight() == int64(uint64(start)+30) {
				// Construct your event with attributes
				logrus.Info("End of round:", round)
				event := sdk.NewEvent(
					"dkg-timeout",
					sdk.NewAttribute("round", strconv.FormatUint(round, 10)),
					sdk.NewAttribute("id", id),
					// Add more attributes as needed
				)

				// Emit the event
				ctx.EventManager().EmitEvent(event)
				am.keeper.NextRound(ctx)
			}
		}
		if round == 1 {
			if ctx.BlockHeight() == int64(uint64(start)+90) {
				// Construct your event with attributes
				logrus.Info("End of round:", round)
				event := sdk.NewEvent(
					"dkg-timeout",
					sdk.NewAttribute("round", strconv.FormatUint(round, 10)),
					sdk.NewAttribute("id", id),
					// Add more attributes as needed
				)

				// Emit the event
				ctx.EventManager().EmitEvent(event)
				am.keeper.NextRound(ctx)
			}

		}

		if round == 3 {

			CalculateMPK(ctx, id, am.keeper.GetMPKData(ctx), faultyList.FaultyList)
			
			id, _ := strconv.Atoi(timeoutData.Id)
			
			am.keeper.InitTimeout(ctx, 0, timeoutData.Timeout, uint64(ctx.BlockHeight())+ 20, strconv.Itoa(id+1))

		}
	}

	return []abci.ValidatorUpdate{}
}

func CalculateMPK(ctx sdk.Context, id string, mpkData types.MPKData, faulters []uint64) {

	logrus.Info("+++++++++++++++++++++++++++++++++++ faulters:", faulters)
	suite := bls.NewBLS12381Suite()

	mpk := suite.G1().Point()

	if id != mpkData.Id {
		logrus.Panic("wrong mpk data")
	}
	for i := 0; i < len(mpkData.Pks); i++ {
		skip := false
		for j := 0; j < len(faulters); j++ {
			if faulters[j] == uint64(i) {
				skip = true
				break
			}
		}
		if !skip {

			if i == 0 {

				mpk.UnmarshalBinary(mpkData.Pks[uint64(i)])

			}
			if i != 0 {
				mpkPrime := suite.G1().Point()
				_ = mpkPrime.UnmarshalBinary(mpkData.Pks[uint64(i)])
				mpk = mpk.Add(mpk, mpkPrime)

			}

		}
	}
	pkb, _ := mpk.MarshalBinary()

	event := sdk.NewEvent(
		"dkg-mpk",
		sdk.NewAttribute("mpk", string(pkb)),
		sdk.NewAttribute("id", id),
		// Add more attributes as needed
	)

	// Emit the event
	ctx.EventManager().EmitEvent(event)
}
