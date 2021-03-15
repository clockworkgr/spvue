package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/clockworkgr/spvue/x/blog/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

func CmdCreatePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-post [title] [body] [votes]",
		Short: "Creates a new post",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsTitle := string(args[0])
			argsBody := string(args[1])
			argsVotes, _ := strconv.ParseInt(args[2], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePost(clientCtx.GetFromAddress().String(), string(argsTitle), string(argsBody), int32(argsVotes))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdatePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-post [id] [title] [body] [votes]",
		Short: "Update a post",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsTitle := string(args[1])
			argsBody := string(args[2])
			argsVotes, _ := strconv.ParseInt(args[3], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePost(clientCtx.GetFromAddress().String(), id, string(argsTitle), string(argsBody), int32(argsVotes))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeletePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-post [id] [title] [body] [votes]",
		Short: "Delete a post by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePost(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
