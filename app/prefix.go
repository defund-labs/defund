package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAddressConfig sets the address configuration.
func SetAddressConfig(AccountAddressPrefix string) {
	var (
		AccountPubKeyPrefix    = AccountAddressPrefix + "pub"
		ValidatorAddressPrefix = AccountAddressPrefix + "valoper"
		ValidatorPubKeyPrefix  = AccountAddressPrefix + "valoperpub"
		ConsNodeAddressPrefix  = AccountAddressPrefix + "valcons"
		ConsNodePubKeyPrefix   = AccountAddressPrefix + "valconspub"
	)

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	config.Seal()
}
