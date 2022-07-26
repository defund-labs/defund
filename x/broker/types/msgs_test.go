package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgAddLiquiditySource_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddLiquiditySource
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddLiquiditySource{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddLiquiditySource{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgAddConnectionBroker_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddConnectionBroker
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddConnectionBroker{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddConnectionBroker{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
