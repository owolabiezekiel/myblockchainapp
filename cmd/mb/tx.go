package main

import (
	"com.owoez/myblockchainapp/database"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const flagFrom = "from"
const flagTo = "to"
const flagValue = "value"
const flagData = "data"

func txCmd() *cobra.Command {
	var txsCmd = &cobra.Command{
		Use:   "tx",
		Short: "Interact with txs (add...).",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageError()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	txsCmd.AddCommand(txAddCmd())
	return txsCmd
}

func txAddCmd() *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Adds new TX to database.",
		Run: func(cmd *cobra.Command, args []string) {
			from, _ := cmd.Flags().GetString(flagFrom)
			to, _ := cmd.Flags().GetString(flagTo)
			value, _ := cmd.Flags().GetUint(flagValue)
			data, _ := cmd.Flags().GetString(flagData)

			tx := database.NewTx(database.NewAccount(from), database.NewAccount(to), value, data)

			state, err := database.NewStateFromDisk()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer state.Close()

			err = state.Add(tx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			_, err = state.Persist()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Println("TX successfully persisted to the ledger.")
		},
	}

	addCmd.Flags().String(flagFrom, "", "From what account to send tokens")
	addCmd.MarkFlagRequired(flagFrom)

	addCmd.Flags().String(flagTo, "", "To what account to send tokens")
	addCmd.MarkFlagRequired(flagTo)

	addCmd.Flags().Uint(flagValue, 0, "How many tokens to send")
	addCmd.MarkFlagRequired(flagValue)

	addCmd.Flags().String(flagData, "", "Possible values: 'reward'")

	return addCmd
}
