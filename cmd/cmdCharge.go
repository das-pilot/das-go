package cmd

import (
	"fmt"

	"github.com/das-pilot/das-go/api"
	"github.com/spf13/cobra"
)

var chargeFromWallet string
var chargeToWallet string
var chargeAmount float32

// chargeCmd represents the charge command
var chargeCmd = &cobra.Command{
	Use:   "charge",
	Short: "Charge another wallet",
	Long:  "Charge another wallet providing its name and amount to charge",
	Run: func(cmd *cobra.Command, args []string) {
		printChargeResult()
	},
}

func init() {
	RootCmd.AddCommand(chargeCmd)
	chargeCmd.Flags().StringVarP(&chargeFromWallet, "fromWallet", "w", "", "Wallet name")
	chargeCmd.Flags().StringVarP(&chargeToWallet, "toWallet", "t", "", "Wallet name")
	chargeCmd.Flags().Float32VarP(&chargeAmount, "amount", "a", 0, "Amount to charge")
}

func printChargeResult() {
	fmt.Printf("%+v\n", api.Charge(chargeFromWallet, chargeToWallet, chargeAmount))
}
