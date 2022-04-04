package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/defund-labs/defund/v1/x/etf/types"
)

var _ = strconv.Itoa(0)

func CmdFundPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fund-price [ticker]",
		Short: "Get the price of a fund",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqTicker := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryFundPriceRequest{

				Ticker: reqTicker,
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
