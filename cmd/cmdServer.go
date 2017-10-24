package cmd

import (
	"fmt"

	"github.com/das-pilot/das-go/srv"
	"github.com/spf13/cobra"
	"github.com/das-pilot/das-go/api"
)

var port int

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "DAS standalone server",
	Long:  "Start the DAS HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("DAS Server started on port %d\n", port)
		api.Init()
		srv.Start(port)
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVarP(&port, "port", "p", 8000, "Server port")
}
