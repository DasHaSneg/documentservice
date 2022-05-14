package cli

import (
	"strconv"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdContractsByInn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contracts-by-inn [inn]",
		Short: "Query contractsByInn",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argsInn := args[0]
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryContractsByInnRequest{
				Inn:        argsInn,
				Pagination: pageReq,
			}

			res, err := queryClient.ContractsByInn(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
