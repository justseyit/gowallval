package models

type Request struct {
	Address        string `json:"address"`
	CurrencySymbol string `json:"currencySymbol"`
	NetworkSymbol  string `json:"networkSymbol"`
}
