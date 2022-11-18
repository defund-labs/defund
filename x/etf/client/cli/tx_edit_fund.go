package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/defund-labs/defund/x/etf/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEditFund() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-fund [symbol] [holdings]",
		Short: "Broadcast message EditFund",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSymbol := args[0]
			argHoldings := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditFund(
				clientCtx.GetFromAddress().String(),
				argSymbol,
				argHoldings,
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
