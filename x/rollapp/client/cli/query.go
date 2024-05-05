package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/eve-network/eve/x/rollapp/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group markets queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdQueryLatestHeight())
	cmd.AddCommand(CmdQueryLatestStateInfoIndex())
	cmd.AddCommand(CmdQueryStateInfo())
	cmd.AddCommand(CmdQueryRollapp())
	cmd.AddCommand(CmdQueryListRollapp())

	return cmd
}