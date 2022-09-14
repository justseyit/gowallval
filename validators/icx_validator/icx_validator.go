package icx_validator

import "regexp"

const icxRegEx string = "^hx[0-9a-f]{40}$" //Begins with hx followed by 40 hex chars

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(icxRegEx, address)
	return match
}
