package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/eve-network/eve/x/rollapp/types"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "shows the parameters of the module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryLatestHeight() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "latest-height [rollapp-id]",
		Short: "Query the last height of the last UpdateState associated with the specified rollapp-id.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argRollappId := args[0]

			argFinalized, err := cmd.Flags().GetBool(FlagFinalized)
			if err != nil {
				return err
			}

			req := &types.QueryGetLatestHeightRequest{
				RollappId: argRollappId,
				Finalized: argFinalized,
			}

			res, err := queryClient.LatestHeight(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().Bool(FlagFinalized, false, "Indicates whether to return the latest finalized state index")

	flags.AddQueryFlagsToCmd(cmd)

	return cmd

}

func CmdQueryLatestStateInfoIndex() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "latest-state-index [rollapp-id]",
		Short: "Query the index of the last UpdateState associated with the specified rollapp-id.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argRollappId := args[0]

			argFinalized, err := cmd.Flags().GetBool(FlagFinalized)
			if err != nil {
				return err
			}

			params := &types.QueryGetLatestStateIndexRequest{
				RollappId: argRollappId,
				Finalized: argFinalized,
			}

			res, err := queryClient.LatestStateIndex(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().Bool(FlagFinalized, false, "Indicates whether to return the latest finalized state index")

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryListRollapp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Query all rollapps currently registered in the hub",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRollappRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RollappAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryRollapp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [rollapp-id]",
		Short: "Query the rollapp associated with the specified rollapp-id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argRollappId := args[0]

			params := &types.QueryGetRollappRequest{
				RollappId: argRollappId,
			}

			res, err := queryClient.Rollapp(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd

}

func CmdQueryStateInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "state [rollapp-id]",
		Short: "Query the state associated with the specified rollapp-id and any other flags. If no flags are provided, return the latest state.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			argRollappId := args[0]

			flagSet := cmd.Flags()
			argIndex, err := flagSet.GetUint64(FlagStateIndex)
			if err != nil {
				return err
			}
			argHeight, err := flagSet.GetUint64(FlagRollappHeight)
			if err != nil {
				return err
			}
			argFinalized, err := flagSet.GetBool(FlagFinalized)
			if err != nil {
				return err
			}

			if (argHeight != 0 && argIndex != 0) || (argHeight != 0 && argFinalized) || (argIndex != 0 && argFinalized) {
				return status.Error(codes.InvalidArgument, fmt.Sprintf("only one flag can be use for %s, %s or %s", FlagStateIndex, FlagRollappHeight, FlagFinalized))
			}

			params := &types.QueryGetStateInfoRequest{
				RollappId: argRollappId,
				Index:     argIndex,
				Height:    argHeight,
				Finalized: argFinalized,
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.StateInfo(context.Background(), params)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().Uint64(FlagStateIndex, 0, "Use a specific state-index to query state-info at")
	cmd.Flags().Uint64(FlagRollappHeight, 0, "Use a specific height of the rollapp to query state-info at")
	cmd.Flags().Bool(FlagFinalized, false, "Indicates whether to return the latest finalized state")

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
