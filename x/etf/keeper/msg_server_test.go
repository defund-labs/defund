package keeper_test

import (
	"fmt"
	"os"

	"github.com/CosmWasm/wasmd/x/wasm/ioutils"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	"github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
	"github.com/stretchr/testify/require"
)

func parseStoreCodeArgs(file string, sender sdk.AccAddress) (wasmtypes.MsgStoreCode, error) {
	wasm, err := os.ReadFile(file)
	if err != nil {
		return wasmtypes.MsgStoreCode{}, err
	}

	// gzip the wasm file
	if ioutils.IsWasm(wasm) {
		wasm, err = ioutils.GzipIt(wasm)

		if err != nil {
			return wasmtypes.MsgStoreCode{}, err
		}
	} else if !ioutils.IsGzip(wasm) {
		return wasmtypes.MsgStoreCode{}, fmt.Errorf("invalid input file. Use wasm binary or gzip")
	}

	perm := &wasmtypes.AllowEverybody
	if err != nil {
		return wasmtypes.MsgStoreCode{}, err
	}

	msg := wasmtypes.MsgStoreCode{
		Sender:                sender.String(),
		WASMByteCode:          wasm,
		InstantiatePermission: perm,
	}
	return msg, nil
}

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
			Symbol:        fmt.Sprintf("test%d", i),
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

func (s *KeeperTestSuite) TestFundMsgServerCreateActive() {
	ctx := s.chainA.GetContext()
	ctx, _, _, _ = s.setup(ctx)
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	srv := keeper.NewMsgServerImpl(s.GetDefundApp(s.chainA).EtfKeeper)
	wctx := sdk.WrapSDKContext(ctx)
	expected := &types.MsgCreateFund{
		Creator:       s.chainA.SenderAccount.GetAddress().String(),
		Name:          "Test 2",
		Description:   "Test 2",
		Symbol:        "test2",
		Holdings:      "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2:75:osmosis:1:spot,uosmo:25:osmosis:1:spot",
		BaseDenom:     "osmo",
		Rebalance:     10,
		StartingPrice: "10000000",
		Active:        true,
		WasmCodeId:    1,
	}
	sender := s.chainA.SenderAccount.GetAddress()
	uploadMsg, err := parseStoreCodeArgs("../../../tests/contracts/odd_number.wasm", sender)
	s.Assert().NoError(err)
	codeId, _, err := s.GetDefundApp(s.chainA).WasmInternalKeeper.Create(ctx, s.chainA.SenderAccount.GetAddress(), uploadMsg.WASMByteCode, uploadMsg.InstantiatePermission)
	s.Assert().NoError(err)
	expected.WasmCodeId = codeId
	_, err = srv.CreateFund(wctx, expected)
	s.Assert().NoError(err)
	// try running the runner within the active contract. This runner will be run at each rebalance
	fund, found := s.GetDefundApp(s.chainA).EtfKeeper.GetFund(ctx, "test2")
	s.Assert().True(found)
	fund.Rebalancing = false
	s.GetDefundApp(s.chainA).EtfKeeper.SetFund(ctx, fund)
	contractAddr, err := sdk.AccAddressFromBech32(fund.Contract)
	s.Assert().NoError(err)
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	_, err = s.GetDefundApp(s.chainA).WasmInternalKeeper.Execute(ctx, contractAddr, contractAddr, []byte(`{"runner": {}}`), sdk.Coins{})
	s.Assert().NoError(err)
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
	ctx, _, _, portId := s.setup(ctx)
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
	etfToken := sdk.NewCoin(fund.Shares.Denom, sdk.NewInt(5000000))
	err = s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), types.ModuleName, sdk.NewCoins(etfToken))
	s.Assert().NoError(err)
	err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(etfToken))
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

	// mock data for packet
	data := []byte{18, 219, 113, 12, 103, 107, 95, 216, 56, 143, 130, 159, 113, 176, 79, 128, 79, 214, 45, 220, 115, 169, 192, 84, 181, 42, 226, 211, 113, 13, 252, 109}
	// create a mock packet
	packet := channeltypes.NewPacket(data, 1, portId, "channel-1", "icahost", "channel-1", clienttypes.NewHeight(0, 10), 0)

	_, err = srv.Redeem(wctx, expectedRedeem)
	require.NoError(s.T(), err)
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	redeems := s.GetDefundApp(s.chainA).BrokerKeeper.GetAllRedeembySymbol(ctx, fund.Symbol)
	s.chainA.Log(redeems)
	s.GetDefundApp(s.chainA).EtfKeeper.OnRedeemSuccess(ctx, packet, redeems[0])
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	fund, found := s.GetDefundApp(s.chainA).EtfKeeper.GetFund(ctx, fund.Symbol)
	s.Assert().True(found)
	s.Assert().Equal(*fund.Shares, sdk.NewCoin(fund.Shares.Denom, sdk.NewInt(0)))
}

func (s *KeeperTestSuite) TestSharesMsgServerEditFund() {
	s.Run("Completed", func() {})

	s.Run("Unauthorized", func() {})
}
