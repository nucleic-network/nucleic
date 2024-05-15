package interchaintest

import (
	"os"

	sdkmath "cosmossdk.io/math"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"

	"github.com/strangelove-ventures/interchaintest/v8/ibc"
)

var (
	EveICTestRepo = "ghcr.io/eve-network/eve"

	repo, version = GetDockerImageInfo()

	EveImage = ibc.DockerImage{
		Repository: repo,
		Version:    version,
		UidGid:     "1025:1025",
	}

	EveConfig = ibc.ChainConfig{
		Type:                "cosmos",
		Name:                "eve",
		ChainID:             "eve-2",
		Images:              []ibc.DockerImage{EveImage},
		Bin:                 "eved",
		Bech32Prefix:        "eve",
		Denom:               "stake",
		CoinType:            "1811",
		GasPrices:           "0.0stake",
		GasAdjustment:       1.1,
		TrustingPeriod:      "112h",
		NoHostMount:         false,
		EncodingConfig:      eveEncoding(),
		ModifyGenesis:       nil,
		ConfigFileOverrides: nil,
	}

	genesisWalletAmount = sdkmath.NewInt(10_000_000)

	DefaultRelayer = ibc.DockerImage{
		Repository: "ghcr.io/cosmos/relayer",
		Version:    "main",
		UidGid:     "1025:1025",
	}
)

// GetDockerImageInfo returns the appropriate repo and branch version string for integration with the CI pipeline.
// The remote runner sets the BRANCH_CI env var. If present, interchaintest will use the docker image pushed up to the repo.
// If testing locally, user should run `make docker-build-debug` and interchaintest will use the local image.
func GetDockerImageInfo() (repo, version string) {
	branchVersion, found := os.LookupEnv("BRANCH_CI")
	repo = EveICTestRepo
	if !found {
		// make local-image
		repo = "eve"
		branchVersion = "debug"
	}
	return repo, branchVersion
}

func eveEncoding() *testutil.TestEncodingConfig {
	cfg := cosmos.DefaultEncoding()

	// register custom types
	wasmtypes.RegisterInterfaces(cfg.InterfaceRegistry)
	return &cfg
}
