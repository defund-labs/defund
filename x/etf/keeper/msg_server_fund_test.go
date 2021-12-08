package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/defundhub/defund/testutil/keeper"
	"github.com/defundhub/defund/x/etf/keeper"
	"github.com/defundhub/defund/x/etf/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestFundMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.EtfKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateFund{Creator: creator,
			Id: strconv.Itoa(i),
		}
		_, err := srv.CreateFund(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetFund(ctx,
			expected.Id,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestFundMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateFund
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateFund{Creator: creator,
				Id: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateFund{Creator: "B",
				Id: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateFund{Creator: creator,
				Id: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EtfKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateFund{Creator: creator,
				Id: strconv.Itoa(0),
			}
			_, err := srv.CreateFund(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateFund(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetFund(ctx,
					expected.Id,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}
