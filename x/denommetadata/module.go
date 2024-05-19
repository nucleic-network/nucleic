package denommetadata

import (
	"context"
	"encoding/json"

	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/eve-network/eve/x/denommetadata/keeper"
	"github.com/eve-network/eve/x/denommetadata/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var (
	_ module.AppModuleBasic   = AppModuleBasic{}
	_ module.HasGenesisBasics = AppModuleBasic{}

	_ appmodule.AppModule        = AppModule{}
	_ module.HasConsensusVersion = AppModule{}
	_ module.HasGenesis          = AppModule{}
	_ module.HasServices         = AppModule{}
)

// AppModuleBasic implements the AppModuleBasic interface that defines the
// independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
}

// DefaultGenesis implements module.HasGenesisBasics.
func (a AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis implements module.HasGenesisBasics.
func (a AppModuleBasic) ValidateGenesis(codec.JSONCodec, client.TxEncodingConfig, json.RawMessage) error {
	return nil
}

// Name implements module.AppModuleBasic.
func (a AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterGRPCGatewayRoutes implements module.AppModuleBasic.
func (a AppModuleBasic) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {
}

// RegisterInterfaces implements module.AppModuleBasic.
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// RegisterLegacyAminoCodec implements module.AppModuleBasic.
func (a AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func NewAppModuleBasic() AppModuleBasic {
	return AppModuleBasic{}
}

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement
type AppModule struct {
	AppModuleBasic

	keeper     keeper.Keeper
	bankKeeper bankkeeper.Keeper
}

// RegisterServices implements module.HasServices.
func (a AppModule) RegisterServices(module.Configurator) {
}

// DefaultGenesis implements module.HasGenesis.
func (a AppModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ExportGenesis implements module.HasGenesis.
func (a AppModule) ExportGenesis(sdk.Context, codec.JSONCodec) json.RawMessage {
	return nil
}

// InitGenesis implements module.HasGenesis.
func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) {
	a.bankKeeper.IterateAllDenomMetaData(ctx, func(metadata banktypes.Metadata) bool {
		// run hooks for each denom metadata, thus `x/denommetadata` genesis init order must be after `x/bank` genesis init
		err := a.keeper.GetHooks().AfterDenomMetadataCreation(ctx, metadata)
		if err != nil {
			panic(err) // error at genesis level should be reported by panic
		}

		return false
	})
}

// ValidateGenesis implements module.HasGenesis.
func (a AppModule) ValidateGenesis(codec.JSONCodec, client.TxEncodingConfig, json.RawMessage) error {
	panic("unimplemented")
}

// ConsensusVersion implements module.HasConsensusVersion.
func (a AppModule) ConsensusVersion() uint64 {
	return 1
}

// IsAppModule implements appmodule.AppModule.
func (a AppModule) IsAppModule() {
}

// IsOnePerModuleType implements appmodule.AppModule.
func (a AppModule) IsOnePerModuleType() {
	panic("unimplemented")
}

// EndBlock implements appmodule.HasEndBlocker.
func (a *AppModule) EndBlock(context.Context) error {
	panic("unimplemented")
}

// BeginBlock implements appmodule.HasBeginBlocker.
func (a *AppModule) BeginBlock(context.Context) error {
	panic("unimplemented")
}

// GenerateGenesisState implements module.AppModuleSimulation.
func (a *AppModule) GenerateGenesisState(input *module.SimulationState) {
	panic("unimplemented")
}

// RegisterStoreDecoder implements module.AppModuleSimulation.
func (a *AppModule) RegisterStoreDecoder(simulation.StoreDecoderRegistry) {
	panic("unimplemented")
}

// WeightedOperations implements module.AppModuleSimulation.
func (a *AppModule) WeightedOperations(simState module.SimulationState) []simulation.WeightedOperation {
	panic("unimplemented")
}

// Name implements module.AppModuleBasic.
func (a *AppModule) Name() string {
	return types.ModuleName
}

// RegisterGRPCGatewayRoutes implements module.AppModuleBasic.
func (a *AppModule) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {
	panic("unimplemented")
}

// RegisterInterfaces implements module.AppModuleBasic.
func (a *AppModule) RegisterInterfaces(cdctypes.InterfaceRegistry) {
	panic("unimplemented")
}

// RegisterLegacyAminoCodec implements module.AppModuleBasic.
func (a *AppModule) RegisterLegacyAminoCodec(*codec.LegacyAmino) {
	panic("unimplemented")
}

func NewAppModule(
	keeper keeper.Keeper,
	bankKeeper bankkeeper.Keeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(),
		keeper:         keeper,
		bankKeeper:     bankKeeper,
	}
}
