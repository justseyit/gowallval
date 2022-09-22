package models

import "fmt"

type Currency struct {
	ID     int    `json:"-"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func (c Currency) String() string {
	return fmt.Sprintf("ID: %d\n Name: %s\n Symbol: %s", c.ID, c.Name, c.Symbol)
}
