package cli

import (
	"strconv"

	"dkg/x/dkg/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdKeygenResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keygen-result [mpk] [commitment]",
		Short: "Broadcast message keygen-result",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMpk := args[0]
			argCommitment := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgKeygenResult(
				clientCtx.GetFromAddress().String(),
				argMpk,
				argCommitment,
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
