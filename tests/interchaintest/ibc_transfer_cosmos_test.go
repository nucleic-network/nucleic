package interchaintest

import (
	"context"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/strangelove-ventures/interchaintest/v8"

	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	interchaintestrelayer "github.com/strangelove-ventures/interchaintest/v8/relayer"
	"github.com/strangelove-ventures/interchaintest/v8/testreporter"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
)

// TestEvejunoIBCTransfer spins up a Eve and juno network, initializes an IBC connection between them,
// and sends an ICS20 token transfer from Eve->juno and then back from juno->Eve.
func TestEveGaiaIBCTransfer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	// Create chain factory with Eve and juno
	numVals := 1
	numFullNodes := 0

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:          "Eve",
			ChainConfig:   EveConfig,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
		{
			Name:          "juno",
			Version:       "v13.0.0",
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
			ChainConfig:   ibc.ChainConfig{GasPrices: "0.0ujuno"},
		},
	})

	const (
		path = "ibc-path"
	)

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	client, network := interchaintest.DockerSetup(t)

	eve, juno := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

	relayerType, relayerName := ibc.CosmosRly, "relay"

	// Get a relayer instance
	rf := interchaintest.NewBuiltinRelayerFactory(
		relayerType,
		zaptest.NewLogger(t),
		interchaintestrelayer.DockerImage(&DefaultRelayer),
	)

	r := rf.Build(t, client, network)

	ic := interchaintest.NewInterchain().
		AddChain(eve).
		AddChain(juno).
		AddRelayer(r, relayerName).
		AddLink(interchaintest.InterchainLink{
			Chain1:  eve,
			Chain2:  juno,
			Relayer: r,
			Path:    path,
			CreateChannelOpts: ibc.CreateChannelOptions{
				SourcePortName: "transfer",
				DestPortName:   "transfer",
				Order:          ibc.Unordered,
				Version:        "{\"fee_version\":\"ics29-1\",\"app_version\":\"ics20-1\"}",
			},
			CreateClientOpts: ibc.DefaultClientOpts(),
		})

	ctx := context.Background()

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,
		// BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
		SkipPathCreation: false,
	}))
	t.Parallel()

	err = testutil.WaitForBlocks(ctx, 5, eve, juno)

	require.NoError(t, err)
	// ChainID of eve
	chainIDA := eve.Config().ChainID

	// Channel of eve
	chA, err := r.GetChannels(ctx, eRep, chainIDA)
	require.NoError(t, err)
	channelA := chA[0]

	// Fund a user account on eve and juno
	initBal := sdkmath.NewInt(1_000_000_000_000)
	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), initBal, eve, juno)
	userA := users[0]
	userAddressA := userA.FormattedAddress()
	userB := users[1]
	userAddressB := userB.FormattedAddress()

	// Addresses of both the chains
	walletA, _ := r.GetWallet(eve.Config().ChainID)
	rlyAddressA := walletA.FormattedAddress()

	walletB, _ := r.GetWallet(juno.Config().ChainID)
	rlyAddressB := walletB.FormattedAddress()

	// // register CounterpartyPayee
	// cmd := []string{
	// 	"tx", "register-counterparty",
	// 	eve.Config().Name,
	// 	channelA.ChannelID,
	// 	"transfer",
	// 	rlyAddressA,
	// 	rlyAddressB,
	// }
	// _ = r.Exec(ctx, eRep, cmd, nil)
	// require.NoError(t, err)
	// err = testutil.WaitForBlocks(ctx, 20, eve, juno)
	// require.NoError(t, err)
	// Query the relayer CounterpartyPayee on a given channel
	// query := []string{
	// 	eve.Config().Bin, "query", "ibc-fee", "counterparty-payee", channelA.ChannelID, rlyAddressA,
	// 	"--node", eve.GetRPCAddress(),
	// 	"--home", eve.HomeDir(),
	// 	"--trace",
	// }
	// _, _, err = eve.Exec(ctx, query, nil)
	// require.NoError(t, err)

	t.Cleanup(func() {
		_ = ic.Close()
	})

	// Get initial account balances
	userAOrigBal, err := eve.GetBalance(ctx, userAddressA, eve.Config().Denom)
	require.NoError(t, err)
	require.True(t, initBal.Equal(userAOrigBal))

	userBOrigBal, err := juno.GetBalance(ctx, userAddressB, juno.Config().Denom)
	require.NoError(t, err)
	require.True(t, initBal.Equal(userBOrigBal))

	rlyAOrigBal, err := eve.GetBalance(ctx, rlyAddressA, eve.Config().Denom)
	require.NoError(t, err)
	require.True(t, initBal.Equal(rlyAOrigBal))

	rlyBOrigBal, err := juno.GetBalance(ctx, rlyAddressB, juno.Config().Denom)
	require.NoError(t, err)
	require.True(t, initBal.Equal(rlyBOrigBal))

	// send tx
	txAmount := sdkmath.NewInt(1000)
	transfer := ibc.WalletAmount{Address: userAddressB, Denom: eve.Config().Denom, Amount: txAmount}
	_, err = eve.SendIBCTransfer(ctx, channelA.ChannelID, userAddressA, transfer, ibc.TransferOptions{})
	require.NoError(t, err)

	// // Incentivizing async packet by returning MsgPayPacketFeeAsync
	// packetFeeAsync := []string{
	// 	eve.Config().Bin, "tx", "ibc-fee", "pay-packet-fee", "transfer", channelA.ChannelID, "1",
	// 	"--recv-fee", fmt.Sprintf("1000%s", eve.Config().Denom),
	// 	"--ack-fee", fmt.Sprintf("1000%s", eve.Config().Denom),
	// 	"--timeout-fee", fmt.Sprintf("1000%s", eve.Config().Denom),
	// 	"--chain-id", chainIDA,
	// 	"--node", eve.GetRPCAddress(),
	// 	"--from", userA.FormattedAddress(),
	// 	"--keyring-backend", "test",
	// 	"--gas", "400000",
	// 	"--yes",
	// 	"--home", eve.HomeDir(),
	// }
	// _, _, err = eve.Exec(ctx, packetFeeAsync, nil)
	// require.NoError(t, err)
	// err = testutil.WaitForBlocks(ctx, 5, eve, juno)
	// require.NoError(t, err)
	// start the relayer
	err = r.StartRelayer(ctx, eRep, path)
	require.NoError(t, err)

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				t.Logf("an error occurred while stopping the relayer: %s", err)
			}
		},
	)

	// Wait for relayer to run
	err = testutil.WaitForBlocks(ctx, 10, eve, juno)
	require.NoError(t, err)

	// Assigning denom
	eveTokenDenom := transfertypes.GetPrefixedDenom(channelA.PortID, channelA.ChannelID, eve.Config().Denom)
	eveDenomTrace := transfertypes.ParseDenomTrace(eveTokenDenom)

	// Get balances after transfer
	expectedBal := userAOrigBal
	eveBal, err := eve.GetBalance(ctx, userAddressA, eve.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, expectedBal.Sub(txAmount), eveBal)

	junoBal, err := juno.GetBalance(ctx, userAddressB, eveDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, txAmount, junoBal)

	// rlyABal, err := eve.GetBalance(ctx, rlyAddressA, eve.Config().Denom)
	// require.NoError(t, err)
	// require.True(t, rlyAOrigBal.AddRaw(1000).Equal(rlyABal))
}
