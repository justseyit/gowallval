package models

import "fmt"

type Request struct {
	Address        string `json:"address"`
	CurrencySymbol string `json:"currencySymbol"`
	NetworkSymbol  string `json:"networkSymbol"`
}

func (r Request) String() string {
	return fmt.Sprintf("Address: %s\n CurrencySymbol: %s\n NetworkSymbol: %s", r.Address, r.CurrencySymbol, r.NetworkSymbol)
}
