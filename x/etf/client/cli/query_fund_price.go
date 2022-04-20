package cli

import (
	"context"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/defund-labs/defund/x/etf/types"
)

var _ = strconv.Itoa(0)

func CmdFundPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fund-price [symbol]",
		Short: "Get the price of a fund",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqSymbol := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryFundPriceRequest{

				Symbol: reqSymbol,
			}

			res, err := queryClient.FundPrice(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdFundPriceAll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fund-prices [symbol]",
		Short: "Get the historical prices of a fund by symbol",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argSymbol := args[0]

			params := &types.QueryAllFundPriceRequest{
				Symbol:     argSymbol,
				Pagination: pageReq,
			}

			res, err := queryClient.FundPriceAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd

}
