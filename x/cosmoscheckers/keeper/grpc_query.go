package keeper

import (
	"github.com/OxKai27/cosmos-checkers/x/cosmoscheckers/types"
)

var _ types.QueryServer = Keeper{}
