package models

type Response struct {
	Address        string   `json:"address"`
	Currency       Currency `json:"currency"`
	IsValidAddress bool     `json:"isValidAddress"`
	Message        string   `json:"message"`
}
