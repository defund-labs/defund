package testutil

import (
	"fmt"
	"testing"
	"time"

	"cosmossdk.io/math"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	chain "defund/app"
	utils "defund/types"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

type TestSuite struct {
	suite.Suite

	App *chain.App
	Ctx sdk.Context
}

func (s *TestSuite) SetupTest() {
	app, ctx := TestApp(s.T())
	s.App = app
	s.Ctx = ctx
	hdr := cmtproto.Header{
		Height: 1,
		Time:   utils.ParseTime("2022-01-01T00:00:00Z"),
	}
	s.App.BeginBlocker(s.Ctx.WithBlockHeader(hdr))
}

func (s *TestSuite) BeginBlock(blockTimeDelta time.Duration) {
	hdr := cmtproto.Header{
		Height: s.App.LastBlockHeight() + 1,
		Time:   s.Ctx.BlockTime().Add(5 * time.Second),
	}
	s.App.BeginBlocker(s.Ctx.WithBlockHeader(hdr))
}

func (s *TestSuite) EndBlock() {
	s.T().Helper()
	s.App.EndBlocker(s.Ctx)
	s.App.Commit()
}

func (s *TestSuite) NextBlock() {
	s.T().Helper()
	s.EndBlock()
	s.BeginBlock(5 * time.Second)
}

func (s *TestSuite) GetBalance(addr sdk.AccAddress, denom string) sdk.Coin {
	s.T().Helper()
	return s.App.BankKeeper.GetBalance(s.Ctx, addr, denom)
}

func (s *TestSuite) GetAllBalances(addr sdk.AccAddress) sdk.Coins {
	s.T().Helper()
	return s.App.BankKeeper.GetAllBalances(s.Ctx, addr)
}

func (s *TestSuite) FundAccount(
	addr sdk.AccAddress, amt sdk.Coins) {
	s.T().Helper()
	if amt.IsAllPositive() {
		s.Require().NoError(s.App.BankKeeper.MintCoins(s.Ctx, minttypes.ModuleName, amt))
		s.Require().NoError(s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, minttypes.ModuleName, addr, amt))
	}
}

func (s *TestSuite) FundedAccount(addrNum int, amt sdk.Coins) sdk.AccAddress {
	s.T().Helper()
	addr := utils.TestAddress(addrNum)
	s.FundAccount(addr, amt)
	return addr
}

func (s *TestSuite) CheckEvent(evtType proto.Message, attrs map[string][]byte) {
	s.T().Helper()
	evtTypeName := proto.MessageName(evtType)
	for _, ev := range s.Ctx.EventManager().ABCIEvents() {
		if ev.Type == evtTypeName {
			attrMap := make(map[string][]byte)
			for _, attr := range ev.Attributes {
				attrMap[string(attr.Key)] = []byte(attr.GetValue())
			}
			for k, v := range attrs {
				value, ok := attrMap[k]
				s.Require().Truef(ok, "key %s not found", k)
				s.Require().Equal(v, value)
			}
			return
		}
	}
	s.FailNowf("CheckEvent failed", "event with type %s not found", evtTypeName)
}

func (s *TestSuite) AssertEqual(exp, got any) {
	s.T().Helper()
	var equal bool
	switch exp := exp.(type) {
	case math.Int:
		equal = exp.Equal(got.(math.Int))
	case math.LegacyDec:
		equal = exp.Equal(got.(math.LegacyDec))
	case sdk.Coin:
		equal = exp.IsEqual(got.(sdk.Coin))
	case sdk.Coins:
		equal = exp.Equal(got.(sdk.Coins))
	case sdk.DecCoin:
		equal = exp.IsEqual(got.(sdk.DecCoin))
	case sdk.DecCoins:
		equal = exp.Equal(got.(sdk.DecCoins))
	default:
		panic(fmt.Sprintf("unsupported type: %T", exp))
	}
	s.Assert().True(equal, "expected:\t%v\ngot:\t\t%v", exp, got)
}

func AssertEqual(t *testing.T, exp, got any) {
	t.Helper()
	var equal bool
	switch exp := exp.(type) {
	case math.Int:
		equal = exp.Equal(got.(math.Int))
	case math.LegacyDec:
		equal = exp.Equal(got.(math.LegacyDec))
	case sdk.Coin:
		equal = exp.IsEqual(got.(sdk.Coin))
	case sdk.Coins:
		equal = exp.Equal(got.(sdk.Coins))
	case sdk.DecCoin:
		equal = exp.IsEqual(got.(sdk.DecCoin))
	case sdk.DecCoins:
		equal = exp.Equal(got.(sdk.DecCoins))
	default:
		panic(fmt.Sprintf("unsupported type: %T", exp))
	}
	assert.True(t, equal, "expected:\t%v\ngot:\t\t%v", exp, got)
}
