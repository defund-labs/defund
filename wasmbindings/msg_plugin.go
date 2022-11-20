package wasmbinding

import (
	"encoding/json"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	etfwasm "github.com/defund-labs/defund/x/etf/client/wasm"
	etftypes "github.com/defund-labs/defund/x/etf/types"
)

// CustomMessageDecorator returns decorator for custom CosmWasm bindings messages
func CustomMessageDecorator(
	router wasmkeeper.MessageRouter,
	accountKeeper *authkeeper.AccountKeeper,
) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			router:        router,
			wrapped:       old,
			accountKeeper: accountKeeper,
		}
	}
}

type CustomMessenger struct {
	router        wasmkeeper.MessageRouter
	wrapped       wasmkeeper.Messenger
	accountKeeper *authkeeper.AccountKeeper
}

type DefundWasmMessage struct {
	EditFund json.RawMessage `json:"edit_fund,omitempty"`
}

var _ wasmkeeper.Messenger = &CustomMessenger{}

// DispatchMsg executes on the bindingMsgs
func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		return m.DispatchCustomMsg(ctx, contractAddr, contractIBCPortID, msg)
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

// DispatchCustomMsg function is forked from wasmd. sdk.Msg will be validated and routed to the corresponding module msg server in this function.
func (m *CustomMessenger) DispatchCustomMsg(
	ctx sdk.Context,
	contractAddr sdk.AccAddress,
	contractIBCPortID string,
	msg wasmvmtypes.CosmosMsg,
) (events []sdk.Event, data [][]byte, err error) {
	var parsedMessage DefundWasmMessage
	if err := json.Unmarshal(msg.Custom, &parsedMessage); err != nil {
		return nil, nil, etftypes.ErrParsingWasmMsg
	}

	var sdkMsgs []sdk.Msg
	switch {
	case parsedMessage.EditFund != nil:
		sdkMsgs, err = etfwasm.EncodeEditFund(parsedMessage.EditFund, contractAddr)
	default:
		sdkMsgs, err = []sdk.Msg{}, wasmvmtypes.UnsupportedRequest{Kind: "Unknown Sei Wasm Message"}
	}
	if err != nil {
		return nil, nil, err
	}

	for _, sdkMsg := range sdkMsgs {
		res, err := m.handleSdkMessage(ctx, contractAddr, sdkMsg)
		if err != nil {
			return nil, nil, err
		}
		// append data
		data = append(data, res.Data)
		// append events
		sdkEvents := make([]sdk.Event, len(res.Events))
		for i := range res.Events {
			sdkEvents[i] = sdk.Event(res.Events[i])
		}
		events = append(events, sdkEvents...)
	}
	return events, data, nil
}

// This function is forked from wasmd. sdk.Msg will be validated and routed to the corresponding module msg server in this function.
func (m *CustomMessenger) handleSdkMessage(ctx sdk.Context, contractAddr sdk.Address, msg sdk.Msg) (*sdk.Result, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	// make sure this account can send it
	for _, acct := range msg.GetSigners() {
		if !acct.Equals(contractAddr) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "contract doesn't have permission")
		}
	}

	// find the handler and execute it
	if handler := m.router.Handler(msg); handler != nil {
		// ADR 031 request type routing
		msgResult, err := handler(ctx, msg)
		return msgResult, err
	}
	// legacy sdk.Msg routing
	// Assuming that the app developer has migrated all their Msgs to
	// proto messages and has registered all `Msg services`, then this
	// path should never be called, because all those Msgs should be
	// registered within the `msgServiceRouter` already.
	return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "can't route message %+v", msg)
}
