package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
	"github.com/stretchr/testify/require"
)

func (s *KeeperTestSuite) setup(ctx sdk.Context) (outctx sdk.Context, fund types.Fund, connectionId string, portId string) {
	path := s.NewTransferPath()
	s.Require().Equal(path.EndpointA.ChannelID, "channel-0")

	fund, connectionId, portId, _ = s.CreateTestFund(path)
	// Commit new block to store info
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	accAddress, found := s.GetDefundApp(s.chainA).ICAControllerKeeper.GetInterchainAccountAddress(ctx, connectionId, portId)
	s.Assert().True(found)
	atomCoin, osmoCoin, aktCoin := s.CreateTestTokens()
	// add them to an account balance
	err := s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(atomCoin))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(osmoCoin))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(aktCoin))
	s.Assert().NoError(err)
	s.CreateOsmosisBroker()
	// create the fake balance query for fund
	s.CreateFundBalanceQuery(accAddress, []sdk.Coin{atomCoin, osmoCoin, aktCoin}, 1)
	s.CreatePoolQueries(fund)

	outctx = ctx

	return outctx, fund, connectionId, portId
}

func (s *KeeperTestSuite) TestFundMsgServerCreate() {
	k := s.GetDefundApp(s.chainA).EtfKeeper
	ctx := s.chainA.GetContext()
	ctx, _, _, _ = s.setup(ctx)
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	srv := keeper.NewMsgServerImpl(k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateFund{
			Symbol:        fmt.Sprintf("%d", i),
			Creator:       creator,
			Holdings:      "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2:33:osmosis:1:spot,uosmo:34:osmosis:1:spot,ibc/1480B8FD20AD5FCAE81EA87584D269547DD4D436843C1D20F15E00EB64743EF4:33:osmosis:3:spot",
			BaseDenom:     "osmo",
			Rebalance:     10,
			StartingPrice: "10000000",
		}
		_, err := srv.CreateFund(wctx, expected)
		s.Assert().NoError(err)
	}
}

func (s *KeeperTestSuite) TestSharesMsgServerCreate() {
	k := s.GetDefundApp(s.chainA).EtfKeeper
	ctx := s.chainA.GetContext()
	ctx, _, _, _ = s.setup(ctx)
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	srv := keeper.NewMsgServerImpl(k)
	wctx := sdk.WrapSDKContext(ctx)

	fund, err := s.GetDefundApp(s.chainA).EtfKeeper.GetFundBySymbol(s.chainA.GetContext(), "test")
	s.Assert().NoError(err)
	// Commit new block to store info
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	atomCoin, osmoCoin, aktCoin := s.CreateTestTokens()
	// add them to an account balance
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(atomCoin))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(osmoCoin))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(aktCoin))
	s.Assert().NoError(err)
	s.CreatePoolQueries(fund)
	token := sdk.NewCoin(fund.BaseDenom.OnDefund, sdk.NewInt(10000000))
	err = s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), types.ModuleName, sdk.NewCoins(token))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(token))
	s.Assert().NoError(err)

	s.coordinator.CommitBlock(s.chainA, s.chainB)

	expected := &types.MsgCreate{
		Creator:          s.chainA.SenderAccounts[1].SenderAccount.GetAddress().String(),
		Fund:             fund.Symbol,
		TokenIn:          &token,
		Channel:          "channel-0",
		TimeoutHeight:    "0-50",
		TimeoutTimestamp: 0,
	}
	_, err = srv.Create(wctx, expected)
	require.NoError(s.T(), err)
}

func (s *KeeperTestSuite) TestSharesMsgServerRedeem() {
	k := s.GetDefundApp(s.chainA).EtfKeeper
	ctx := s.chainA.GetContext()
	ctx, _, _, _ = s.setup(ctx)
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	srv := keeper.NewMsgServerImpl(k)
	wctx := sdk.WrapSDKContext(ctx)

	fund, err := s.GetDefundApp(s.chainA).EtfKeeper.GetFundBySymbol(s.chainA.GetContext(), "test")
	s.Assert().NoError(err)
	// Commit new block to store info
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	atomCoin, osmoCoin, aktCoin := s.CreateTestTokens()
	// add them to an account balance
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(atomCoin))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(osmoCoin))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(aktCoin))
	s.Assert().NoError(err)
	s.CreatePoolQueries(fund)
	// create the etf shares we will redeem
	etfToken := sdk.NewCoin(fund.Shares.Denom, sdk.NewInt(224387))
	err = s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), types.ModuleName, sdk.NewCoins(etfToken))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(etfToken))
	s.Assert().NoError(err)
	// add the etf shares to the fund
	newShares := fund.Shares.Add(etfToken)
	fund.Shares = &newShares
	s.GetDefundApp(s.chainA).EtfKeeper.SetFund(s.chainA.GetContext(), fund)
	s.Assert().NoError(err)

	s.coordinator.CommitBlock(s.chainA, s.chainB)

	expectedRedeem := &types.MsgRedeem{
		Creator: s.chainA.SenderAccounts[1].SenderAccount.GetAddress().String(),
		Fund:    fund.Symbol,
		Amount:  &etfToken,
		Addresses: &types.AddressMap{
			OsmosisAddress: "osmo1234",
		},
	}

	_, err = srv.Redeem(wctx, expectedRedeem)
	require.NoError(s.T(), err)
}

func (s *KeeperTestSuite) TestSharesMsgServerEditFund() {}
