package wasmbinding

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	etfwasm "github.com/defund-labs/defund/x/etf/client/wasm"
)

type WasmMessage struct {
	EditFund json.RawMessage `json:"edit_fund,omitempty"`
}

func CustomEncoder(sender sdk.AccAddress, msg json.RawMessage) ([]sdk.Msg, error) {
	var parsedMessage WasmMessage
	if err := json.Unmarshal(msg, &parsedMessage); err != nil {
		return []sdk.Msg{}, sdkerrors.Wrap(err, "Error parsing Wasm Message")
	}
	switch {
	case parsedMessage.EditFund != nil:
		return etfwasm.EncodeEditFund(parsedMessage.EditFund, sender)
	default:
		return []sdk.Msg{}, wasmvmtypes.UnsupportedRequest{Kind: "Unknown Wasm Message"}
	}
}
