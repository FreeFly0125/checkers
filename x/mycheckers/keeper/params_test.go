package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "myCheckers/testutil/keeper"
	"myCheckers/x/mycheckers/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MycheckersKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
