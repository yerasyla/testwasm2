package testwasm2_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "testwasm2/testutil/keeper"
	"testwasm2/testutil/nullify"
	"testwasm2/x/testwasm2"
	"testwasm2/x/testwasm2/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Testwasm2Keeper(t)
	testwasm2.InitGenesis(ctx, *k, genesisState)
	got := testwasm2.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
