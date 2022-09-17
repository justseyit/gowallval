package models

type Network struct {
	ID     int    `json:"-"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
