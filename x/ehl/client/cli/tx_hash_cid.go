package cli

import (
	"strconv"

	"bitmail/x/ehl/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateHashCid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-hash-cid [receiver] [hashlink] [vaultid]",
		Short: "Create a new hash-cid",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argHashlink := args[1]
			argVaultid := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateHashCid(clientCtx.GetFromAddress().String(), argReceiver, argHashlink, argVaultid)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateHashCid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-hash-cid [id] [receiver] [hashlink] [vaultid]",
		Short: "Update a hash-cid",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argReceiver := args[1]

			argHashlink := args[2]

			argVaultid := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateHashCid(clientCtx.GetFromAddress().String(), id, argReceiver, argHashlink, argVaultid)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteHashCid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-hash-cid [id]",
		Short: "Delete a hash-cid by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteHashCid(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
