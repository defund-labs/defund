package keeper_test

import (
	"fmt"
	"math/rand"
	"time"

	math "cosmossdk.io/math"
	store "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	utils "defund/types"
	"defund/x/dex/amm"
	dex "defund/x/dex/module"
	"defund/x/dex/types"

	_ "github.com/stretchr/testify/suite"
)

func (s *KeeperTestSuite) TestLimitOrder() {
	// Create a denom1/denom2 pair and set last price to 1.0
	pair1 := s.createPair(s.addr(0), "denom1", "denom2", true)
	lastPrice := utils.ParseDec("1.0")
	pair1.LastPrice = &lastPrice
	s.app.DexKeeper.SetPair(s.ctx, pair1)

	// denom2/denom1 pair doesn't have last price
	pair2 := s.createPair(s.addr(0), "denom2", "denom1", true)

	orderer := s.addr(1)
	s.fundAddr(orderer, utils.ParseCoins("1000000000denom1,1000000000denom2"))

	for _, tc := range []struct {
		name        string
		msg         *types.MsgLimitOrder
		expectedErr string
	}{
		{
			"happy case",
			types.NewMsgLimitOrder(
				orderer, pair1.Id, types.OrderDirectionBuy, utils.ParseCoin("1000000denom2"), "denom1",
				utils.ParseDec("1.0"), newInt(1000000), 0),
			"",
		},
		{
			"wrong offer coin and demand coin denom",
			types.NewMsgLimitOrder(
				orderer, pair1.Id, types.OrderDirectionBuy, utils.ParseCoin("1000000denom1"), "denom2",
				utils.ParseDec("1.0"), newInt(1000000), 0),
			"denom pair (denom2, denom1) != (denom1, denom2): wrong denom pair",
		},
		{
			"correct offer coin and demand coin denom",
			types.NewMsgLimitOrder(
				orderer, pair2.Id, types.OrderDirectionBuy, utils.ParseCoin("1000000denom1"), "denom2",
				utils.ParseDec("1.0"), newInt(1000000), 0),
			"",
		},
		{
			"price not fit in ticks",
			types.NewMsgLimitOrder(
				orderer, pair1.Id, types.OrderDirectionSell, utils.ParseCoin("1000000denom1"), "denom2",
				utils.ParseDec("1.0005"), newInt(1000000), 0),
			"",
		},
		{
			"too long order lifespan",
			types.NewMsgLimitOrder(
				orderer, pair1.Id, types.OrderDirectionSell, utils.ParseCoin("1000000denom1"), "denom2",
				utils.ParseDec("1.0"), newInt(1000000), 48*time.Hour),
			"48h0m0s is longer than 24h0m0s: order lifespan is too long",
		},
		{
			"pair not found",
			types.NewMsgLimitOrder(
				orderer, 3, types.OrderDirectionBuy, utils.ParseCoin("1000000denom1"), "denom2",
				utils.ParseDec("1.0"), newInt(1000000), 0),
			"pair 3 not found: not found",
		},
		{
			"price out of lower limit",
			types.NewMsgLimitOrder(
				orderer, pair1.Id, types.OrderDirectionBuy, utils.ParseCoin("1000000denom2"), "denom1",
				utils.ParseDec("0.8"), newInt(1000000), 0),
			"0.800000000000000000 is lower than 0.900000000000000000: price out of range limit",
		},
		{
			"price out of upper limit",
			types.NewMsgLimitOrder(
				orderer, pair1.Id, types.OrderDirectionBuy, utils.ParseCoin("2000000denom2"), "denom1",
				utils.ParseDec("1.2"), newInt(1000000), 0),
			"1.200000000000000000 is higher than 1.100000000000000000: price out of range limit",
		},
		{
			"no price limit without last price",
			types.NewMsgLimitOrder(
				orderer, pair2.Id, types.OrderDirectionSell, utils.ParseCoin("1000000denom2"), "denom1",
				utils.ParseDec("100.0"), newInt(1000000), 0),
			"",
		},
	} {
		s.Run(tc.name, func() {
			// The msg is valid, but may cause an error when it's being handled in the msg server.
			s.Require().NoError(tc.msg.ValidateBasic())
			req, err := s.app.DexKeeper.LimitOrder(s.ctx, tc.msg)
			if tc.expectedErr == "" {
				s.Require().NoError(err)
				switch tc.msg.Direction {
				case types.OrderDirectionBuy:
					s.Require().True(req.Price.LTE(tc.msg.Price))
				case types.OrderDirectionSell:
					s.Require().True(req.Price.GTE(tc.msg.Price))
				}
			} else {
				s.Require().EqualError(err, tc.expectedErr)
			}
		})
	}
}

func (s *KeeperTestSuite) TestLimitOrderInsufficientOfferCoin() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	orderer := s.addr(1)
	s.fundAddr(orderer, utils.ParseCoins("1000000denom2"))
	_, err := s.app.DexKeeper.LimitOrder(s.ctx, types.NewMsgLimitOrder(
		orderer, pair.Id, types.OrderDirectionBuy, utils.ParseCoin("1000001denom2"), "denom1",
		utils.ParseDec("1.0"), math.NewInt(1000000), 0))
	s.Require().ErrorIs(err, sdkerrors.ErrInsufficientFunds)

	s.fundAddr(orderer, utils.ParseCoins("1000000denom1"))
	_, err = s.app.DexKeeper.LimitOrder(s.ctx, types.NewMsgLimitOrder(
		orderer, pair.Id, types.OrderDirectionSell, utils.ParseCoin("1000001denom1"), "denom2",
		utils.ParseDec("1.0"), math.NewInt(1000000), 0))
	s.Require().ErrorIs(err, sdkerrors.ErrInsufficientFunds)
}

func (s *KeeperTestSuite) TestLimitOrderRefund() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	orderer := s.addr(1)
	s.fundAddr(orderer, utils.ParseCoins("1000000000denom1,1000000000denom2"))

	for _, tc := range []struct {
		msg          *types.MsgLimitOrder
		refundedCoin sdk.Coin
	}{
		{
			types.NewMsgLimitOrder(
				orderer, pair.Id, types.OrderDirectionBuy, utils.ParseCoin("1000000denom2"), "denom1",
				utils.ParseDec("1.0"), newInt(1000000), 0),
			utils.ParseCoin("0denom2"),
		},
		{
			types.NewMsgLimitOrder(
				orderer, pair.Id, types.OrderDirectionBuy, utils.ParseCoin("1000000denom2"), "denom1",
				utils.ParseDec("1.0"), newInt(10000), 0),
			utils.ParseCoin("990000denom2"),
		},
		{
			types.NewMsgLimitOrder(
				orderer, pair.Id, types.OrderDirectionBuy, utils.ParseCoin("1000denom2"), "denom1",
				utils.ParseDec("0.9999"), newInt(1000), 0),
			utils.ParseCoin("0denom2"),
		},
		{
			types.NewMsgLimitOrder(
				orderer, pair.Id, types.OrderDirectionBuy, utils.ParseCoin("102denom2"), "denom1",
				utils.ParseDec("1.001"), newInt(100), 0),
			utils.ParseCoin("1denom2"),
		},
		{
			types.NewMsgLimitOrder(
				orderer, pair.Id, types.OrderDirectionSell, utils.ParseCoin("1000denom1"), "denom2",
				utils.ParseDec("1.100"), newInt(1000), 0),
			utils.ParseCoin("0denom1"),
		},
		{
			types.NewMsgLimitOrder(
				orderer, pair.Id, types.OrderDirectionSell, utils.ParseCoin("1000denom1"), "denom2",
				utils.ParseDec("1.100"), newInt(100), 0),
			utils.ParseCoin("900denom1"),
		},
	} {
		s.Run("", func() {
			s.Require().NoError(tc.msg.ValidateBasic())

			balanceBefore := s.getBalance(orderer, tc.msg.OfferCoin.Denom)
			_, err := s.app.DexKeeper.LimitOrder(s.ctx, tc.msg)
			s.Require().NoError(err)

			balanceAfter := s.getBalance(orderer, tc.msg.OfferCoin.Denom)

			refundedCoin := balanceAfter.Sub(balanceBefore.Sub(tc.msg.OfferCoin))
			s.Require().True(coinEq(tc.refundedCoin, refundedCoin))
		})
	}
}

func (s *KeeperTestSuite) TestMarketOrder() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	// When there is no last price in the pair, only limit orders can be made.
	// These two orders will be matched.
	s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("1.0"), math.NewInt(10000), 0, true)
	s.sellLimitOrder(s.addr(2), pair.Id, utils.ParseDec("1.0"), math.NewInt(10000), 0, true)
	s.nextBlock(false)

	// Now users can make market orders.
	// In this case, addr(3) user's order takes higher priority than addr(4) user's,
	// because market buy orders have 10% higher price than the last price(1.0).
	s.buyMarketOrder(s.addr(3), pair.Id, math.NewInt(10000), 0, true)
	s.buyLimitOrder(s.addr(4), pair.Id, utils.ParseDec("1.08"), math.NewInt(10000), 0, true)
	s.sellLimitOrder(s.addr(5), pair.Id, utils.ParseDec("1.07"), math.NewInt(10000), 0, true)
	s.nextBlock(false)

	// Check the result.
	s.Require().True(coinEq(utils.ParseCoin("10000denom1"), s.getBalance(s.addr(3), "denom1")))
	s.Require().True(coinsEq(utils.ParseCoins("10800denom2"), s.getBalances(s.addr(4))))
}

func (s *KeeperTestSuite) TestMarketOrderInsufficientOfferCoin() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	orderer := s.addr(1)
	s.fundAddr(orderer, utils.ParseCoins("1000000denom2"))
	_, err := s.app.DexKeeper.MarketOrder(s.ctx, types.NewMsgMarketOrder(
		orderer, pair.Id, types.OrderDirectionBuy, utils.ParseCoin("1000001denom2"), "denom1",
		math.NewInt(1000000), 0))
	s.Require().ErrorIs(err, sdkerrors.ErrInsufficientFunds)

	s.fundAddr(orderer, utils.ParseCoins("1000000denom1"))
	_, err = s.app.DexKeeper.MarketOrder(s.ctx, types.NewMsgMarketOrder(
		orderer, pair.Id, types.OrderDirectionSell, utils.ParseCoin("1000001denom1"), "denom2",
		math.NewInt(1000000), 0))
	s.Require().ErrorIs(err, sdkerrors.ErrInsufficientFunds)
}

func (s *KeeperTestSuite) TestMarketOrderRefund() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	p := utils.ParseDec("1.0")
	pair.LastPrice = &p
	s.app.DexKeeper.SetPair(s.ctx, pair)
	orderer := s.addr(1)
	s.fundAddr(orderer, utils.ParseCoins("1000000000denom1,1000000000denom2"))

	for _, tc := range []struct {
		msg          *types.MsgMarketOrder
		refundedCoin sdk.Coin
	}{
		{
			types.NewMsgMarketOrder(
				orderer, pair.Id, types.OrderDirectionBuy, utils.ParseCoin("1100000denom2"), "denom1",
				newInt(1000000), 0),
			utils.ParseCoin("0denom2"),
		},
		{
			types.NewMsgMarketOrder(
				orderer, pair.Id, types.OrderDirectionBuy, utils.ParseCoin("1000000denom2"), "denom1",
				newInt(10000), 0),
			utils.ParseCoin("989000denom2"),
		},
		{
			types.NewMsgMarketOrder(
				orderer, pair.Id, types.OrderDirectionSell, utils.ParseCoin("1000000denom1"), "denom2",
				newInt(10000), 0),
			utils.ParseCoin("990000denom1"),
		},
	} {
		s.Run("", func() {
			s.Require().NoError(tc.msg.ValidateBasic())

			balanceBefore := s.getBalance(orderer, tc.msg.OfferCoin.Denom)
			_, err := s.app.DexKeeper.MarketOrder(s.ctx, tc.msg)
			s.Require().NoError(err)

			balanceAfter := s.getBalance(orderer, tc.msg.OfferCoin.Denom)

			refundedCoin := balanceAfter.Sub(balanceBefore.Sub(tc.msg.OfferCoin))
			s.Require().True(coinEq(tc.refundedCoin, refundedCoin))
		})
	}
}

func (s *KeeperTestSuite) TestMarketOrderWithNoLastPrice() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	s.Require().Nil(pair.LastPrice)
	offerCoin := utils.ParseCoin("10000denom2")
	s.fundAddr(s.addr(1), sdk.NewCoins(offerCoin))
	msg := types.NewMsgMarketOrder(
		s.addr(1), pair.Id, types.OrderDirectionBuy, offerCoin, "denom1", math.NewInt(10000), 0)
	_, err := s.app.DexKeeper.MarketOrder(s.ctx, msg)
	s.Require().ErrorIs(err, types.ErrNoLastPrice)
}

func (s *KeeperTestSuite) TestSingleOrderNoMatch() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	order := s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("1.0"), math.NewInt(1000000), 10*time.Second, true)
	// Execute matching
	s.nextBlock(false)

	order, found := s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	s.Require().True(found)
	s.Require().Equal(types.OrderStatusNotMatched, order.Status)

	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(10 * time.Second))
	// Expire the order, here BeginBlocker is not called to check
	// the request's changed status
	s.nextBlock(false)

	order, _ = s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	s.Require().Equal(types.OrderStatusExpired, order.Status)

	s.Require().True(coinsEq(utils.ParseCoins("1000000denom2"), s.getBalances(s.addr(1))))
}

func (s *KeeperTestSuite) TestTwoOrderExactMatch() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	req1 := s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("1.0"), newInt(10000), time.Hour, true)
	req2 := s.sellLimitOrder(s.addr(2), pair.Id, utils.ParseDec("1.0"), newInt(10000), time.Hour, true)
	s.nextBlock(false)

	req1, _ = s.app.DexKeeper.GetOrder(s.ctx, req1.PairId, req1.Id)
	s.Require().Equal(types.OrderStatusCompleted, req1.Status)
	req2, _ = s.app.DexKeeper.GetOrder(s.ctx, req2.PairId, req2.Id)
	s.Require().Equal(types.OrderStatusCompleted, req2.Status)

	s.Require().True(coinsEq(utils.ParseCoins("10000denom1"), s.getBalances(s.addr(1))))
	s.Require().True(coinsEq(utils.ParseCoins("10000denom2"), s.getBalances(s.addr(2))))

	pair, _ = s.app.DexKeeper.GetPair(s.ctx, pair.Id)
	s.Require().NotNil(pair.LastPrice)
	s.Require().True(decEq(utils.ParseDec("1.0"), *pair.LastPrice))
}

func (s *KeeperTestSuite) TestPartialMatch() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	order := s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("1.0"), math.NewInt(10000), time.Hour, true)
	s.sellLimitOrder(s.addr(2), pair.Id, utils.ParseDec("1.0"), math.NewInt(5000), 0, true)
	// FIX ME: Using s.nextBlock was not working properly so we manually run the end blocker calls for now
	dex.EndBlocker(s.ctx, s.app.DexKeeper)
	dex.BeginBlocker(s.ctx, s.app.DexKeeper)

	order, found := s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	s.Require().True(found)
	s.Require().Equal(types.OrderStatusPartiallyMatched, order.Status)
	s.Require().True(coinEq(utils.ParseCoin("5000denom2"), order.RemainingOfferCoin))
	s.Require().True(coinEq(utils.ParseCoin("5000denom1"), order.ReceivedCoin))
	s.Require().True(intEq(math.NewInt(5000), order.OpenAmount))

	s.sellMarketOrder(s.addr(3), pair.Id, math.NewInt(5000), 0, true)
	// FIX ME: Using s.nextBlock was not working properly so we manually run the end blocker calls for now
	dex.EndBlocker(s.ctx, s.app.DexKeeper)
	dex.BeginBlocker(s.ctx, s.app.DexKeeper)

	// Now completely matched.
	_, found = s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	s.Require().False(found)
}

func (s *KeeperTestSuite) TestMatchWithLowPricePool() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	// Create a pool with very low price.
	s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100000000000000000denom1,1000000denom2"), true)
	order := s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("0.000000000000010000"), math.NewInt(100000000000000000), 10*time.Second, true)
	s.nextBlock(false)
	order, found := s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	s.Require().True(found)
	s.Require().Equal(types.OrderStatusNotMatched, order.Status)
}

func (s *KeeperTestSuite) TestMMOrderNumberLimit() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pair.LastPrice = utils.ParseDecP("1.0")
	s.app.DexKeeper.SetPair(s.ctx, pair)

	for i := uint32(0); i < s.app.DexKeeper.GetMaxNumMarketMakingOrdersPerPair(s.ctx); i++ {
		price := math.LegacyNewDecWithPrec(int64(1000+i*5), 3)
		s.mmOrder(
			s.addr(1), pair.Id, types.OrderDirectionBuy,
			price, math.NewInt(1_000000), time.Hour, true)
	}

	s.nextBlock(false)
	// Cannot place an MM order anymore.
	s.fundAddr(s.addr(1), utils.ParseCoins("1000000000denom1,1000000000denom2"))
	_, err := s.app.DexKeeper.MMOrder(
		s.ctx,
		types.NewMsgMMOrder(
			s.addr(1), pair.Id, types.OrderDirectionBuy,
			utils.ParseCoin("1_100000denom2"), "denom1",
			utils.ParseDec("1.1"), math.NewInt(1_000000), 0))
	s.Require().ErrorIs(err, types.ErrMaxNumMMOrdersExceeded)

	// Complete one MM order.
	s.sellLimitOrder(s.addr(2), pair.Id, utils.ParseDec("0.9"), math.NewInt(1_000000), 0, true)
	s.nextBlock(false)

	// Now it is possible to place one more MM order.
	order := s.mmOrder(
		s.addr(1), pair.Id, types.OrderDirectionBuy,
		utils.ParseDec("1.1"), math.NewInt(1_000000), time.Hour, true)

	s.nextBlock(false)
	_, err = s.app.DexKeeper.MMOrder(
		s.ctx,
		types.NewMsgMMOrder(
			s.addr(1), pair.Id, types.OrderDirectionBuy,
			utils.ParseCoin("1_100000denom2"), "denom1",
			utils.ParseDec("1.1"), math.NewInt(1_000000), 0))
	s.Require().ErrorIs(err, types.ErrMaxNumMMOrdersExceeded)
	// Cancelling a market making order makes it possible to place more
	// MM orders again.
	s.cancelOrder(s.addr(1), order.PairId, order.Id)
	s.mmOrder(
		s.addr(1), pair.Id, types.OrderDirectionBuy,
		utils.ParseDec("1.1"), math.NewInt(1_000000), time.Hour, true)
}

func (s *KeeperTestSuite) TestCancelOrder() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	order := s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("1.0"), newInt(10000), types.DefaultMaxOrderLifespan, true)

	// Cannot cancel an order within a same batch
	err := s.app.DexKeeper.CancelOrder(s.ctx, types.NewMsgCancelOrder(s.addr(1), order.PairId, order.Id))
	s.Require().ErrorIs(err, types.ErrSameBatch)

	// FIX ME: Using s.nextBlock was not working properly so we manually run the end blocker calls for now
	dex.EndBlocker(s.ctx, s.app.DexKeeper)
	dex.BeginBlocker(s.ctx, s.app.DexKeeper)

	// Now an order can be canceled
	err = s.app.DexKeeper.CancelOrder(s.ctx, types.NewMsgCancelOrder(s.addr(1), order.PairId, order.Id))
	s.Require().NoError(err)

	order, found := s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	s.Require().True(found)
	s.Require().Equal(types.OrderStatusCanceled, order.Status)

	// Coins are refunded
	s.Require().True(coinsEq(utils.ParseCoins("10000denom2"), s.getBalances(s.addr(1))))

	// FIX ME: Using s.nextBlock was not working properly so we manually run the end blocker calls for now
	dex.EndBlocker(s.ctx, s.app.DexKeeper)
	dex.BeginBlocker(s.ctx, s.app.DexKeeper)

	// Order is deleted
	_, found = s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	s.Require().False(found)
}

func (s *KeeperTestSuite) TestCancelAllOrders() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	order := s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("1.0"), math.NewInt(10000), time.Hour, true)
	s.cancelAllOrders(s.addr(1), nil) // CancelAllOrders doesn't cancel orders within in same batch
	s.nextBlock(false)

	// The order is still alive.
	_, found := s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	s.Require().True(found)

	s.cancelAllOrders(s.addr(1), nil) // This time, it cancels the order.
	order, found = s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	// Canceling an order doesn't delete the order immediately.
	s.Require().True(found)
	// Instead, the order becomes canceled.
	s.Require().Equal(types.OrderStatusCanceled, order.Status)

	// The order won't be matched with this market order, since the order is
	// already canceled.
	s.sellLimitOrder(s.addr(3), pair.Id, utils.ParseDec("1.0"), math.NewInt(10000), 0, true)
	s.nextBlock(false)
	s.Require().True(coinsEq(utils.ParseCoins("10000denom2"), s.getBalances(s.addr(1))))

	pair2 := s.createPair(s.addr(0), "denom2", "denom3", true)
	s.buyLimitOrder(s.addr(2), pair.Id, utils.ParseDec("1.0"), math.NewInt(10000), time.Hour, true)
	s.sellLimitOrder(s.addr(2), pair.Id, utils.ParseDec("1.5"), math.NewInt(10000), time.Hour, true)
	s.sellLimitOrder(s.addr(2), pair2.Id, utils.ParseDec("1.0"), math.NewInt(10000), time.Hour, true)
	s.nextBlock(false)
	// CancelAllOrders can cancel orders in specific pairs.
	s.cancelAllOrders(s.addr(2), []uint64{pair.Id})
	// Coins from first two orders are refunded, but not from the last order.
	s.Require().True(coinsEq(utils.ParseCoins("10000denom2,10000denom1"), s.getBalances(s.addr(2))))
}

func (s *KeeperTestSuite) TestCancelAllOrdersGasUsage() {
	// Ensure that the number of other orders in pairs doesn't affect
	// the msg's gas usage.

	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	// 1000 other users make orders.
	for i := 1; i <= 1000; i++ {
		s.buyLimitOrder(s.addr(i), pair.Id, utils.ParseDec("0.9"), math.NewInt(10000), time.Minute, true)
		s.sellLimitOrder(s.addr(i), pair.Id, utils.ParseDec("1.1"), math.NewInt(10000), time.Minute, true)
	}

	// The orderer makes an order.
	orderer := s.addr(1001)
	s.sellLimitOrder(orderer, pair.Id, utils.ParseDec("1.1"), math.NewInt(10000), time.Minute, true)

	// New batch begins, now the orderer can cancel his/her order.
	s.nextBlock(false)

	s.ctx = s.ctx.WithGasMeter(store.NewInfiniteGasMeter()) // to record gas consumption
	s.cancelAllOrders(orderer, nil)                         // cancel all orders in all pairs
	s.Require().Less(s.ctx.GasMeter().GasConsumed(), store.Gas(50000))
}

func (s *KeeperTestSuite) TestDustCollector() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("0.9005"), newInt(1000), 0, true)
	s.sellLimitOrder(s.addr(2), pair.Id, utils.ParseDec("0.9005"), newInt(1000), 0, true)
	s.nextBlock(false)

	s.Require().True(coinsEq(utils.ParseCoins("1000denom1"), s.getBalances(s.addr(1))))
	s.Require().True(coinsEq(utils.ParseCoins("900denom2"), s.getBalances(s.addr(2))))

	s.Require().True(coinsEq(sdk.Coins{}, s.getBalances(pair.GetEscrowAddressAcc())))
	s.Require().True(coinsEq(utils.ParseCoins("1denom2"), s.getBalances(s.app.DexKeeper.GetDustCollector(s.ctx))))
}

func (s *KeeperTestSuite) TestFitPrice() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	lastPrice := utils.ParseDec("1")
	pair.LastPrice = &lastPrice
	s.app.DexKeeper.SetPair(s.ctx, pair)

	for _, tc := range []struct {
		name        string
		price       math.LegacyDec
		dir         types.OrderDirection
		expectedErr string
	}{
		{
			"",
			utils.ParseDec("1"),
			types.OrderDirectionBuy,
			"",
		},
		{
			"",
			utils.ParseDec("1"),
			types.OrderDirectionSell,
			"",
		},
		{
			"",
			utils.ParseDec("1.1"),
			types.OrderDirectionBuy,
			"",
		},
		{
			"",
			utils.ParseDec("0.9"),
			types.OrderDirectionSell,
			"",
		},
		{
			"",
			utils.ParseDec("1.099999999"),
			types.OrderDirectionBuy,
			"",
		},
		{
			"",
			utils.ParseDec("0.900000001"),
			types.OrderDirectionSell,
			"",
		},
		{
			"",
			utils.ParseDec("1.10000001"),
			types.OrderDirectionBuy,
			"1.100000010000000000 is higher than 1.100000000000000000: price out of range limit",
		},
		{
			"",
			utils.ParseDec("0.8999999"),
			types.OrderDirectionSell,
			"0.899999900000000000 is lower than 0.900000000000000000: price out of range limit",
		},
	} {
		s.Run(tc.name, func() {
			amt := newInt(10000)
			var offerCoin sdk.Coin
			var demandCoinDenom string
			switch tc.dir {
			case types.OrderDirectionBuy:
				offerCoin = sdk.NewCoin(pair.QuoteCoinDenom, tc.price.MulInt(amt).Ceil().TruncateInt())
				demandCoinDenom = pair.BaseCoinDenom
			case types.OrderDirectionSell:
				offerCoin = sdk.NewCoin(pair.BaseCoinDenom, amt)
				demandCoinDenom = pair.QuoteCoinDenom
			}
			s.fundAddr(s.addr(1), sdk.NewCoins(offerCoin))
			msg := types.NewMsgLimitOrder(s.addr(1), pair.Id, tc.dir, offerCoin, demandCoinDenom, tc.price, amt, 0)
			req, err := s.app.DexKeeper.LimitOrder(s.ctx, msg)
			if tc.expectedErr == "" {
				s.Require().NoError(err)
				switch tc.dir {
				case types.OrderDirectionBuy:
					s.Require().True(req.Price.LTE(tc.price))
				case types.OrderDirectionSell:
					s.Require().True(req.Price.GTE(tc.price))
				}
			} else {
				s.Require().EqualError(err, tc.expectedErr)
			}
		})
	}
}

func (s *KeeperTestSuite) TestGetOrdersByOrderer() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pair2 := s.createPair(s.addr(0), "denom2", "denom3", true)

	order1 := s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("1.0"), math.NewInt(10000), 0, true)
	order2 := s.sellLimitOrder(s.addr(1), pair2.Id, utils.ParseDec("1.0"), math.NewInt(10000), 0, true)

	orders := s.app.DexKeeper.GetOrdersByOrderer(s.ctx, s.addr(1))
	s.Require().Len(orders, 2)
	s.Require().Equal(order1.PairId, orders[0].PairId)
	s.Require().Equal(order1.Id, orders[0].Id)
	s.Require().Equal(order2.PairId, orders[1].PairId)
	s.Require().Equal(order2.Id, orders[1].Id)
}

func (s *KeeperTestSuite) TestInsufficientRemainingOfferCoin() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	order := s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("0.5"), math.NewInt(10000), time.Minute, true)
	s.sellLimitOrder(s.addr(2), pair.Id, utils.ParseDec("0.5"), math.NewInt(1001), 0, true)
	s.nextBlock(false)

	s.sellLimitOrder(s.addr(2), pair.Id, utils.ParseDec("0.5"), math.NewInt(8999), 0, true)
	s.nextBlock(false)
	order, found := s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	s.Require().True(found)
	s.Require().Equal(types.OrderStatusExpired, order.Status)
	s.Require().True(intEq(math.OneInt(), order.OpenAmount))
}

func (s *KeeperTestSuite) TestNegativeOpenAmount() {
	s.ctx = s.ctx.WithBlockHeight(1).WithBlockTime(utils.ParseTime("2022-03-01T00:00:00Z"))

	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	order := s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("0.82"), math.NewInt(648744), 0, true)
	s.sellLimitOrder(s.addr(2), pair.Id, utils.ParseDec("0.82"), math.NewInt(648745), 0, true)
	s.nextBlock(false)

	order, found := s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	s.Require().True(found)
	s.Require().False(order.OpenAmount.IsNegative())

	genState := s.app.DexKeeper.ExportGenesis(s.ctx)
	s.Require().NotPanics(func() {
		s.app.DexKeeper.InitGenesis(s.ctx, *genState)
	})
}

func (s *KeeperTestSuite) TestRejectSmallOrders() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	s.fundAddr(s.addr(1), utils.ParseCoins("10000000denom1,10000000denom2"))

	// Too small offer coin amount.
	msg := types.NewMsgLimitOrder(
		s.addr(1), pair.Id, types.OrderDirectionBuy, utils.ParseCoin("99denom2"),
		"denom1", utils.ParseDec("0.1"), math.NewInt(990), 0)
	s.Require().EqualError(msg.ValidateBasic(), "offer coin 99denom2 is smaller than the min amount 100: invalid request")

	// Too small order amount.
	msg = types.NewMsgLimitOrder(
		s.addr(1), pair.Id, types.OrderDirectionBuy, utils.ParseCoin("990denom2"),
		"denom1", utils.ParseDec("10.0"), math.NewInt(99), 0)
	s.Require().EqualError(msg.ValidateBasic(), "order amount 99 is smaller than the min amount 100: invalid request")

	// Too small orders.
	msg = types.NewMsgLimitOrder(
		s.addr(1), pair.Id, types.OrderDirectionBuy, utils.ParseCoin("101denom2"),
		"denom1", utils.ParseDec("0.00010001"), math.NewInt(999000), 0)
	s.Require().NoError(msg.ValidateBasic())
	_, err := s.app.DexKeeper.LimitOrder(s.ctx, msg)
	s.Require().ErrorIs(err, types.ErrTooSmallOrder)

	msg = types.NewMsgLimitOrder(
		s.addr(1), pair.Id, types.OrderDirectionSell, utils.ParseCoin("999999denom1"),
		"denom2", utils.ParseDec("0.0001"), math.NewInt(999999), 0)
	s.Require().NoError(msg.ValidateBasic())
	_, err = s.app.DexKeeper.LimitOrder(s.ctx, msg)
	s.Require().ErrorIs(err, types.ErrTooSmallOrder)

	// Too small offer coin amount.
	msg2 := types.NewMsgMarketOrder(
		s.addr(1), pair.Id, types.OrderDirectionSell, utils.ParseCoin("99denom1"),
		"denom2", math.NewInt(99), 0)
	s.Require().EqualError(msg2.ValidateBasic(), "offer coin 99denom1 is smaller than the min amount 100: invalid request")

	// Too small order amount.
	msg2 = types.NewMsgMarketOrder(
		s.addr(1), pair.Id, types.OrderDirectionSell, utils.ParseCoin("100denom1"),
		"denom2", math.NewInt(99), 0)
	s.Require().EqualError(msg2.ValidateBasic(), "order amount 99 is smaller than the min amount 100: invalid request")

	p := utils.ParseDec("0.0001")
	pair.LastPrice = &p
	s.app.DexKeeper.SetPair(s.ctx, pair)

	// Too small orders.
	msg2 = types.NewMsgMarketOrder(
		s.addr(1), pair.Id, types.OrderDirectionBuy, utils.ParseCoin("100denom2"),
		"denom1", math.NewInt(909090), 0)
	s.Require().NoError(msg2.ValidateBasic())
	_, err = s.app.DexKeeper.MarketOrder(s.ctx, msg2)
	s.Require().ErrorIs(err, types.ErrTooSmallOrder)

	msg2 = types.NewMsgMarketOrder(
		s.addr(1), pair.Id, types.OrderDirectionSell, utils.ParseCoin("1000denom1"),
		"denom2", math.NewInt(1000), 0)
	s.Require().NoError(msg2.ValidateBasic())
	_, err = s.app.DexKeeper.MarketOrder(s.ctx, msg2)
	s.Require().ErrorIs(err, types.ErrTooSmallOrder)
}

func (s *KeeperTestSuite) TestExpireSmallOrders() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("0.000018"), math.NewInt(10000000), time.Minute, true)
	// This order should have 10000 open amount after matching.
	// If this order would be matched after that, then the orderer will receive
	// floor(10000*0.000018) demand coin, which is zero.
	// So the order must have been expired after matching.
	order := s.sellLimitOrder(s.addr(2), pair.Id, utils.ParseDec("0.000018"), math.NewInt(10010000), time.Minute, true)
	s.nextBlock(false)
	order, found := s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	s.Require().True(found)
	s.Require().Equal(types.OrderStatusExpired, order.Status)

	s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("0.000019"), math.NewInt(100000000), time.Minute, true)
	s.sellLimitOrder(s.addr(3), pair.Id, utils.ParseDec("0.000019"), math.NewInt(100000000), time.Minute, true)
	s.nextBlock(false)
}

func (s *KeeperTestSuite) TestPoolOrderOverflow() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	i, _ := math.NewIntFromString("10000000000000000000000000")
	s.createPool(s.addr(0), pair.Id, sdk.NewCoins(sdk.NewInt64Coin("denom1", 1e6), sdk.NewCoin("denom2", i)), true)

	s.sellLimitOrder(s.addr(1), pair.Id, utils.ParseDec("0.000000000000010000"), math.NewInt(1e17), 0, true)
	s.Require().NotPanics(func() {
		s.nextBlock(false)
	})
}

func (s *KeeperTestSuite) TestRangeddex() {
	orderPrice := utils.ParseDec("1.05")
	orderAmt := math.NewInt(100000)

	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pair.LastPrice = utils.ParseDecP("1.0")
	s.app.DexKeeper.SetPair(s.ctx, pair)

	s.createPool(s.addr(1), pair.Id, utils.ParseCoins("1000000denom1,1000000denom2"), true)

	order := s.buyLimitOrder(s.addr(2), pair.Id, orderPrice, orderAmt, 0, true)
	s.nextBlock(false)
	order, _ = s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	paid := order.OfferCoin.Sub(order.RemainingOfferCoin).Amount
	received := order.ReceivedCoin.Amount
	s.Require().True(received.LT(orderAmt))
	s.Require().True(paid.ToLegacyDec().QuoInt(received).LTE(orderPrice))

	pair = s.createPair(s.addr(0), "denom3", "denom4", true)
	pair.LastPrice = utils.ParseDecP("1.0")
	s.app.DexKeeper.SetPair(s.ctx, pair)

	s.createRangedPool(
		s.addr(1), pair.Id, utils.ParseCoins("1000000denom3,1000000denom4"),
		utils.ParseDec("0.8"), utils.ParseDec("1.3"), utils.ParseDec("1.0"), true)
	order = s.buyLimitOrder(s.addr(2), pair.Id, orderPrice, orderAmt, 0, true)
	s.nextBlock(false)
	order, _ = s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	paid = order.OfferCoin.Sub(order.RemainingOfferCoin).Amount
	received = order.ReceivedCoin.Amount
	s.Require().True(intEq(orderAmt, received))
	s.Require().True(paid.ToLegacyDec().QuoInt(received).LTE(orderPrice))
}

func (s *KeeperTestSuite) TestOneSidedRangedPool() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pair.LastPrice = utils.ParseDecP("1.0")
	s.app.DexKeeper.SetPair(s.ctx, pair)

	pool := s.createRangedPool(
		s.addr(1), pair.Id, utils.ParseCoins("1000000denom1,1000000denom2"),
		utils.ParseDec("1.0"), utils.ParseDec("1.2"), utils.ParseDec("1.0"), true)
	rx, ry := s.app.DexKeeper.GetPoolBalances(s.ctx, pool)
	ammPool := pool.AMMPool(rx.Amount, ry.Amount, math.Int{})
	s.Require().True(utils.DecApproxEqual(utils.ParseDec("1.0"), ammPool.Price()))
	s.Require().True(intEq(math.ZeroInt(), rx.Amount))
	s.Require().True(intEq(math.NewInt(1000000), ry.Amount))

	orderPrice := utils.ParseDec("1.1")
	orderAmt := math.NewInt(100000)
	order := s.buyLimitOrder(s.addr(2), pair.Id, utils.ParseDec("1.1"), math.NewInt(100000), 0, true)
	s.nextBlock(false)
	order, _ = s.app.DexKeeper.GetOrder(s.ctx, order.PairId, order.Id)
	paid := order.OfferCoin.Sub(order.RemainingOfferCoin).Amount
	received := order.ReceivedCoin.Amount
	s.Require().True(intEq(orderAmt, received))
	s.Require().True(paid.ToLegacyDec().QuoInt(received).LTE(orderPrice))

	rx, _ = s.app.DexKeeper.GetPoolBalances(s.ctx, pool)
	s.Require().True(rx.IsPositive())
}

func (s *KeeperTestSuite) TestExhaustRangedPool() {
	r := rand.New(rand.NewSource(0))

	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	minPrice, maxPrice := utils.ParseDec("0.5"), utils.ParseDec("2.0")
	initialPrice := utils.ParseDec("1.0")
	pool := s.createRangedPool(
		s.addr(1), pair.Id, utils.ParseCoins("1000000denom1,1000000denom2"),
		minPrice, maxPrice, initialPrice, true)

	orderer := s.addr(2)
	s.fundAddr(orderer, utils.ParseCoins("10000000denom1,10000000denom2"))

	// Buy
	for {
		rx, ry := s.app.DexKeeper.GetPoolBalances(s.ctx, pool)
		ammPool := pool.AMMPool(rx.Amount, ry.Amount, math.Int{})
		poolPrice := ammPool.Price()
		if ry.Amount.LT(math.NewInt(100)) {
			s.Require().True(utils.DecApproxEqual(maxPrice, poolPrice))
			break
		}
		orderPrice := utils.RandomDec(r, poolPrice, poolPrice.Mul(math.LegacyNewDecWithPrec(105, 2)))
		amt := utils.RandomInt(r, math.NewInt(5000), math.NewInt(15000))
		s.buyLimitOrder(orderer, pair.Id, orderPrice, amt, 0, false)
		s.nextBlock(false)
	}

	// Sell
	for {
		rx, ry := s.app.DexKeeper.GetPoolBalances(s.ctx, pool)
		ammPool := pool.AMMPool(rx.Amount, ry.Amount, math.Int{})
		poolPrice := ammPool.Price()
		if rx.Amount.LT(math.NewInt(100)) {
			s.Require().True(utils.DecApproxEqual(minPrice, poolPrice))
			break
		}
		orderPrice := utils.RandomDec(r, poolPrice.Mul(math.LegacyNewDecWithPrec(95, 2)), poolPrice)
		amt := utils.RandomInt(r, math.NewInt(5000), math.NewInt(15000))
		s.sellLimitOrder(orderer, pair.Id, orderPrice, amt, 0, false)
		s.nextBlock(false)
	}

	// Buy again
	for {
		rx, ry := s.app.DexKeeper.GetPoolBalances(s.ctx, pool)
		ammPool := pool.AMMPool(rx.Amount, ry.Amount, math.Int{})
		poolPrice := ammPool.Price()
		if poolPrice.GTE(initialPrice) {
			break
		}
		orderPrice := utils.RandomDec(r, poolPrice, poolPrice.Mul(math.LegacyNewDecWithPrec(105, 2)))
		amt := utils.RandomInt(r, math.NewInt(5000), math.NewInt(15000))
		s.buyLimitOrder(orderer, pair.Id, orderPrice, amt, 0, false)
		s.nextBlock(false)
	}

	rx, ry := s.app.DexKeeper.GetPoolBalances(s.ctx, pool)
	ammPool := pool.AMMPool(rx.Amount, ry.Amount, math.Int{})
	fmt.Println(rx, ry, ammPool.Price())

	fmt.Println(s.getBalances(s.app.DexKeeper.GetDustCollector(s.ctx)))
	fmt.Println(s.getBalances(orderer))
}

func (s *KeeperTestSuite) TestSwap_edgecase1() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	s.sellLimitOrder(s.addr(2), pair.Id, utils.ParseDec("0.102"), math.NewInt(10000), 0, true)
	s.sellLimitOrder(s.addr(3), pair.Id, utils.ParseDec("0.101"), math.NewInt(9995), 0, true)
	s.buyLimitOrder(s.addr(4), pair.Id, utils.ParseDec("0.102"), math.NewInt(10000), 0, true)
	s.nextBlock(false)
	pair, _ = s.app.DexKeeper.GetPair(s.ctx, pair.Id)
	s.Require().True(decEq(utils.ParseDec("0.102"), *pair.LastPrice))

	s.sellLimitOrder(s.addr(2), pair.Id, utils.ParseDec("0.102"), math.NewInt(10000), 0, true)
	s.sellLimitOrder(s.addr(3), pair.Id, utils.ParseDec("0.101"), math.NewInt(9995), 0, true)
	s.buyLimitOrder(s.addr(4), pair.Id, utils.ParseDec("0.102"), math.NewInt(10000), 0, true)
	s.nextBlock(false)
}

func (s *KeeperTestSuite) TestSwap_edgecase2() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pair.LastPrice = utils.ParseDecP("1.6724")
	s.app.DexKeeper.SetPair(s.ctx, pair)

	s.createPool(s.addr(0), pair.Id, utils.ParseCoins("1005184935980denom2,601040339855denom1"), true)
	s.createRangedPool(
		s.addr(0), pair.Id, utils.ParseCoins("17335058855denom2"),
		utils.ParseDec("1.15"), utils.ParseDec("1.55"), utils.ParseDec("1.55"), true)
	s.createRangedPool(
		s.addr(0), pair.Id, utils.ParseCoins("217771046279denom2"),
		utils.ParseDec("1.25"), utils.ParseDec("1.45"), utils.ParseDec("1.45"), true)

	s.sellMarketOrder(s.addr(1), pair.Id, math.NewInt(4336_000000), 0, true)
	s.nextBlock(false)

	pair, _ = s.app.DexKeeper.GetPair(s.ctx, pair.Id)
	s.Require().True(decEq(utils.ParseDec("1.6484"), *pair.LastPrice))

	s.nextBlock(false)
	pair, _ = s.app.DexKeeper.GetPair(s.ctx, pair.Id)
	s.Require().True(decEq(utils.ParseDec("1.6484"), *pair.LastPrice))

	s.sellMarketOrder(s.addr(1), pair.Id, math.NewInt(4450_000000), 0, true)
	s.nextBlock(false)

	pair, _ = s.app.DexKeeper.GetPair(s.ctx, pair.Id)
	s.Require().True(decEq(utils.ParseDec("1.6248"), *pair.LastPrice))
}

func (s *KeeperTestSuite) TestSwap_edgecase3() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pair.LastPrice = utils.ParseDecP("0.99992")
	s.app.DexKeeper.SetPair(s.ctx, pair)

	s.createPool(s.addr(0), pair.Id, utils.ParseCoins("110001546090denom2,110013588106denom1"), true)
	s.createRangedPool(
		s.addr(0), pair.Id, utils.ParseCoins("140913832254denom2,130634675302denom1"),
		utils.ParseDec("0.92"), utils.ParseDec("1.08"), utils.ParseDec("0.99989"), true)

	s.buyMarketOrder(s.addr(1), pair.Id, math.NewInt(30_000000), 0, true)
	s.nextBlock(false)

	pair, _ = s.app.DexKeeper.GetPair(s.ctx, pair.Id)
	s.Require().True(decEq(utils.ParseDec("0.99992"), *pair.LastPrice))
}

func (s *KeeperTestSuite) TestSwap_edgecase4() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pair.LastPrice = utils.ParseDecP("0.99999")
	s.app.DexKeeper.SetPair(s.ctx, pair)

	s.createPool(s.addr(0), pair.Id, utils.ParseCoins("1000_000000denom1,100_000000denom2"), true)

	s.createRangedPool(s.addr(0), pair.Id, utils.ParseCoins("1000_000000denom1,1000_000000denom2"),
		utils.ParseDec("0.95"), utils.ParseDec("1.05"), utils.ParseDec("1.02"), true)
	s.createRangedPool(s.addr(0), pair.Id, utils.ParseCoins("1000_000000denom1,1000_000000denom2"),
		utils.ParseDec("0.9"), utils.ParseDec("1.2"), utils.ParseDec("0.98"), true)

	s.sellLimitOrder(s.addr(1), pair.Id, utils.ParseDec("1.05"), math.NewInt(50_000000), 0, true)
	s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("0.97"), math.NewInt(100_000000), 0, true)

	s.nextBlock(false)
}

func (s *KeeperTestSuite) TestOrderBooks_edgecase1() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pair.LastPrice = utils.ParseDecP("0.57472")
	s.app.DexKeeper.SetPair(s.ctx, pair)

	s.createPool(s.addr(0), pair.Id, utils.ParseCoins("991883358661denom2,620800303846denom1"), true)
	s.createRangedPool(
		s.addr(0), pair.Id, utils.ParseCoins("155025981873denom2,4703143223denom1"),
		utils.ParseDec("1.15"), utils.ParseDec("1.55"), utils.ParseDec("1.5308"), true)
	s.createRangedPool(
		s.addr(0), pair.Id, utils.ParseCoins("223122824634denom2,26528571912denom1"),
		utils.ParseDec("1.25"), utils.ParseDec("1.45"), utils.ParseDec("1.4199"), true)

	resp, err := s.app.DexKeeper.OrderBooks(sdk.WrapSDKContext(s.ctx), &types.QueryOrderBooksRequest{
		PairIds:  []uint64{pair.Id},
		NumTicks: 10,
	})
	s.Require().NoError(err)
	s.Require().Len(resp.Pairs, 1)
	s.Require().Len(resp.Pairs[0].OrderBooks, 3)

	s.Require().Len(resp.Pairs[0].OrderBooks[0].Buys, 2)
	s.Require().True(decEq(utils.ParseDec("0.63219"), resp.Pairs[0].OrderBooks[0].Buys[0].Price))
	s.Require().True(intEq(math.NewInt(1178846737645), resp.Pairs[0].OrderBooks[0].Buys[0].UserOrderAmount))
	s.Require().True(decEq(utils.ParseDec("0.5187"), resp.Pairs[0].OrderBooks[0].Buys[1].Price))
	s.Require().True(intEq(math.NewInt(13340086), resp.Pairs[0].OrderBooks[0].Buys[1].UserOrderAmount))
	s.Require().Len(resp.Pairs[0].OrderBooks[0].Sells, 0)
}

func (s *KeeperTestSuite) TestPoolPreserveK() {
	r := rand.New(rand.NewSource(0))

	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	tickPrec := s.app.DexKeeper.GetTickPrecision(s.ctx)
	for i := 0; i < 10; i++ {
		minPrice := amm.RandomTick(r, utils.ParseDec("0.001"), utils.ParseDec("10.0"), int(tickPrec))
		maxPrice := amm.RandomTick(r, minPrice.Mul(utils.ParseDec("1.01")), utils.ParseDec("100.0"), int(tickPrec))
		initialPrice := amm.RandomTick(r, minPrice, maxPrice, int(tickPrec))
		s.createRangedPool(
			s.addr(1), pair.Id, utils.ParseCoins("1_000000000000denom1,1_000000000000denom2"),
			minPrice, maxPrice, initialPrice, true)
	}

	pools := s.app.DexKeeper.GetAllPools(s.ctx)

	ks := map[uint64]math.LegacyDec{}
	for _, pool := range pools {
		rx, ry := s.app.DexKeeper.GetPoolBalances(s.ctx, pool)
		ammPool := pool.AMMPool(rx.Amount, ry.Amount, math.Int{}).(*amm.RangedPool)
		transX, transY := ammPool.Translation()
		ks[pool.Id] = rx.Amount.ToLegacyDec().Add(transX).Mul(ry.Amount.ToLegacyDec().Add(transY))
	}

	for i := 0; i < 20; i++ {
		pair, _ = s.app.DexKeeper.GetPair(s.ctx, pair.Id)
		for j := 0; j < 50; j++ {
			var price math.LegacyDec
			if pair.LastPrice == nil {
				price = utils.RandomDec(r, utils.ParseDec("0.001"), utils.ParseDec("100.0"))
			} else {
				price = utils.RandomDec(r, utils.ParseDec("0.91"), utils.ParseDec("1.09")).Mul(*pair.LastPrice)
			}
			amt := utils.RandomInt(r, math.NewInt(10000), math.NewInt(1000000))
			lifespan := time.Duration(r.Intn(60)) * time.Second
			if r.Intn(2) == 0 {
				s.buyLimitOrder(s.addr(j+2), pair.Id, price, amt, lifespan, true)
			} else {
				s.buyLimitOrder(s.addr(j+2), pair.Id, price, amt, lifespan, true)
			}
		}

		dex.EndBlocker(s.ctx, s.app.DexKeeper)
		s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(3 * time.Second))
		dex.BeginBlocker(s.ctx, s.app.DexKeeper)

		for _, pool := range pools {
			rx, ry := s.app.DexKeeper.GetPoolBalances(s.ctx, pool)
			ammPool := pool.AMMPool(rx.Amount, ry.Amount, math.Int{}).(*amm.RangedPool)
			transX, transY := ammPool.Translation()
			k := rx.Amount.ToLegacyDec().Add(transX).Mul(ry.Amount.ToLegacyDec().Add(transY))
			s.Require().True(k.GTE(ks[pool.Id].Mul(utils.ParseDec("0.99999")))) // there may be a small error
			ks[pool.Id] = k
		}
	}
}
