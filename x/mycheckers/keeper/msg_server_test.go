package keeper_test

import (
	"context"
	"testing"

	keepertest "myCheckers/testutil/keeper"
	"myCheckers/x/mycheckers/keeper"
	"myCheckers/x/mycheckers/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MycheckersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	require.NotNil(t, msgServer)
	require.NotNil(t, context)
}