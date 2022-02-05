package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/defundhub/defund/x/query/types"
)

func CmdCreateInterquery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-interquery [index] [height] [path] [chain-id] [type-name]",
		Short: "Create a new interquery",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argHeight := args[1]
			argPath := args[2]
			argChainId := args[3]
			argTypeName := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateInterquery(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argHeight,
				argPath,
				argChainId,
				argTypeName,
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

func CmdUpdateInterquery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-interquery [index] [height] [path] [chain-id] [type-name]",
		Short: "Update a interquery",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argHeight := args[1]
			argPath := args[2]
			argChainId := args[3]
			argTypeName := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateInterquery(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argHeight,
				argPath,
				argChainId,
				argTypeName,
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

func CmdDeleteInterquery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-interquery [index]",
		Short: "Delete a interquery",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteInterquery(
				clientCtx.GetFromAddress().String(),
				indexIndex,
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
