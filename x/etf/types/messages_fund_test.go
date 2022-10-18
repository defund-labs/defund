package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateFund_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateFund
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateFund{
				Creator:   "invalid_address",
				BaseDenom: "uosmo",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address",
			msg: MsgCreateFund{
				Creator:   sample.AccAddress(),
				BaseDenom: "uosmo",
			},
		},
		{
			name: "invalid base denom",
			msg: MsgCreateFund{
				Creator:   sample.AccAddress(),
				BaseDenom: "ujuno",
			},
			err: sdkerrors.Wrapf(ErrWrongBaseDenom, "invalid base denom (ujuno). must be uosmo or ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2 (uatom)"),
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
