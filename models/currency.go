package models

type Currency struct {
	Symbol         string
	Name           string
	ExpectedLength int
	HashFunction   string
	RegEx          string
	AddressTypes   map[string]interface{}
}
