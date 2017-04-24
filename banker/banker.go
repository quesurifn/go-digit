package banker

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/plaid/plaid-go/plaid"
)

// Plaid wraps any interactions with Plaid
type Plaid struct {
	Client *plaid.Client
}

// PlaidConfig wraps the configs for Plaid
type PlaidConfig struct {
	ClientID  string
	Secret    string
	PublicKey string
}

// NewPlaid instantiates new Plaid struct
// TODO: Implement configs and read configs for whether or not this is development mode or not
func NewPlaid(cfg *PlaidConfig) *Plaid {
	return &Plaid{
		Client: plaid.NewClient(cfg.ClientID, cfg.Secret, plaid.Tartan),
	}
}

// ExchangeToken gets exchange tokens
func (p *Plaid) ExchangeToken(publicToken string) (string, error) {
	res, err := p.Client.ExchangeToken(publicToken)
	if err != nil {
		return "", err
	}
	spew.Dump(res)
	return res.AccessToken, nil
}
