package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/proto/tendermint/crypto"

	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	etftypes "github.com/defund-labs/defund/x/etf/types"
	"github.com/defund-labs/defund/x/query/keeper"
	"github.com/defund-labs/defund/x/query/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func (s *KeeperTestSuite) setup(ctx sdk.Context) (outctx sdk.Context, fund etftypes.Fund, connectionId string, portId string) {
	path := s.NewTransferPath()
	s.Require().Equal(path.EndpointA.ChannelID, "channel-0")

	// Commit new block to store info
	s.coordinator.CommitBlock(s.chainA, s.chainB)

	outctx = ctx

	return outctx, fund, connectionId, portId
}

func (s *KeeperTestSuite) TestInterqueryMsgServerResult() {
	creator := "A"

	k := s.GetDefundApp(s.chainA).QueryKeeper
	ctx := s.chainA.GetContext()
	ctx, _, _, _ = s.setup(ctx)
	s.coordinator.CommitBlock(s.chainA, s.chainB)
	srv := keeper.NewMsgServerImpl(k)
	wctx := sdk.WrapSDKContext(ctx)

	h := clienttypes.NewHeight(0, 0)

	s.coordinator.CommitBlock(s.chainA, s.chainB)

	for _, tc := range []struct {
		desc    string
		request *types.MsgCreateInterqueryResult
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgCreateInterqueryResult{Creator: creator,
				Storeid: strconv.Itoa(0),
				Proof:   &crypto.ProofOps{},
				Height:  &h,
			},
		},
		{
			desc: "Invalid proof",
			request: &types.MsgCreateInterqueryResult{Creator: creator,
				Storeid: strconv.Itoa(0),
				Proof:   &crypto.ProofOps{},
				Height:  &h,
			},
			err: sdkerrors.Wrapf(types.ErInvalidProof, "no proof provided"),
		},
		{
			desc: "InterqueryNotFound",
			request: &types.MsgCreateInterqueryResult{Creator: creator,
				Storeid: strconv.Itoa(100000),
				Proof:   &crypto.ProofOps{},
				Height:  &h,
			},
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Interquery with StoreId %s could not be found.", strconv.Itoa(0))),
		},
	} {
		s.T().Run(tc.desc, func(t *testing.T) {
			expected := &types.MsgCreateInterqueryResult{Creator: creator,
				Storeid: strconv.Itoa(0),
			}

			iq := types.Interquery{
				Storeid:      tc.request.Storeid,
				ConnectionId: "connection-0",
			}
			k.SetInterquery(ctx, iq)

			_, err := srv.CreateInterqueryResult(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetInterquery(ctx,
					expected.Storeid,
				)
				require.True(t, found)
				require.Equal(t, expected.Storeid, rst.Storeid)
			}
		})
	}
}
