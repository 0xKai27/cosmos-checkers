package cosmoscheckers_test

import (
	"testing"

	keepertest "github.com/OxKai27/cosmos-checkers/testutil/keeper"
	"github.com/OxKai27/cosmos-checkers/testutil/nullify"
	"github.com/OxKai27/cosmos-checkers/x/cosmoscheckers"
	"github.com/OxKai27/cosmos-checkers/x/cosmoscheckers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		NextGame: &types.NextGame{
			IdValue: 11,
		},
		StoredGameList: []types.StoredGame{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmoscheckersKeeper(t)
	cosmoscheckers.InitGenesis(ctx, *k, genesisState)
	got := cosmoscheckers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.NextGame, got.NextGame)
	require.ElementsMatch(t, genesisState.StoredGameList, got.StoredGameList)
	// this line is used by starport scaffolding # genesis/test/assert
}
