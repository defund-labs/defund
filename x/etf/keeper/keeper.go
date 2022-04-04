package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/etf/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		accountKeeper: accountKeeper, bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// End keeper that runs at each block. This end keeper checks to see if there is any balance left
// in the funds invest account (on the opposite chain). If there is a balance, the end keeper
// creates ICA instructions to swap funds approporiately for the underlying fund assets. Once complete
// assets are sent back to Defunds fund holding account and investment is complete.
func (k Keeper) InvestEndBlocker() {
}

// End keeper that runs at each block. This end keeper checks to see if there is any balance left
// in the funds uninvest account (on the opposite chain). If there is a balance, the end keeper
// creates ICA instructions to swap funds to ATOM or DETF to redeem for the investor. Once complete
// assets are sent back to Defund's fund uninvest account and assets are sent to redeemer.
func (k Keeper) UninvestEndBlocker() {
}
