package models


type Network struct {
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
	Validator func(string) bool
}