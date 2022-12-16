package wasmbinding_test

func (s *WasmTestSuite) TestWasmEtfQuery() {
	s.chainA.Log(s.GetDefundApp(s.chainA).EtfKeeper.GetAllFund(s.chainA.GetContext()))
}

func (s *WasmTestSuite) TestWasmBrokerQuery() {}
