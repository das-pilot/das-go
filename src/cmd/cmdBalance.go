package cmd

import (
	"fmt"

	"github.com/das-pilot/das-go/api"
	"github.com/spf13/cobra"
)

var balanceWallet string

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance -w <wallet name?>",
	Short: "Show your balance",
	Long:  "Show your current balance. Optionally - with particular wallet",
	Run: func(cmd *cobra.Command, args []string) {
		printBalance()
	},
}

func init() {
	RootCmd.AddCommand(balanceCmd)
	balanceCmd.Flags().StringVarP(&balanceWallet, "wallet", "w", "", "Wallet name")
}

func printBalance() {
	fmt.Printf("%+v\n", api.Balance(balanceWallet))
}
