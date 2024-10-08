package keeper_test

import (
	"encoding/binary"
	"testing"
	"time"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/suite"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"defund/app"
	"defund/x/dex/amm"
	"defund/x/dex/keeper"
	"defund/x/dex/types"
)

type KeeperTestSuite struct {
	suite.Suite

	app       *app.App
	ctx       sdk.Context
	msgServer types.MsgServer
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	app := app.InitTestApp(true)
	s.app = app
	hdr := cmtproto.Header{
		Height: 1,
	}
	s.ctx = s.app.BaseApp.NewContext(true).WithBlockHeader(hdr)
	// Initialize params
	s.app.DexKeeper.SetParams(s.ctx, types.DefaultParams())
	s.msgServer = keeper.NewMsgServerImpl(s.app.DexKeeper)
}

// Below are just shortcuts to frequently-used functions.
func (s *KeeperTestSuite) getBalances(addr sdk.AccAddress) sdk.Coins {
	return s.app.BankKeeper.GetAllBalances(s.ctx, addr)
}

func (s *KeeperTestSuite) getBalance(addr sdk.AccAddress, denom string) sdk.Coin {
	return s.app.BankKeeper.GetBalance(s.ctx, addr, denom)
}

func (s *KeeperTestSuite) sendCoins(fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) {
	s.T().Helper()
	err := s.app.BankKeeper.SendCoins(s.ctx, fromAddr, toAddr, amt)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) nextBlock(commit bool) {
	s.T().Helper()
	s.app.ModuleManager.EndBlock(s.ctx)
	s.app.EndBlocker(s.ctx)
	if commit {
		s.app.Commit()
	}
	s.app.BeginBlocker(s.ctx)
}

// Below are useful helpers to write test code easily.
func (s *KeeperTestSuite) addr(addrNum int) sdk.AccAddress {
	addr := make(sdk.AccAddress, 20)
	binary.PutVarint(addr, int64(addrNum))
	return addr
}

func (s *KeeperTestSuite) fundAddr(addr sdk.AccAddress, amt sdk.Coins) {
	s.T().Helper()
	err := s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, amt)
	s.Require().NoError(err)
	err = s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, addr, amt)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) createPair(creator sdk.AccAddress, baseCoinDenom, quoteCoinDenom string, fund bool) types.Pair {
	s.T().Helper()
	if fund {
		s.fundAddr(creator, s.app.DexKeeper.GetPairCreationFee(s.ctx))
	}
	msg := types.NewMsgCreatePair(creator, baseCoinDenom, quoteCoinDenom)
	s.Require().NoError(msg.ValidateBasic())
	pair, err := s.app.DexKeeper.CreatePair(s.ctx, msg)
	s.Require().NoError(err)
	return pair
}

func (s *KeeperTestSuite) createPool(creator sdk.AccAddress, pairId uint64, depositCoins sdk.Coins, fund bool) types.Pool {
	s.T().Helper()
	if fund {
		s.fundAddr(creator, depositCoins.Add(s.app.DexKeeper.GetPoolCreationFee(s.ctx)...))
	}
	msg := types.NewMsgCreatePool(creator, pairId, depositCoins)
	s.Require().NoError(msg.ValidateBasic())
	pool, err := s.app.DexKeeper.CreatePool(s.ctx, msg)
	s.Require().NoError(err)
	return pool
}

func (s *KeeperTestSuite) createRangedPool(creator sdk.AccAddress, pairId uint64, depositCoins sdk.Coins, minPrice, maxPrice, initialPrice math.LegacyDec, fund bool) types.Pool {
	s.T().Helper()
	if fund {
		s.fundAddr(creator, depositCoins.Add(s.app.DexKeeper.GetPoolCreationFee(s.ctx)...))
	}
	msg := types.NewMsgCreateRangedPool(creator, pairId, depositCoins, minPrice, maxPrice, initialPrice)
	s.Require().NoError(msg.ValidateBasic())
	pool, err := s.app.DexKeeper.CreateRangedPool(s.ctx, msg)
	s.Require().NoError(err)
	return pool
}

func (s *KeeperTestSuite) deposit(depositor sdk.AccAddress, poolId uint64, depositCoins sdk.Coins, fund bool) types.DepositRequest {
	s.T().Helper()
	if fund {
		s.fundAddr(depositor, depositCoins)
	}
	req, err := s.app.DexKeeper.Deposit(s.ctx, types.NewMsgDeposit(depositor, poolId, depositCoins))
	s.Require().NoError(err)
	return req
}

func (s *KeeperTestSuite) withdraw(withdrawer sdk.AccAddress, poolId uint64, poolCoin sdk.Coin) types.WithdrawRequest {
	s.T().Helper()
	req, err := s.app.DexKeeper.Withdraw(s.ctx, types.NewMsgWithdraw(withdrawer, poolId, poolCoin))
	s.Require().NoError(err)
	return req
}

func (s *KeeperTestSuite) limitOrder(
	orderer sdk.AccAddress, pairId uint64, dir types.OrderDirection,
	price math.LegacyDec, amt math.Int, orderLifespan time.Duration, fund bool) types.Order {
	s.T().Helper()
	pair, found := s.app.DexKeeper.GetPair(s.ctx, pairId)
	s.Require().True(found)
	var ammDir amm.OrderDirection
	var offerCoinDenom, demandCoinDenom string
	switch dir {
	case types.OrderDirectionBuy:
		ammDir = amm.Buy
		offerCoinDenom, demandCoinDenom = pair.QuoteCoinDenom, pair.BaseCoinDenom
	case types.OrderDirectionSell:
		ammDir = amm.Sell
		offerCoinDenom, demandCoinDenom = pair.BaseCoinDenom, pair.QuoteCoinDenom
	}
	offerCoin := sdk.NewCoin(offerCoinDenom, amm.OfferCoinAmount(ammDir, price, amt))
	if fund {
		s.fundAddr(orderer, sdk.NewCoins(offerCoin))
	}
	msg := types.NewMsgLimitOrder(
		orderer, pairId, dir, offerCoin, demandCoinDenom,
		price, amt, orderLifespan)
	s.Require().NoError(msg.ValidateBasic())
	req, err := s.app.DexKeeper.LimitOrder(s.ctx, msg)
	s.Require().NoError(err)
	return req
}

func (s *KeeperTestSuite) buyLimitOrder(
	orderer sdk.AccAddress, pairId uint64, price math.LegacyDec,
	amt math.Int, orderLifespan time.Duration, fund bool) types.Order {
	s.T().Helper()
	return s.limitOrder(
		orderer, pairId, types.OrderDirectionBuy, price, amt, orderLifespan, fund)
}

func (s *KeeperTestSuite) sellLimitOrder(
	orderer sdk.AccAddress, pairId uint64, price math.LegacyDec,
	amt math.Int, orderLifespan time.Duration, fund bool) types.Order {
	s.T().Helper()
	return s.limitOrder(
		orderer, pairId, types.OrderDirectionSell, price, amt, orderLifespan, fund)
}

func (s *KeeperTestSuite) marketOrder(
	orderer sdk.AccAddress, pairId uint64, dir types.OrderDirection,
	amt math.Int, orderLifespan time.Duration, fund bool) types.Order {
	s.T().Helper()
	pair, found := s.app.DexKeeper.GetPair(s.ctx, pairId)
	s.Require().True(found)
	s.Require().NotNil(pair.LastPrice)
	lastPrice := *pair.LastPrice
	var offerCoin sdk.Coin
	var demandCoinDenom string
	switch dir {
	case types.OrderDirectionBuy:
		maxPrice := lastPrice.Mul(math.LegacyOneDec().Add(s.app.DexKeeper.GetMaxPriceLimitRatio(s.ctx)))
		offerCoin = sdk.NewCoin(pair.QuoteCoinDenom, amm.OfferCoinAmount(amm.Buy, maxPrice, amt))
		demandCoinDenom = pair.BaseCoinDenom
	case types.OrderDirectionSell:
		offerCoin = sdk.NewCoin(pair.BaseCoinDenom, amt)
		demandCoinDenom = pair.QuoteCoinDenom
	}
	if fund {
		s.fundAddr(orderer, sdk.NewCoins(offerCoin))
	}
	msg := types.NewMsgMarketOrder(
		orderer, pairId, dir, offerCoin, demandCoinDenom,
		amt, orderLifespan)
	s.Require().NoError(msg.ValidateBasic())
	req, err := s.app.DexKeeper.MarketOrder(s.ctx, msg)
	s.Require().NoError(err)
	return req
}

func (s *KeeperTestSuite) buyMarketOrder(
	orderer sdk.AccAddress, pairId uint64,
	amt math.Int, orderLifespan time.Duration, fund bool) types.Order {
	s.T().Helper()
	return s.marketOrder(
		orderer, pairId, types.OrderDirectionBuy, amt, orderLifespan, fund)
}

func (s *KeeperTestSuite) sellMarketOrder(
	orderer sdk.AccAddress, pairId uint64,
	amt math.Int, orderLifespan time.Duration, fund bool) types.Order {
	s.T().Helper()
	return s.marketOrder(
		orderer, pairId, types.OrderDirectionSell, amt, orderLifespan, fund)
}

func (s *KeeperTestSuite) mmOrder(
	orderer sdk.AccAddress, pairId uint64, dir types.OrderDirection,
	price math.LegacyDec, amt math.Int, orderLifespan time.Duration, fund bool) types.Order {
	s.T().Helper()
	pair, found := s.app.DexKeeper.GetPair(s.ctx, pairId)
	s.Require().True(found)
	var ammDir amm.OrderDirection
	var offerCoinDenom, demandCoinDenom string
	switch dir {
	case types.OrderDirectionBuy:
		ammDir = amm.Buy
		offerCoinDenom, demandCoinDenom = pair.QuoteCoinDenom, pair.BaseCoinDenom
	case types.OrderDirectionSell:
		ammDir = amm.Sell
		offerCoinDenom, demandCoinDenom = pair.BaseCoinDenom, pair.QuoteCoinDenom
	}
	offerCoin := sdk.NewCoin(offerCoinDenom, amm.OfferCoinAmount(ammDir, price, amt))
	if fund {
		s.fundAddr(orderer, sdk.NewCoins(offerCoin))
	}
	msg := types.NewMsgMMOrder(
		orderer, pairId, dir, offerCoin, demandCoinDenom,
		price, amt, orderLifespan)
	s.Require().NoError(msg.ValidateBasic())
	req, err := s.app.DexKeeper.MMOrder(s.ctx, msg)
	s.Require().NoError(err)
	return req
}

// nolint
func (s *KeeperTestSuite) cancelOrder(orderer sdk.AccAddress, pairId, orderId uint64) {
	s.T().Helper()
	err := s.app.DexKeeper.CancelOrder(s.ctx, types.NewMsgCancelOrder(orderer, pairId, orderId))
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) cancelAllOrders(orderer sdk.AccAddress, pairIds []uint64) {
	s.T().Helper()
	err := s.app.DexKeeper.CancelAllOrders(s.ctx, types.NewMsgCancelAllOrders(orderer, pairIds))
	s.Require().NoError(err)
}

func coinEq(exp, got sdk.Coin) (bool, string, string, string) {
	return exp.IsEqual(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func coinsEq(exp, got sdk.Coins) (bool, string, string, string) {
	return exp.Equal(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func intEq(exp, got math.Int) (bool, string, string, string) {
	return exp.Equal(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func decEq(exp, got math.LegacyDec) (bool, string, string, string) {
	return exp.Equal(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func newInt(i int64) math.Int {
	return math.NewInt(i)
}
