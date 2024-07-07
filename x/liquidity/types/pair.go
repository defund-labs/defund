package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cometbft/cometbft/crypto"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

// AddressType enumerates the available types of a address.
type AddressType int32

const (
	// the 32 bytes length address type of ADR 028.
	AddressType32Bytes AddressType = 0
	// the default 20 bytes length address type.
	AddressType20Bytes AddressType = 1
)

func (pair Pair) GetEscrowAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(pair.EscrowAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewPair returns a new pair object.
func NewPair(id uint64, baseCoinDenom, quoteCoinDenom string) Pair {
	return Pair{
		Id:             id,
		BaseCoinDenom:  baseCoinDenom,
		QuoteCoinDenom: quoteCoinDenom,
		EscrowAddress:  PairEscrowAddress(id).String(),
		LastOrderId:    0,
		LastPrice:      nil,
		CurrentBatchId: 1,
	}
}

// Validate validates Pair for genesis.
func (pair Pair) Validate() error {
	if pair.Id == 0 {
		return fmt.Errorf("pair id must not be 0")
	}
	if err := sdk.ValidateDenom(pair.BaseCoinDenom); err != nil {
		return fmt.Errorf("invalid base coin denom: %w", err)
	}
	if err := sdk.ValidateDenom(pair.QuoteCoinDenom); err != nil {
		return fmt.Errorf("invalid quote coin denom: %w", err)
	}
	if _, err := sdk.AccAddressFromBech32(pair.EscrowAddress); err != nil {
		return fmt.Errorf("invalid escrow address %s: %w", pair.EscrowAddress, err)
	}
	if pair.LastPrice != nil {
		if !pair.LastPrice.IsPositive() {
			return fmt.Errorf("last price must be positive: %s", pair.LastPrice)
		}
	}
	if pair.CurrentBatchId == 0 {
		return fmt.Errorf("current batch id must not be 0")
	}
	return nil
}

// DeriveAddress derives an address with the given address length type, module name, and
// address derivation name. It is used to derive private plan farming pool address, and staking reserve address.
func DeriveAddress(addressType AddressType, moduleName, name string) sdk.AccAddress {
	switch addressType {
	case AddressType32Bytes:
		return sdk.AccAddress(address.Module(moduleName, []byte(name)))
	case AddressType20Bytes:
		return sdk.AccAddress(crypto.AddressHash([]byte(moduleName + name)))
	default:
		return sdk.AccAddress{}
	}
}

// PairEscrowAddress returns a unique address of the pair's escrow.
func PairEscrowAddress(pairId uint64) sdk.AccAddress {
	return DeriveAddress(
		AddressType32Bytes,
		ModuleName,
		strings.Join([]string{PairEscrowAddressPrefix, strconv.FormatUint(pairId, 10)}, ModuleAddressNameSplitter))
}

// MustMarshalPair returns the pair bytes.
// It throws panic if it fails.
func MustMarshalPair(cdc codec.BinaryCodec, pair Pair) []byte {
	return cdc.MustMarshal(&pair)
}

// MustUnmarshalPair return the unmarshalled pair from bytes.
// It throws panic if it fails.
func MustUnmarshalPair(cdc codec.BinaryCodec, value []byte) Pair {
	pair, err := UnmarshalPair(cdc, value)
	if err != nil {
		panic(err)
	}

	return pair
}

// UnmarshalPair returns the pair from bytes.
func UnmarshalPair(cdc codec.BinaryCodec, value []byte) (pair Pair, err error) {
	err = cdc.Unmarshal(value, &pair)
	return pair, err
}
