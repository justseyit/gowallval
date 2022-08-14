package main

import (
	"fmt"
	"gowallval/services/steem_validator"
)

func main() {
	m := steem_validator.IsValidAddress("ebookwriter")
	fmt.Println(m)
}
