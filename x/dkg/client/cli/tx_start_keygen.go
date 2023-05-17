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

func CmdStartKeygen() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-keygen [key-id] [threshold] [timeout] [participants]",
		Short: "Broadcast message startKeygen",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argKeyID := args[0]
			argThreshold := args[1]
			argTimeout := args[2]
			argParticipants := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgStartKeygen(
				clientCtx.GetFromAddress().String(),
				argKeyID,
				argThreshold,
				argTimeout,
				argParticipants,
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
