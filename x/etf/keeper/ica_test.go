package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/etf/types"
)

func (s *KeeperTestSuite) TestCallbackActions() {

	path := s.NewTransferPath()
	s.Require().Equal(path.EndpointA.ChannelID, "channel-0")

	fund, _, _, _ := s.CreateTestFund(path)
	// Commit new block to store info
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	// We must create an ICA channel here on the broker chain with the test fund address as the owner
	connectionId, portId := s.CreateChannelICA(fund.Address, path)
	accAddress, found := s.GetDefundApp(s.chainA).ICAControllerKeeper.GetInterchainAccountAddress(s.chainA.GetContext(), connectionId, portId)
	s.Assert().True(found)
	atomCoin, osmoCoin, aktCoin := s.CreateTestTokens()
	// add them to an account balance
	err := s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(atomCoin))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(osmoCoin))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(aktCoin))
	s.Assert().NoError(err)
	// create the fake balance query for fund
	s.CreateFundBalanceQuery(accAddress, []sdk.Coin{atomCoin, osmoCoin, aktCoin}, 1)
	s.CreatePoolQueries(fund)

	s.Run("OnAcknowledgementPacketSuccess", func() {})

	s.Run("OnAcknowledgementPacketFailure", func() {})

	s.Run("OnAcknowledgementPacket", func() {})
}
