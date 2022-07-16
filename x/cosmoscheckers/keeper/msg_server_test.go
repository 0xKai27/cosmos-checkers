package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/OxKai27/cosmos-checkers/testutil/keeper"
	"github.com/OxKai27/cosmos-checkers/x/cosmoscheckers/keeper"
	"github.com/OxKai27/cosmos-checkers/x/cosmoscheckers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmoscheckersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
