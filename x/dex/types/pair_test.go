package types_test

import (
	"testing"

	math "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"defund/x/dex/types"
)

func TestPair_Validate(t *testing.T) {
	for _, tc := range []struct {
		name        string
		malleate    func(pair *types.Pair)
		expectedErr string
	}{
		{
			"happy case",
			func(pair *types.Pair) {},
			"",
		},
		{
			"zero id",
			func(pair *types.Pair) {
				pair.Id = 0
			},
			"pair id must not be 0",
		},
		{
			"invalid base coin denom",
			func(pair *types.Pair) {
				pair.BaseCoinDenom = "invalliddenom!"
			},
			"invalid base coin denom: invalid denom: invalliddenom!",
		},
		{
			"invalid quote coin denom",
			func(pair *types.Pair) {
				pair.QuoteCoinDenom = "invaliddenom!"
			},
			"invalid quote coin denom: invalid denom: invaliddenom!",
		},
		{
			"invalid escrow address",
			func(pair *types.Pair) {
				pair.EscrowAddress = "invalidaddr"
			},
			"invalid escrow address invalidaddr: decoding bech32 failed: invalid separator index -1",
		},
		{
			"",
			func(pair *types.Pair) {
				p := math.LegacyNewDec(-1)
				pair.LastPrice = &p
			},
			"last price must be positive: -1.000000000000000000",
		},
		{
			"",
			func(pair *types.Pair) {
				pair.CurrentBatchId = 0
			},
			"current batch id must not be 0",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pair := types.NewPair(1, "denom1", "denom2")
			tc.malleate(&pair)
			err := pair.Validate()
			if tc.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}

func TestPairEscrowAddress(t *testing.T) {
	for _, tc := range []struct {
		pairId   uint64
		expected string
	}{
		{1, "cosmos1xutvc4z4z4gvqryesqk0f555gw7ggd7wtpzcrl0300g4vz5g4mtqvv554k"},
		{2, "cosmos1anyy4u9ufk0xmac84nl5ae4rgyhryql4c68vxulmaxdrxs9aa3gqnya8q0"},
	} {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tc.expected, types.PairEscrowAddress(tc.pairId).String())
		})
	}
}
