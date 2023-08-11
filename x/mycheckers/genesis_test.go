package mycheckers_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "myCheckers/testutil/keeper"
	"myCheckers/testutil/nullify"
	"myCheckers/x/mycheckers"
	"myCheckers/x/mycheckers/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MycheckersKeeper(t)
	mycheckers.InitGenesis(ctx, *k, genesisState)
	got := mycheckers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
