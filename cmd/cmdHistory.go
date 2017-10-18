package cmd

import (
	"fmt"

	"github.com/das-pilot/das-go/api"
	"github.com/spf13/cobra"
)

var historyWallet string
var startTime int64
var endTime int64

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show transactions history",
	Long:  "Show history of transactions with specified wallets for specified period of time",
	Run: func(cmd *cobra.Command, args []string) {
		printHistory()
	},
}

func init() {
	RootCmd.AddCommand(historyCmd)
	historyCmd.Flags().StringVarP(&historyWallet, "wallet", "w", "", "Wallet name")
	historyCmd.Flags().Int64VarP(&startTime, "from", "f", -1, "Start time")
	historyCmd.Flags().Int64VarP(&endTime, "to", "t", -1, "End time")
}

func printHistory() {
	filter := api.HistoryFilter{
		Wallet:    historyWallet,
		TimeStart: startTime,
		TimeEnd:   endTime}

	history := api.History(filter)
	for i, h := range history {
		fmt.Printf("%d) %+v\n", i+1, h)
	}

}
