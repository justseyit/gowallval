package main

import (
	"fmt"
	"gowallval/utils/base58"

	iota_validator "gowallval/services/iota_validator"
)

func main() {
	base58.Init()
	m := iota_validator.IsValidAddress("OGMMQJUDMNNYSOAXMJWAMNAJPHWMGVAY9UWBXRGTXXVEDIEWSNYRNDQY99NDJQB9QQBPCRRNFAIUPGPLZ")
	fmt.Println(m)
}
