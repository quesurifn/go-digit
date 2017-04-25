package cmd

import (
	"github.com/quesurifn/go-digit/banker"
	"github.com/spf13/cobra"
)

const (
	defaultPublicToken string = "replaceme"
)

// Flags
var (
	publicToken string
)

var plaidCmd = &cobra.Command{
	Use:   "plaid",
	Short: "Wraps useful plaid commands",
	Long:  "Wraps useful plaid commands",
}

var plaidExchangeTokenCmd = &cobra.Command{
	Use:   "exchange",
	Short: "Exchanges a public token for an access one",
	Long:  "Exchanges a public token for an access one. Use the --token flag to specify the token",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := &banker.PlaidConfig{
		// Stuff here
		}
		p := banker.NewPlaid(cfg)
		_, err := p.ExchangeToken(publicToken)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	plaidExchangeTokenCmd.PersistentFlags().StringVarP(&publicToken, "token", "t", defaultPublicToken, "The public token that will be exchanged for an access token")
	plaidCmd.AddCommand(plaidExchangeTokenCmd)
}
