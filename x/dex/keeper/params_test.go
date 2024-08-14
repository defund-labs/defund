package keeper_test

import (
	"defund/x/dex/types"
)

func (s *KeeperTestSuite) TestGetBatchSize() {
	s.Require().EqualValues(types.DefaultBatchSize, s.app.DexKeeper.GetBatchSize(s.ctx))
}

func (s *KeeperTestSuite) TestGetTickPrecision() {
	s.Require().EqualValues(types.DefaultTickPrecision, s.app.DexKeeper.GetTickPrecision(s.ctx))
}

func (s *KeeperTestSuite) TestGetFeeCollector() {
	s.Require().EqualValues(types.DefaultFeeCollectorAddress, s.app.DexKeeper.GetFeeCollector(s.ctx))
}

func (s *KeeperTestSuite) TestGetDustCollector() {
	s.Require().EqualValues(types.DefaultDustCollectorAddress, s.app.DexKeeper.GetDustCollector(s.ctx))
}

func (s *KeeperTestSuite) TestGetMinInitialPoolCoinSupply() {
	s.Require().EqualValues(types.DefaultMinInitialPoolCoinSupply, s.app.DexKeeper.GetMinInitialPoolCoinSupply(s.ctx))
}

func (s *KeeperTestSuite) TestGetPairCreationFee() {
	s.Require().EqualValues(types.DefaultPairCreationFee, s.app.DexKeeper.GetPairCreationFee(s.ctx))
}

func (s *KeeperTestSuite) TestGetPoolCreationFee() {
	s.Require().EqualValues(types.DefaultPoolCreationFee, s.app.DexKeeper.GetPoolCreationFee(s.ctx))
}

func (s *KeeperTestSuite) TestGetMinInitialDepositAmount() {
	s.Require().EqualValues(types.DefaultMinInitialDepositAmount, s.app.DexKeeper.GetMinInitialDepositAmount(s.ctx))
}

func (s *KeeperTestSuite) TestGetMaxPriceLimitRatio() {
	s.Require().EqualValues(types.DefaultMaxPriceLimitRatio, s.app.DexKeeper.GetMaxPriceLimitRatio(s.ctx))
}

func (s *KeeperTestSuite) TestGetMaxNumMarketMakingOrderTicks() {
	s.Require().EqualValues(types.DefaultMaxNumMarketMakingOrderTicks, s.app.DexKeeper.GetMaxNumMarketMakingOrderTicks(s.ctx))
}

func (s *KeeperTestSuite) TestGetMaxNumMarketMakingOrdersPerPair() {
	s.Require().EqualValues(types.DefaultMaxNumMarketMakingOrdersPerPair, s.app.DexKeeper.GetMaxNumMarketMakingOrdersPerPair(s.ctx))
}

func (s *KeeperTestSuite) TestGetMaxOrderLifespan() {
	s.Require().EqualValues(types.DefaultMaxOrderLifespan, s.app.DexKeeper.GetMaxOrderLifespan(s.ctx))
}

func (s *KeeperTestSuite) TestGetWithdrawFeeRate() {
	s.Require().EqualValues(types.DefaultWithdrawFeeRate, s.app.DexKeeper.GetWithdrawFeeRate(s.ctx))
}

func (s *KeeperTestSuite) TestGetDepositExtraGas() {
	s.Require().EqualValues(types.DefaultDepositExtraGas, s.app.DexKeeper.GetDepositExtraGas(s.ctx))
}

func (s *KeeperTestSuite) TestGetWithdrawExtraGas() {
	s.Require().EqualValues(types.DefaultWithdrawExtraGas, s.app.DexKeeper.GetWithdrawExtraGas(s.ctx))
}

func (s *KeeperTestSuite) TestGetOrderExtraGas() {
	s.Require().EqualValues(types.DefaultOrderExtraGas, s.app.DexKeeper.GetOrderExtraGas(s.ctx))
}

func (s *KeeperTestSuite) TestGetMaxNumActivePoolsPerPair() {
	s.Require().EqualValues(types.DefaultMaxNumActivePoolsPerPair, s.app.DexKeeper.GetMaxNumActivePoolsPerPair(s.ctx))
}
