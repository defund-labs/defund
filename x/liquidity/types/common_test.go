package types_test

import (
	sdk "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

var testAddr = types.AccAddress(crypto.AddressHash([]byte("test")))

func newInt(i int64) sdk.Int {
	return sdk.NewInt(i)
}
