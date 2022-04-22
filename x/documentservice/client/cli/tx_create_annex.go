package cli

import (
	"strconv"

	"github.com/cosmonaut/documentservice/x/documentservice/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
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
			argBuyer := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argContractId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAnnex(
				clientCtx.GetFromAddress().String(),
				argAnnexHash,
				argContractId,
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