package cmd

import (
	"fmt"

	"github.com/quesurifn/go-digit/digit"
	"github.com/spf13/cobra"
)

var port string

var digitRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the go-digit gin server",
	Long:  `Runs the go-digit gin server. It will default to running on 3030 if a port flag (--port or -p) is not specified`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world!")
		s := digit.NewServer(port)
		err := s.Run()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	digitRunCmd.PersistentFlags().StringVarP(&port, "port", "p", "3030", "The port that the go-digit gin server will run on.")
}
