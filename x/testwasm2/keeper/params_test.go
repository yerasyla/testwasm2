package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "testwasm2/testutil/keeper"
	"testwasm2/x/testwasm2/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Testwasm2Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
