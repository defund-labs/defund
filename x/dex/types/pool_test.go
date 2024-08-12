package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"defund/x/dex/types"
)

func TestPoolReserveAddress(t *testing.T) {
	for _, tc := range []struct {
		poolId   uint64
		expected string
	}{
		{1, "cosmos1n4tt7fvpjhduvyhj2rezm0m7gdwuuf08azrszdalmn9kle66j8xqa2t0fa"},
		{2, "cosmos167wmk7wkd7u7hevqy75t8jc6kwsgv6xngx74908eulqcklhpykzqqjahag"},
	} {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tc.expected, types.PoolReserveAddress(tc.poolId).String())
		})
	}
}

func TestPoolCoinDenom(t *testing.T) {
	for _, tc := range []struct {
		poolId   uint64
		expected string
	}{
		{1, "pool1"},
		{10, "pool10"},
		{18446744073709551615, "pool18446744073709551615"},
	} {
		t.Run("", func(t *testing.T) {
			poolCoinDenom := types.PoolCoinDenom(tc.poolId)
			require.Equal(t, tc.expected, poolCoinDenom)
		})
	}
}

func TestParsePoolCoinDenomFailure(t *testing.T) {
	for _, tc := range []struct {
		denom      string
		expectsErr bool
	}{
		{"pool1", false},
		{"pool10", false},
		{"pool18446744073709551615", false},
		{"pool18446744073709551616", true},
		{"pool01", true},
		{"pool-10", true},
		{"pool+10", true},
		{"ucre", true},
		{"denom1", true},
	} {
		t.Run("", func(t *testing.T) {
			poolId, err := types.ParsePoolCoinDenom(tc.denom)
			if tc.expectsErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.denom, types.PoolCoinDenom(poolId))
			}
		})
	}
}
