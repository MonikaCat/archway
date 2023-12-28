package genmsg_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"

	"github.com/MonikaCat/archway/v5/app"
	e2eTesting "github.com/MonikaCat/archway/v5/e2e/testing"
	"github.com/MonikaCat/archway/v5/x/genmsg"
	"github.com/MonikaCat/archway/v5/x/genmsg/types"
)

func TestIntegration(t *testing.T) {
	recvAddr := sdk.AccAddress("recv")
	chain := e2eTesting.NewTestChain(t, 1, e2eTesting.WithDummyTestAddress(), e2eTesting.WithGenDefaultCoinBalance("100000000000000000000000000000000000"), e2eTesting.TestChainGenesisOption(func(cdc codec.Codec, genesis app.GenesisState) {
		testMsg := &banktypes.MsgSend{
			FromAddress: e2eTesting.TestAccountAddr.String(),
			ToAddress:   recvAddr.String(),
			Amount:      sdk.NewCoins(sdk.NewInt64Coin("stake", 1000)),
		}
		anyMsg, err := codectypes.NewAnyWithValue(testMsg)
		require.NoError(t, err)
		genesis[genmsg.ModuleName] = cdc.MustMarshalJSON(&types.GenesisState{Messages: []*codectypes.Any{anyMsg}})
	}))
	bankQuery := banktypes.NewQueryClient(chain.Client())
	resp, err := bankQuery.Balance(sdk.WrapSDKContext(chain.GetContext()), &banktypes.QueryBalanceRequest{
		Address: recvAddr.String(),
		Denom:   "stake",
	})

	require.NoError(t, err)
	t.Log(resp)
}
