package iost_validator

import "regexp"

const iostRegEx string = "^[a-z0-9_]{5,11}$"

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(iostRegEx, address)
	return match
}
