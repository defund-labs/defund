package wasm

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/etf/types"
)

type EditFund struct {
	Symbol   string           `json:"symbol"`
	Holdings []*types.Holding `json:"holdings"`
}

func EncodeEditFund(rawMsg json.RawMessage, sender sdk.AccAddress) ([]sdk.Msg, error) {
	decodedMsg := EditFund{}
	if err := json.Unmarshal(rawMsg, &decodedMsg); err != nil {
		return []sdk.Msg{}, types.ErrUnmarshallJson
	}
	msg := types.MsgEditFund{
		Creator:  sender.String(),
		Symbol:   decodedMsg.Symbol,
		Holdings: decodedMsg.Holdings,
	}
	return []sdk.Msg{&msg}, nil
}
