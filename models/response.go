package models

type Response struct {
	Address       string `json:"address"`
	NetworkSymbol string `json:"network_symbol"`
	NetworkName   string `json:"network_name"`
	Message       string `json:"message"`
	Valid         bool   `json:"valid"`
}
