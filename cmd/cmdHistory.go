package cmd

import (
	"fmt"

	"github.com/das-pilot/das-go/api"
	"github.com/spf13/cobra"
)

var fromHistoryWallet string
var toHistoryWallet string
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
	historyCmd.Flags().StringVarP(&fromHistoryWallet, "fromWallet", "w", "", "Wallet name")
	historyCmd.Flags().StringVarP(&toHistoryWallet, "toWallet", "t", "", "Wallet name")
	historyCmd.Flags().Int64VarP(&startTime, "from", "s", -1, "Start time")
	historyCmd.Flags().Int64VarP(&endTime, "to", "e", -1, "End time")
}

func printHistory() {
	filter := api.HistoryFilter{
		FromWallet: fromHistoryWallet,
		ToWallet:   toHistoryWallet,
		TimeStart:  startTime,
		TimeEnd:    endTime}

	history := api.History(filter)
	for i, h := range history {
		fmt.Printf("%d) %+v\n", i+1, h)
	}

}
