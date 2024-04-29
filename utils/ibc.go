package utils

import (
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
)

const (
	ibcPort = "transfer"
)

func GetForeignIBCDenom(channelId, denom string) string {
	return GetForeignDenomTrace(channelId, denom).IBCDenom()
}

func GetForeignDenomTrace(channelId, denom string) transfertypes.DenomTrace {
	sourcePrefix := transfertypes.GetDenomPrefix(ibcPort, channelId)
	// NOTE: sourcePrefix contains the trailing "/"
	prefixedDenom := sourcePrefix + denom
	// construct the denomination trace from the full raw denomination
	return transfertypes.ParseDenomTrace(prefixedDenom)
}
