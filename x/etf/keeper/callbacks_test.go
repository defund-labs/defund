package keeper_test

import (
	"encoding/base64"
	"fmt"

	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	querytypes "github.com/defund-labs/defund/x/query/types"
)

func (s *KeeperTestSuite) TestICQCallbacks() {

	path := s.NewTransferPath()
	s.Require().Equal(path.EndpointA.ChannelID, "channel-0")

	fund, _, _, _ := s.CreateTestFund(path)

	s.Run("OnFundBalanceSubmissionCallback", func() {

		portID, err := icatypes.NewControllerPortID(fund.Address)
		s.Assert().NoError(err)
		address, found := s.GetDefundApp(s.chainA).BrokerKeeper.GetBrokerAccount(s.chainA.GetContext(), "connection-0", portID)
		s.Assert().True(found)
		testQueryId := fmt.Sprintf("balance:test:osmosis:%s:uosmo", address)
		height := clienttypes.NewHeight(4, 7266410)
		data, err := base64.StdEncoding.DecodeString("CgV1b3NtbxIIMTAwMDAwMDA=")
		s.Assert().NoError(err)
		result := querytypes.InterqueryResult{
			Creator:     "defund1mjk79fjjgpplak5wq838w0yd982gzkyftuze76",
			Storeid:     testQueryId,
			Chainid:     "osmo-test-4",
			Data:        data,
			Height:      &height,
			LocalHeight: 0,
			Success:     true,
			Proved:      true,
		}

		err = s.GetDefundApp(s.chainA).EtfKeeper.OnFundBalanceSubmissionCallback(s.chainA.GetContext(), &result)
		s.Assert().NoError(err)
		fund, found := s.GetDefundApp(s.chainA).EtfKeeper.GetFund(s.chainA.GetContext(), "test")
		s.Assert().True(found)
		s.coordinator.Log(fund)
	})
}
