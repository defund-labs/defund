package keeper_test

import (
	"time"

	"cosmossdk.io/math"

	utils "defund/types"
	"defund/x/dex/keeper"
	"defund/x/dex/types"
)

func (s *KeeperTestSuite) TestDepositCoinsEscrowInvariant() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("1000000denom1,1000000denom2"), true)

	req := s.deposit(s.addr(1), pool.Id, utils.ParseCoins("1000000denom1,1000000denom2"), true)
	_, broken := keeper.DepositCoinsEscrowInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().False(broken)

	oldReq := req
	req.DepositCoins = utils.ParseCoins("2000000denom1,2000000denom2")
	s.app.DexKeeper.SetDepositRequest(s.ctx, req)
	_, broken = keeper.DepositCoinsEscrowInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().True(broken)

	req = oldReq
	s.app.DexKeeper.SetDepositRequest(s.ctx, req)
	s.nextBlock(false)
	_, broken = keeper.DepositCoinsEscrowInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().False(broken)
}

func (s *KeeperTestSuite) TestPoolCoinEscrowInvariant() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("1000000denom1,1000000denom2"), true)

	s.deposit(s.addr(1), pool.Id, utils.ParseCoins("1000000denom1,1000000denom2"), true)
	s.nextBlock(false)

	req := s.withdraw(s.addr(1), pool.Id, utils.ParseCoin("1000000pool1"))
	_, broken := keeper.PoolCoinEscrowInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().False(broken)

	oldReq := req
	req.PoolCoin = utils.ParseCoin("2000000pool1")
	s.app.DexKeeper.SetWithdrawRequest(s.ctx, req)
	_, broken = keeper.PoolCoinEscrowInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().True(broken)

	req = oldReq
	s.app.DexKeeper.SetWithdrawRequest(s.ctx, req)
	s.nextBlock(false)
	_, broken = keeper.PoolCoinEscrowInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().False(broken)
}

func (s *KeeperTestSuite) TestRemainingOfferCoinEscrowInvariant() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	order := s.buyLimitOrder(s.addr(1), pair.Id, utils.ParseDec("1.0"), newInt(1000000), 0, true)
	_, broken := keeper.RemainingOfferCoinEscrowInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().False(broken)

	oldOrder := order
	order.RemainingOfferCoin = utils.ParseCoin("2000000denom1")
	s.app.DexKeeper.SetOrder(s.ctx, order)
	_, broken = keeper.RemainingOfferCoinEscrowInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().True(broken)

	order = oldOrder
	s.app.DexKeeper.SetOrder(s.ctx, order)
	s.nextBlock(false)
	_, broken = keeper.RemainingOfferCoinEscrowInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().False(broken)
}

func (s *KeeperTestSuite) TestPoolStatusInvariant() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("1000000denom1,1000000denom2"), true)

	_, broken := keeper.PoolStatusInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().False(broken)

	s.withdraw(s.addr(0), pool.Id, s.getBalance(s.addr(0), pool.PoolCoinDenom))
	s.nextBlock(false)

	_, broken = keeper.PoolStatusInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().False(broken)

	pool, _ = s.app.DexKeeper.GetPool(s.ctx, pool.Id)
	pool.Disabled = false
	s.app.DexKeeper.SetPool(s.ctx, pool)
	_, broken = keeper.PoolStatusInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().True(broken)
}

func (s *KeeperTestSuite) TestNumMMOrdersInvariant() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	orderer := s.addr(1)
	// Place random MM orders
	s.mmOrder(
		orderer, pair.Id, types.OrderDirectionBuy,
		utils.ParseDec("0.99"), math.NewInt(1000000), time.Hour, true)
	s.mmOrder(
		orderer, pair.Id, types.OrderDirectionBuy,
		utils.ParseDec("0.98"), math.NewInt(1000000), time.Hour, true)
	s.mmOrder(
		orderer, pair.Id, types.OrderDirectionBuy,
		utils.ParseDec("0.97"), math.NewInt(1000000), time.Hour, true)
	s.mmOrder(
		orderer, pair.Id, types.OrderDirectionSell,
		utils.ParseDec("1.01"), math.NewInt(1000000), time.Hour, true)
	s.mmOrder(
		orderer, pair.Id, types.OrderDirectionSell,
		utils.ParseDec("1.02"), math.NewInt(1000000), time.Hour, true)
	s.mmOrder(
		orderer, pair.Id, types.OrderDirectionSell,
		utils.ParseDec("1.03"), math.NewInt(1000000), time.Hour, true)

	_, broken := keeper.NumMMOrdersInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().False(broken)

	s.nextBlock(false)

	// Cancel some MM orders and place another order
	s.cancelOrder(orderer, pair.Id, 1)
	s.cancelOrder(orderer, pair.Id, 2)
	s.mmOrder(
		orderer, pair.Id, types.OrderDirectionSell,
		utils.ParseDec("1.04"), math.NewInt(1000000), time.Hour, true)

	_, broken = keeper.NumMMOrdersInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().False(broken)

	// After deleting canceled orders, the invariant must not be broken
	s.nextBlock(false)
	_, broken = keeper.NumMMOrdersInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().False(broken)

	// Break it
	s.app.DexKeeper.SetNumMMOrders(s.ctx, orderer, pair.Id, 3)
	_, broken = keeper.NumMMOrdersInvariant(s.app.DexKeeper)(s.ctx)
	s.Require().True(broken)
}
