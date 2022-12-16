package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	brokertypes "github.com/defund-labs/defund/x/broker/types"
	"github.com/defund-labs/defund/x/etf/types"
)

func (s *KeeperTestSuite) TestBrokerCallbacks() {

	path := s.NewTransferPath()
	s.Require().Equal(path.EndpointA.ChannelID, "channel-0")

	fund, connectionId, portId, _ := s.CreateTestFund(path)
	// Commit new block to store info
	s.coordinator.CommitBlock(s.chainA, s.chainB)
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

	s.Run("OnRedeemSuccess", func() {
		// mock data for packet
		data := []byte{18, 219, 113, 12, 103, 107, 95, 216, 56, 143, 130, 159, 113, 176, 79, 128, 79, 214, 45, 220, 115, 169, 192, 84, 181, 42, 226, 211, 113, 13, 252, 109}
		// create a mock packet
		packet := channeltypes.NewPacket(data, 1, portId, "channel-1", "icahost", "channel-1", clienttypes.NewHeight(0, 10), 0)
		amount := sdk.NewCoin(fund.Shares.Denom, sdk.NewInt(1000000))
		err := s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), "etf", sdk.NewCoins(amount))
		s.Assert().NoError(err)
		// check that the module account received etf shares
		moduleAccount := s.GetDefundApp(s.chainA).AccountKeeper.GetModuleAddress("etf")
		balance := s.GetDefundApp(s.chainA).BankKeeper.GetAllBalances(s.chainA.GetContext(), moduleAccount)
		s.Assert().Contains(balance, amount)
		redeem := brokertypes.Redeem{
			Creator: s.chainA.SenderAccount.GetAddress().String(),
			Fund:    fund.Symbol,
			Amount:  &amount,
		}
		err = s.GetDefundApp(s.chainA).EtfKeeper.OnRedeemFailure(s.chainA.GetContext(), packet, redeem)
		s.Assert().NoError(err)
		// confirm that the balance of etf shares no longer included in balance
		balance = s.GetDefundApp(s.chainA).BankKeeper.GetAllBalances(s.chainA.GetContext(), moduleAccount)
		s.Assert().NotContains(balance, amount)
	})

	s.Run("OnRedeemFailure", func() {
		// mock data for packet
		data := []byte{18, 219, 113, 12, 103, 107, 95, 216, 56, 143, 130, 159, 113, 176, 79, 128, 79, 214, 45, 220, 115, 169, 192, 84, 181, 42, 226, 211, 113, 13, 252, 109}
		// create a mock packet
		packet := channeltypes.NewPacket(data, 1, portId, "channel-1", "icahost", "channel-1", clienttypes.NewHeight(0, 10), 0)
		amount := sdk.NewCoin(fund.Shares.Denom, sdk.NewInt(1000000))
		err := s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), "etf", sdk.NewCoins(amount))
		s.Assert().NoError(err)
		// check that the module account received etf shares
		moduleAccount := s.GetDefundApp(s.chainA).AccountKeeper.GetModuleAddress("etf")
		balance := s.GetDefundApp(s.chainA).BankKeeper.GetAllBalances(s.chainA.GetContext(), moduleAccount)
		s.Assert().Contains(balance, amount)
		redeem := brokertypes.Redeem{
			Creator: s.chainA.SenderAccount.GetAddress().String(),
			Fund:    fund.Symbol,
			Amount:  &amount,
		}
		addr, err := sdk.AccAddressFromBech32(s.chainA.SenderAccount.GetAddress().String())
		s.Assert().NoError(err)
		balance = s.GetDefundApp(s.chainA).BankKeeper.GetAllBalances(s.chainA.GetContext(), addr)
		// account starts off with 1000000etf/test
		s.Assert().Contains(balance, amount)
		err = s.GetDefundApp(s.chainA).EtfKeeper.OnRedeemFailure(s.chainA.GetContext(), packet, redeem)
		s.Assert().NoError(err)
		// confirm that the balance of etf shares no longer included in module balance
		balance = s.GetDefundApp(s.chainA).BankKeeper.GetAllBalances(s.chainA.GetContext(), moduleAccount)
		s.Assert().NotContains(balance, amount)
		// confirm that the balance of etf shares was sent back to the redeemer
		balance = s.GetDefundApp(s.chainA).BankKeeper.GetAllBalances(s.chainA.GetContext(), addr)
		// account should now have 2000000etf/test
		s.Assert().Contains(balance, sdk.NewCoin(fund.Shares.Denom, amount.Amount.Add(sdk.NewInt(1000000))))
	})

	s.Run("OnRebalanceSuccess", func() {
		rebalance := brokertypes.Rebalance{
			Id:     "channel-1-1",
			Fund:   fund.Symbol,
			Height: 1,
			Broker: "osmosis",
		}
		// set the rebalance
		s.GetDefundApp(s.chainA).BrokerKeeper.SetRebalance(s.chainA.GetContext(), rebalance)
		fund, found := s.GetDefundApp(s.chainA).EtfKeeper.GetFund(s.chainA.GetContext(), fund.Symbol)
		s.Assert().True(found)
		s.Assert().Equal(fund.LastRebalanceHeight, int64(0))
		err := s.GetDefundApp(s.chainA).EtfKeeper.OnRebalanceSuccess(s.chainA.GetContext(), rebalance, fund.Symbol)
		s.Assert().NoError(err)
		fund, found = s.GetDefundApp(s.chainA).EtfKeeper.GetFund(s.chainA.GetContext(), fund.Symbol)
		s.Assert().True(found)
		// check to make sure fund last rebelance height was updated
		s.Assert().Equal(fund.LastRebalanceHeight, int64(25))
		// ensure the rebalance store was deleted
		_, found = s.GetDefundApp(s.chainA).BrokerKeeper.GetRebalance(s.chainA.GetContext(), rebalance.Id)
		s.Assert().False(found)
	})

	s.Run("OnRebalanceFailure", func() {
		rebalance := brokertypes.Rebalance{
			Id:     "channel-1-1",
			Fund:   fund.Symbol,
			Height: 1,
			Broker: "osmosis",
		}
		// set the rebalance
		s.GetDefundApp(s.chainA).BrokerKeeper.SetRebalance(s.chainA.GetContext(), rebalance)
		err := s.GetDefundApp(s.chainA).EtfKeeper.OnRebalanceFailure(s.chainA.GetContext(), rebalance, fund.Symbol)
		s.Assert().NoError(err)
		// check to make sure fund last rebelance height was not updated
		s.Assert().Equal(fund.LastRebalanceHeight, int64(0))
		// ensure the rebalance store was deleted
		_, found = s.GetDefundApp(s.chainA).BrokerKeeper.GetRebalance(s.chainA.GetContext(), rebalance.Id)
		s.Assert().False(found)
	})
}
