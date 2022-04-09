package keeper

import (
	"github.com/cosmonaut/documentservice/x/documentservice/types"
)

var _ types.QueryServer = Keeper{}
