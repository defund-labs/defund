package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func (s *KeeperTestSuite) TestCreatePrice() {
	path := s.NewTransferPath()
	s.CreateTestFund(path)
	fund, err := s.GetDefundApp(s.chainA).EtfKeeper.GetFundBySymbol(s.chainA.GetContext(), "test")
	s.Assert().NoError(err)
	s.CreatePoolQueries(fund)

	s.Run("CreateFundPrice", func() {

		price, err := s.GetDefundApp(s.chainA).EtfKeeper.CreateFundPrice(s.chainA.GetContext(), fund.Symbol)
		s.Assert().NoError(err)
		s.Assert().Equal(price, sdk.NewCoin("uosmo", sdk.NewInt(44565793)))
	})

	s.Run("GetOwnershipSharesInFund", func() {
		// create fund shares
		newShares := sdk.NewCoin(fund.Shares.Denom, sdk.NewInt(500000))

		ownership, err := s.GetDefundApp(s.chainA).EtfKeeper.GetOwnershipSharesInFund(s.chainA.GetContext(), fund, newShares)
		s.Assert().NoError(err)

		ret := sdk.Coins(ownership).IsEqual(sdk.NewCoins(sdk.NewCoin("uosmo", sdk.NewInt(5000000)), sdk.NewCoin("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", sdk.NewInt(5000000)), sdk.NewCoin("ibc/1480B8FD20AD5FCAE81EA87584D269547DD4D436843C1D20F15E00EB64743EF4", sdk.NewInt(5000000))))
		s.Assert().True(ret)
	})

	s.Run("GetAmountETFSharesForTokens", func() {
		// create fund shares
		newShares := sdk.NewCoin(fund.BaseDenom, sdk.NewInt(44565793))

		ownership, err := s.GetDefundApp(s.chainA).EtfKeeper.GetAmountETFSharesForToken(s.chainA.GetContext(), fund, newShares)
		s.Assert().NoError(err)

		// make sure we have the amount of etf shares we want
		s.Assert().Equal(sdk.NewCoin(fund.Shares.Denom, sdk.NewInt(1000000)), ownership)
	})
}
