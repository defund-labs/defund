package amm

import "cosmossdk.io/math"

// The minimum and maximum coin amount used in the amm package.
var (
	MinCoinAmount = math.NewInt(100)
	MaxCoinAmount = math.NewIntWithDecimal(1, 40)
)

var (
	MinPoolPrice               = math.LegacyNewDecWithPrec(1, 15)            // 10^-15
	MaxPoolPrice               = math.NewIntWithDecimal(1, 20).ToLegacyDec() // 10^20
	MinRangedPoolPriceGapRatio = math.LegacyNewDecWithPrec(1, 3)             // 0.001, 0.1%
)
