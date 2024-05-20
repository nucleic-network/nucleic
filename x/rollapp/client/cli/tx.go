package cli

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/eve-network/eve/utils"
	"github.com/eve-network/eve/x/rollapp/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

type PermissionedAddresses struct {
	Addresses []string `json:"addresses"`
}

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateRollapp())

	return cmd
}

func CmdCreateRollapp() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-rollapp [rollapp-id] [max-sequencers] [permissioned-addresses] [metadata.json]",
		Short:   "Create a new rollapp",
		Example: "nucleicd tx rollapp create-rollapp ROLLAPP_CHAIN_ID 10 '{\"Addresses\":[]}' metadata.json",
		Args:    cobra.RangeArgs(3, 4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRollappId := args[0]

			argMaxSequencers, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			var argPermissionedAddresses PermissionedAddresses
			if err := json.Unmarshal([]byte(args[2]), &argPermissionedAddresses); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// Parse metadata
			var metadatas []types.TokenMetadata
			if len(args) == 4 {
				metadatas, err = utils.ParseJsonFromFile[types.TokenMetadata](args[3])
				if err != nil {
					return err
				}
			}
			// Parse genesis accounts
			genesisAccountsPath, _ := cmd.Flags().GetString(FlagGenesisAccountsPath)
			genesisAccounts, err := utils.ParseJsonFromFile[types.GenesisAccount](genesisAccountsPath)
			if err != nil && genesisAccountsPath != "" {
				return err
			}

			msg := types.NewMsgCreateRollapp(
				clientCtx.GetFromAddress().String(),
				argRollappId,
				argMaxSequencers,
				argPermissionedAddresses.Addresses,
				metadatas,
				genesisAccounts,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetCreateRollapp())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-state [rollapp-id] [start-height] [num-blocks] [da-path] [version] [block-desc]",
		Short: "Update rollapp state",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRollappId := args[0]
			argStartHeight, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argNumBlocks, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argDAPath := args[3]
			argVersion, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			argBlockDescs := new(types.BlockDescriptors)
			err = json.Unmarshal([]byte(args[5]), argBlockDescs)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateState(
				clientCtx.GetFromAddress().String(),
				argRollappId,
				argStartHeight,
				argNumBlocks,
				argDAPath,
				argVersion,
				argBlockDescs,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
