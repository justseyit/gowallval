package models

type Currency struct {
	ID     int    `json:"-"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
