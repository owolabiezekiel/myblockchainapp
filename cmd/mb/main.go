package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var mbCmd = &cobra.Command{
		Use:   "mb",
		Short: "My Blockchain CLI",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	mbCmd.AddCommand(versionCmd)
	mbCmd.AddCommand(balancesCmd())
	mbCmd.AddCommand(txCmd())

	err := mbCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func incorrectUsageError() error {
	return fmt.Errorf("incorrect usage")
}
