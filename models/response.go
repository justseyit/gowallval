package models

import "fmt"

type Response struct {
	Address        string   `json:"address"`
	Currency       Currency `json:"currency"`
	Network        Network  `json:"network"`
	IsValidAddress bool     `json:"isValidAddress"`
	Message        string   `json:"message"`
}

func (r Response) String() string {
	return fmt.Sprintf("Address: %s\n Currency: %s\n Network: %s\n IsValidAddress: %t\n Message: %s", r.Address, r.Currency, r.Network, r.IsValidAddress, r.Message)
}