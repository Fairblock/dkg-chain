package cli

import (
	"encoding/json"
	"strconv"

	"dkg/x/dkg/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/spf13/cobra"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
)

var _ = strconv.Itoa(0)

func CmdRefundMsgRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "refund-msg-request [creator] [sender] [inner-message]",
		Short: "Broadcast message RefundMsgRequest",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddr := new(github_com_cosmos_cosmos_sdk_types.AccAddress)
			argMsg := new(types1.Any)

			err = json.Unmarshal([]byte(args[1]), argAddr)
			err = json.Unmarshal([]byte(args[2]), argMsg)
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRefundMsgRequest(
				clientCtx.GetFromAddress().String(),
				
				*argAddr,
				argMsg,
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
