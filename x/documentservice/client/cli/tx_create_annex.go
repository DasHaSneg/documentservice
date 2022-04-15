package cli

import (
	"strconv"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateAnnex() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-annex [annex-hash] [contract-id] [buyer]",
		Short: "Broadcast message create-annex",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAnnexHash := args[0]
			argContractId := args[1]
			argBuyer := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			contractId, err := strconv.ParseInt(argContractId, 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAnnex(
				clientCtx.GetFromAddress().String(),
				argAnnexHash,
				uint64(contractId),
				argBuyer,
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
