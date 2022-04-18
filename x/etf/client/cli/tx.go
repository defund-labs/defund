package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/cosmos/cosmos-sdk/client/flags"

	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
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
	cmd.AddCommand(CmdInvest())
	cmd.AddCommand(CmdUninvest())
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdCreateFund() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-fund [symbol] [name] [description] [basedenom] [broker] [holdings] [rebalance] [connectionId]",
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
			argConnection := args[7]

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
				argConnection,
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

func CmdInvest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invest [fund] [amount] [channel]",
		Short: "Invest the specified amount into the dETF ticker using the IBC channel specified.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFund := args[0]
			argAmount := args[1]
			argChannel := args[2]

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

			msg := types.NewMsgInvest(
				clientCtx.GetFromAddress().String(),
				argFund,
				&amount,
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

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUninvest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninvest [fund] [amount]",
		Short: "Uninvest the specified amount from the dETF specified.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFund := args[0]
			argAmount := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(argAmount)
			if err != nil {
				return err
			}

			msg := types.NewMsgUninvest(
				clientCtx.GetFromAddress().String(),
				argFund,
				&amount,
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
