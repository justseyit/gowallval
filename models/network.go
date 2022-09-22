package models

import "fmt"

type Network struct {
	ID     int    `json:"-"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func (n Network) String() string {
	return fmt.Sprintf("ID: %d\n Name: %s\n Symbol: %s", n.ID, n.Name, n.Symbol)
}
