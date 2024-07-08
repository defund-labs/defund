package amm

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "cosmossdk.io/math"

	utils "defund/types"
)

func Test_char(t *testing.T) {
	require.Panics(t, func() {
		char(sdk.LegacyZeroDec())
	})

	for _, tc := range []struct {
		x        sdk.LegacyDec
		expected int
	}{
		{sdk.LegacyMustNewDecFromStr("999.99999999999999999"), 20},
		{sdk.LegacyMustNewDecFromStr("100"), 20},
		{sdk.LegacyMustNewDecFromStr("99.999999999999999999"), 19},
		{sdk.LegacyMustNewDecFromStr("10"), 19},
		{sdk.LegacyMustNewDecFromStr("9.999999999999999999"), 18},
		{sdk.LegacyMustNewDecFromStr("1"), 18},
		{sdk.LegacyMustNewDecFromStr("0.999999999999999999"), 17},
		{sdk.LegacyMustNewDecFromStr("0.1"), 17},
		{sdk.LegacyMustNewDecFromStr("0.099999999999999999"), 16},
		{sdk.LegacyMustNewDecFromStr("0.01"), 16},
		{sdk.LegacyMustNewDecFromStr("0.000000000000000009"), 0},
		{sdk.LegacyMustNewDecFromStr("0.000000000000000001"), 0},
	} {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tc.expected, char(tc.x))
		})
	}
}

func Test_pow10(t *testing.T) {
	for _, tc := range []struct {
		power    int
		expected sdk.LegacyDec
	}{
		{18, sdk.LegacyNewDec(1)},
		{19, sdk.LegacyNewDec(10)},
		{20, sdk.LegacyNewDec(100)},
		{17, sdk.LegacyNewDecWithPrec(1, 1)},
		{16, sdk.LegacyNewDecWithPrec(1, 2)},
	} {
		t.Run("", func(t *testing.T) {
			require.True(sdk.LegacyDecEq(t, tc.expected, pow10(tc.power)))
		})
	}
}

func Test_isPow10(t *testing.T) {
	for _, tc := range []struct {
		x        sdk.LegacyDec
		expected bool
	}{
		{utils.ParseDec("100"), true},
		{utils.ParseDec("101"), false},
		{utils.ParseDec("10"), true},
		{utils.ParseDec("1"), true},
		{utils.ParseDec("1.000000000000000001"), false},
		{utils.ParseDec("0.11"), false},
		{utils.ParseDec("0.000000000000000001"), true},
		{utils.ParseDec("10000000000000000000000000001"), false},
		{utils.ParseDec("10000000000000000000000000000"), true},
		{utils.ParseDec("123456789"), false},
	} {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tc.expected, isPow10(tc.x))
		})
	}
}
