package cli

import "github.com/spf13/pflag"

const (
	FlagStateIndex          = "index"
	FlagRollappHeight       = "rollapp-height"
	FlagFinalized           = "finalized"
	FlagGenesisAccountsPath = "genesis-accounts-path"
)

// FlagSetCreateRollapp returns flags for creating gauges.
func FlagSetCreateRollapp() *pflag.FlagSet {
	fs := pflag.NewFlagSet("", pflag.ContinueOnError)

	fs.String(FlagGenesisAccountsPath, "", "path to a json file containing genesis accounts")
	return fs
}
