package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/defund-labs/defund/v1/x/query/types"
)

func CmdShowInterquery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "interquery [key] [id]",
		Short: "Gets an interquery for key-id pair",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			storeid := fmt.Sprintf("%s-%s", args[0], args[1])

			params := &types.QueryGetInterqueryRequest{
				Storeid: storeid,
			}

			res, err := queryClient.Interquery(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowInterqueryAll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "interqueries",
		Short: "List all interqueries",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllInterqueryRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.InterqueryAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowInterqueryResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "interquery-result [key] [id]",
		Short: "Gets an interquery result for key-id pair",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			storeid := fmt.Sprintf("%s-%s", args[0], args[1])

			params := &types.QueryGetInterqueryRequest{
				Storeid: storeid,
			}

			res, err := queryClient.Interquery(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowInterqueryTimeoutResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "interquery-timeout [key] [id]",
		Short: "Gets an interquery timeout result for key-id pair",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			storeid := fmt.Sprintf("%s-%s", args[0], args[1])

			params := &types.QueryGetInterqueryRequest{
				Storeid: storeid,
			}

			res, err := queryClient.Interquery(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
