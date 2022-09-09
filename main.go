package main

import (
	"fmt"
	"gowallval/utils/base58"

	ripple_validator "gowallval/services/ripple_validator"
)

func main() {
	base58.Init()
	m := ripple_validator.IsValidAddress("rpSpKdgBL4ecPx1b6bE43dRybbSv1TRADe")
	fmt.Println(m)
}
