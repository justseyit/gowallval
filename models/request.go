package models

type Request struct {
	Address       string `json:"address"`
	NetworkName   string `json:"network_name"`
	NetworkSymbol string `json:"network_symbol"`
	Valid         bool   `json:"valid"`
}
