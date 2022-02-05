package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defundhub/defund/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateInterquery_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateInterquery
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateInterquery{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateInterquery{
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

func TestMsgUpdateInterquery_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateInterquery
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateInterquery{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateInterquery{
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

func TestMsgDeleteInterquery_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteInterquery
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteInterquery{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteInterquery{
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
