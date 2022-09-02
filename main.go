package main

import (
	"fmt"
	"gowallval/utils/base58"

	xtz_validator "gowallval/services/xtz_validator"
)

func main() {
	base58.Init()
	m := xtz_validator.IsValidAddress("tz1b5xhEHFGkaN3fc6Aw8KsGxxxxQ5XUftps")
	fmt.Println(m)
}
