package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	"github.com/defund-labs/defund/x/etf/types"
)

func (s *KeeperTestSuite) TestCallbacks() {

	path := s.NewTransferPath()
	s.Require().Equal(path.EndpointA.ChannelID, "channel-0")
	fund, _, _, _ := s.CreateTestFund(path)
	fund, err := s.GetDefundApp(s.chainA).EtfKeeper.GetFundBySymbol(s.chainA.GetContext(), fund.Symbol)
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

	s.Run("OnTransferSuccess", func() {
		currentFund, found := s.GetDefundApp(s.chainA).EtfKeeper.GetFund(s.chainA.GetContext(), fund.Symbol)
		s.Assert().True(found)
		beforeShares := currentFund.Shares.Amount
		tokenIn := sdk.NewCoin(fund.BaseDenom.OnDefund, sdk.NewInt(44565793))
		err := s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), "etf", sdk.NewCoins(tokenIn))
		s.Assert().NoError(err)
		err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(tokenIn))
		s.Assert().NoError(err)
		err = s.GetDefundApp(s.chainA).EtfKeeper.CreateShares(s.chainA.GetContext(), fund, "channel-0", tokenIn, s.chainA.SenderAccounts[1].SenderAccount.GetAddress().String(), clienttypes.NewHeight(0, 100), 0)
		s.Assert().NoError(err)
		// query the transfers and get the one for the create shares above
		transfers := s.GetDefundApp(s.chainA).BrokerKeeper.GetAllTransfer(s.chainA.GetContext())
		data := transfertypes.NewFungibleTokenPacketData(transfers[0].Token.Denom, transfers[0].Token.Amount.String(), transfers[0].Sender, transfers[0].Receiver)
		packet := channeltypes.NewPacket(data.GetBytes(), transfers[0].Sequence, "transfer", path.EndpointA.ChannelID, "transfer", path.EndpointB.ChannelID, clienttypes.NewHeight(0, 10), 0)
		ack := channeltypes.NewResultAcknowledgement([]byte("success"))
		err = s.GetDefundApp(s.chainA).EtfKeeper.OnTransferSuccess(s.chainA.GetContext(), packet, ack)
		s.Assert().NoError(err)
		currentFund, found = s.GetDefundApp(s.chainA).EtfKeeper.GetFund(s.chainA.GetContext(), fund.Symbol)
		s.Assert().True(found)
		afterShares := currentFund.Shares.Amount
		// make sure that the shares of funds are not equal since we should have created more on success
		s.Assert().NotEqual(beforeShares, afterShares)
		// make sure the creator received created shares in fund
		etfBalance := s.GetDefundApp(s.chainA).BankKeeper.GetBalance(s.chainA.GetContext(), s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), fund.Shares.Denom)
		s.Assert().Equal(sdk.NewCoin(fund.Shares.Denom, sdk.NewInt(340000)), etfBalance)
	})

	s.Run("OnTransferFailure", func() {
		currentFund, found := s.GetDefundApp(s.chainA).EtfKeeper.GetFund(s.chainA.GetContext(), fund.Symbol)
		s.Assert().True(found)
		beforeShares := currentFund.Shares.Amount
		tokenIn := sdk.NewCoin(fund.BaseDenom.OnDefund, sdk.NewInt(44565793))
		err := s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), "etf", sdk.NewCoins(tokenIn))
		s.Assert().NoError(err)
		err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(tokenIn))
		s.Assert().NoError(err)
		err = s.GetDefundApp(s.chainA).EtfKeeper.CreateShares(s.chainA.GetContext(), fund, "channel-0", tokenIn, s.chainA.SenderAccounts[1].SenderAccount.GetAddress().String(), clienttypes.NewHeight(0, 100), 0)
		s.Assert().NoError(err)
		// query the transfers and get the one for the create shares above
		transfers := s.GetDefundApp(s.chainA).BrokerKeeper.GetAllTransfer(s.chainA.GetContext())
		data := transfertypes.NewFungibleTokenPacketData(transfers[0].Token.Denom, transfers[0].Token.Amount.String(), transfers[0].Sender, transfers[0].Receiver)
		packet := channeltypes.NewPacket(data.GetBytes(), transfers[0].Sequence, "transfer", path.EndpointA.ChannelID, "transfer", path.EndpointB.ChannelID, clienttypes.NewHeight(0, 10), 0)
		ack := channeltypes.NewErrorAcknowledgement(sdkerrors.New("create error", 1000, "create error"))
		err = s.GetDefundApp(s.chainA).EtfKeeper.OnTransferFailure(s.chainA.GetContext(), packet, ack)
		s.Assert().NoError(err)
		currentFund, found = s.GetDefundApp(s.chainA).EtfKeeper.GetFund(s.chainA.GetContext(), fund.Symbol)
		s.Assert().True(found)
		afterShares := currentFund.Shares.Amount
		// make sure that the shares of funds are equal since we should not have created more on failure
		s.Assert().Equal(beforeShares, afterShares)
		// make sure the creator received there escrowed tokenin back
		err = s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), "etf", sdk.NewCoins(sdk.NewCoin("ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518", sdk.NewInt(14706711))))
		s.Assert().NoError(err)
		escrowAddress := transfertypes.GetEscrowAddress(packet.GetSourcePort(), packet.GetSourceChannel())
		err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), "etf", escrowAddress, sdk.NewCoins(sdk.NewCoin("ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518", sdk.NewInt(14706711))))
		s.Assert().NoError(err)
		module, _, err := s.GetDefundApp(s.chainA).GetIBCKeeper().PortKeeper.LookupModuleByPort(s.chainA.GetContext(), "transfer")
		s.Assert().NoError(err)
		ibcModule, ok := s.GetDefundApp(s.chainA).GetIBCKeeper().Router.GetRoute(module)
		s.Assert().True(ok)
		relayer := s.chainA.SenderAccounts[2].SenderAccount.GetAddress()
		err = ibcModule.OnAcknowledgementPacket(s.chainA.GetContext(), packet, ack.Acknowledgement(), relayer)
		s.Assert().NoError(err)
		tokenInBalance := s.GetDefundApp(s.chainA).BankKeeper.GetBalance(s.chainA.GetContext(), s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), tokenIn.Denom)
		// we add 4 to the check since the balance starts with 4 for some reason
		s.Assert().Equal(sdk.NewCoin("ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518", sdk.NewInt(14706715)), tokenInBalance)
	})

	s.Run("OnTransferTimeout", func() {
		currentFund, found := s.GetDefundApp(s.chainA).EtfKeeper.GetFund(s.chainA.GetContext(), fund.Symbol)
		s.Assert().True(found)
		beforeShares := currentFund.Shares.Amount
		tokenIn := sdk.NewCoin(fund.BaseDenom.OnDefund, sdk.NewInt(44565793))
		err := s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), "etf", sdk.NewCoins(tokenIn))
		s.Assert().NoError(err)
		err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), types.ModuleName, s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), sdk.NewCoins(tokenIn))
		s.Assert().NoError(err)
		err = s.GetDefundApp(s.chainA).EtfKeeper.CreateShares(s.chainA.GetContext(), fund, "channel-0", tokenIn, s.chainA.SenderAccounts[1].SenderAccount.GetAddress().String(), clienttypes.NewHeight(0, 100), 0)
		s.Assert().NoError(err)
		// query the transfers and get the one for the create shares above
		transfers := s.GetDefundApp(s.chainA).BrokerKeeper.GetAllTransfer(s.chainA.GetContext())
		data := transfertypes.NewFungibleTokenPacketData(transfers[0].Token.Denom, transfers[0].Token.Amount.String(), transfers[0].Sender, transfers[0].Receiver)
		packet := channeltypes.NewPacket(data.GetBytes(), transfers[0].Sequence, "transfer", path.EndpointA.ChannelID, "transfer", path.EndpointB.ChannelID, clienttypes.NewHeight(0, 10), 0)
		ack := channeltypes.NewErrorAcknowledgement(sdkerrors.New("create error", 1000, "create error"))
		err = s.GetDefundApp(s.chainA).EtfKeeper.OnTransferTimeout(s.chainA.GetContext(), packet)
		s.Assert().NoError(err)
		currentFund, found = s.GetDefundApp(s.chainA).EtfKeeper.GetFund(s.chainA.GetContext(), fund.Symbol)
		s.Assert().True(found)
		afterShares := currentFund.Shares.Amount
		// make sure that the shares of funds are equal since we should not have created more on failure
		s.Assert().Equal(beforeShares, afterShares)
		// make sure the creator received there escrowed tokenin back
		err = s.GetDefundApp(s.chainA).BankKeeper.MintCoins(s.chainA.GetContext(), "etf", sdk.NewCoins(sdk.NewCoin("ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518", sdk.NewInt(14706711))))
		s.Assert().NoError(err)
		escrowAddress := transfertypes.GetEscrowAddress(packet.GetSourcePort(), packet.GetSourceChannel())
		err = s.GetDefundApp(s.chainA).BankKeeper.SendCoinsFromModuleToAccount(s.chainA.GetContext(), "etf", escrowAddress, sdk.NewCoins(sdk.NewCoin("ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518", sdk.NewInt(14706711))))
		s.Assert().NoError(err)
		module, _, err := s.GetDefundApp(s.chainA).GetIBCKeeper().PortKeeper.LookupModuleByPort(s.chainA.GetContext(), "transfer")
		s.Assert().NoError(err)
		ibcModule, ok := s.GetDefundApp(s.chainA).GetIBCKeeper().Router.GetRoute(module)
		s.Assert().True(ok)
		relayer := s.chainA.SenderAccounts[2].SenderAccount.GetAddress()
		err = ibcModule.OnAcknowledgementPacket(s.chainA.GetContext(), packet, ack.Acknowledgement(), relayer)
		s.Assert().NoError(err)
		tokenInBalance := s.GetDefundApp(s.chainA).BankKeeper.GetBalance(s.chainA.GetContext(), s.chainA.SenderAccounts[1].SenderAccount.GetAddress(), tokenIn.Denom)
		// we add 4 to the check since the balance starts with 4 for some reason
		s.Assert().Equal(sdk.NewCoin("ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518", sdk.NewInt(29413428)), tokenInBalance)
	})
}
