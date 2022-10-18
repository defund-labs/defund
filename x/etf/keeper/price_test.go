package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/etf/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func (s *KeeperTestSuite) TestCreatePrice() {
	path := s.NewTransferPath()
	// value to set the remoteAddress, connection and port id to
	var connectionId string
	var portId string
	var accAddress string
	var fund types.Fund

	s.Run("CreateFundPrice", func() {
		fund = s.CreateTestFund()
		// We must create an ICA channel here on the broker chain with the test fund address as the owner
		connectionId, portId = s.CreateChannelICA(fund.Address, path)
		accAddress, found := s.GetDefundApp(s.chainA).ICAControllerKeeper.GetInterchainAccountAddress(s.chainA.GetContext(), connectionId, portId)
		s.Assert().True(found)
		atomCoin, osmoCoin, aktCoin := s.CreateTestTokens()
		// add them to an account balance
		s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccount.GetAddress(), sdk.NewCoins(atomCoin))
		s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccount.GetAddress(), sdk.NewCoins(osmoCoin))
		s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccount.GetAddress(), sdk.NewCoins(aktCoin))
		// create the fake balance query for fund
		s.CreateFundBalanceQuery(accAddress, []sdk.Coin{atomCoin, osmoCoin, aktCoin}, 1)
		s.CreatePoolQueries(fund)

		price, err := s.GetDefundApp(s.chainA).EtfKeeper.CreateFundPrice(s.chainA.GetContext(), fund.Symbol)
		s.Assert().NoError(err)
		s.Assert().Equal(price, sdk.NewCoin("uosmo", sdk.NewInt(44565793)))
	})

	s.Run("GetOwnershipSharesInFund", func() {
		atomCoin, osmoCoin, aktCoin := s.CreateTestTokens()
		// add them to an account balance
		s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccount.GetAddress(), sdk.NewCoins(atomCoin))
		s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccount.GetAddress(), sdk.NewCoins(osmoCoin))
		s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccount.GetAddress(), sdk.NewCoins(aktCoin))
		// create the fake balance query for fund
		s.CreateFundBalanceQuery(accAddress, []sdk.Coin{atomCoin, osmoCoin, aktCoin}, 1)
		s.CreatePoolQueries(fund)
		// create fund shares
		newShares := sdk.NewCoin(fund.Shares.Denom, sdk.NewInt(500000))

		ownership, err := s.GetDefundApp(s.chainA).EtfKeeper.GetOwnershipSharesInFund(s.chainA.GetContext(), fund, newShares)
		s.Assert().NoError(err)

		ret := sdk.Coins(ownership).IsEqual(sdk.NewCoins(sdk.NewCoin("uosmo", sdk.NewInt(5000000)), sdk.NewCoin("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", sdk.NewInt(5000000)), sdk.NewCoin("ibc/1480B8FD20AD5FCAE81EA87584D269547DD4D436843C1D20F15E00EB64743EF4", sdk.NewInt(5000000))))
		s.Assert().True(ret)
	})

	s.Run("GetAmountETFSharesForTokens", func() {
		atomCoin, osmoCoin, aktCoin := s.CreateTestTokens()
		// add them to an account balance
		s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccount.GetAddress(), sdk.NewCoins(atomCoin))
		s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccount.GetAddress(), sdk.NewCoins(osmoCoin))
		s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccount.GetAddress(), sdk.NewCoins(aktCoin))
		// create the fake balance query for fund
		s.CreateFundBalanceQuery(accAddress, []sdk.Coin{atomCoin, osmoCoin, aktCoin}, 1)
		s.CreatePoolQueries(fund)
		// create fund shares
		newShares := sdk.NewCoin(fund.BaseDenom, sdk.NewInt(44565793))

		ownership, err := s.GetDefundApp(s.chainA).EtfKeeper.GetAmountETFSharesForToken(s.chainA.GetContext(), fund, newShares)
		s.Assert().NoError(err)

		// make sure we have the amount of etf shares we want
		s.Assert().Equal(sdk.NewCoin(fund.Shares.Denom, sdk.NewInt(1000000)), ownership)
	})
}
