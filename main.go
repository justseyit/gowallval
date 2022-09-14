package main

import (
	"fmt"
	"gowallval/utils/base58"

	ethereum_validator "gowallval/validators/ethereum_validator"
)

func main() {
	base58.Init()
	m := ethereum_validator.IsValidAddress("0x4E656459ed25bF986Eea1196Bc1B00665401645d")
	fmt.Println(m)
}
