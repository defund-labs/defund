package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	ibctesting "github.com/defund-labs/defund/testing"
)

func (s *KeeperTestSuite) TestETFTransfer_Valid() {
	s.SetupTest()

	// setup transfer channels
	path := ibctesting.NewPath(s.chainA, s.chainB)
	path.EndpointA.ChannelConfig.PortID = "transfer"
	path.EndpointB.ChannelConfig.PortID = "transfer"
	path.EndpointA.ChannelConfig.Version = "ics20-1"
	path.EndpointB.ChannelConfig.Version = "ics20-1"
	path.EndpointA.ConnectionID = "connection-0"
	path.EndpointB.ConnectionID = "connection-0"
	path.EndpointA.ChannelID = "channel-0"
	path.EndpointB.ChannelID = "channel-0"
	path.EndpointA.ClientID = "07-tendermint-0"
	path.EndpointB.ClientID = "07-tendermint-0"
	s.coordinator.SetupConnections(path)
	s.coordinator.CreateChannels(path)

	_, err := s.GetDefundApp(s.chainA).EtfKeeper.SendTransfer(s.chainA.GetContext(), "channel-0", sdk.NewCoin("stake", sdk.NewInt(1000000)), s.chainA.SenderAccount.GetAddress().String(), s.chainB.SenderAccount.GetAddress().String(), clienttypes.NewHeight(0, 100), 0)
	s.Assert().NoError(err)
}
