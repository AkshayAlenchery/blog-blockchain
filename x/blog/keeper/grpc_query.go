package keeper

import (
	"github.com/akshayalenchery/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
