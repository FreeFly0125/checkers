package keeper

import (
	"myCheckers/x/mycheckers/types"
)

var _ types.QueryServer = Keeper{}
