package keeper_test

func (s *KeeperTestSuite) TestCallbacks() {

	path := s.NewTransferPath()
	s.Require().Equal(path.EndpointA.ChannelID, "channel-0")

	s.Run("OnTransferSuccess", func() {})

	s.Run("OnTransferFailure", func() {})

	s.Run("OnTransferTimeout", func() {})

	s.Run("OnAcknowledgementPacket", func() {})

	s.Run("OnTimeoutPacket", func() {})
}
