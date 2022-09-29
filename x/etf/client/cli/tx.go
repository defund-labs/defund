package cli

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/cosmos/cosmos-sdk/client/flags"

	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	"github.com/defund-labs/defund/x/etf/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutHeight    = "packet-timeout-height"
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	flagAbsoluteTimeouts       = "absolute-timeouts"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateFund())
	cmd.AddCommand(CmdCreate())
	cmd.AddCommand(CmdRedeem())
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdCreateFund() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-fund [symbol] [name] [description] [basedenom] [broker] [holdings] [rebalance] [startingPrice]",
		Short: "Create a new fund",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			// Get value arguments
			argSymbol := args[0]
			argName := args[1]
			argDescription := args[2]
			argBaseDenom := args[3]
			argBroker := args[4]
			argHoldings := args[5]
			argRebalance := args[6]
			argStartingPrice := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Convert rebalance to int
			rebalance, err := strconv.ParseInt(argRebalance, 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateFund(
				clientCtx.GetFromAddress().String(),
				argSymbol,
				argName,
				argDescription,
				argBroker,
				argHoldings,
				rebalance,
				argBaseDenom,
				argStartingPrice,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [fund] [tokenIn] [channel]",
		Short: "Create shares for the dETF ticker using the IBC channel specified and the tokens supplied (comma seperated list of coins i.e 1000000uosmo,1000000uatom).",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFund := args[0]
			argTokenIn := args[1]
			argChannel := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			timeoutHeightStr, err := cmd.Flags().GetString(flagPacketTimeoutHeight)
			if err != nil {
				return err
			}
			timeoutHeight, err := clienttypes.ParseHeight(timeoutHeightStr)
			if err != nil {
				return err
			}

			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}

			tokenIn, err := sdk.ParseCoinNormalized(argTokenIn)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreate(
				clientCtx.GetFromAddress().String(),
				argFund,
				&tokenIn,
				argChannel,
				timeoutHeight.String(),
				timeoutTimestamp,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagPacketTimeoutHeight, types.DefaultRelativePacketTimeoutHeight, "Packet timeout block height. The timeout is disabled when set to 0-0.")
	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, types.DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds from now. Default is 10 minutes. The timeout is disabled when set to 0.")
	cmd.Flags().Bool(flagAbsoluteTimeouts, false, "Timeout flags are used as absolute timeouts.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRedeem() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redeem [fund] [amount] [channel] [addresses]",
		Short: "Redeem shares for the dETF ticker to supplied addresses using the IBC channel specified, and the tokens supplied. Sends redeemed assets to the addresses supplied i.e '{osmosisAddress: 'osmo123456789'}'",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFund := args[0]
			argAmount := args[1]
			argChannel := args[2]
			argAddresses := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(argAmount)
			if err != nil {
				return err
			}

			timeoutHeightStr, err := cmd.Flags().GetString(flagPacketTimeoutHeight)
			if err != nil {
				return err
			}
			timeoutHeight, err := clienttypes.ParseHeight(timeoutHeightStr)
			if err != nil {
				return err
			}

			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}

			addresses := parseJSONAddressMap(argAddresses)

			msg := types.NewMsgRedeem(
				clientCtx.GetFromAddress().String(),
				argFund,
				&amount,
				argChannel,
				timeoutHeight.String(),
				timeoutTimestamp,
				addresses,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagPacketTimeoutHeight, types.DefaultRelativePacketTimeoutHeight, "Packet timeout block height. The timeout is disabled when set to 0-0.")
	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, types.DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds from now. Default is 10 minutes. The timeout is disabled when set to 0.")
	cmd.Flags().Bool(flagAbsoluteTimeouts, false, "Timeout flags are used as absolute timeouts.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// parseJSONAddressMap takes a json string and converts it into JSON address map
func parseJSONAddressMap(jsonString string) types.AddressMap {
	data := types.AddressMap{}
	json.Unmarshal([]byte(jsonString), &data)

	return data
}
