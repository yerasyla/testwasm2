package keeper

import (
	"testwasm2/x/testwasm2/types"
)

var _ types.QueryServer = Keeper{}
