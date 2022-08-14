package sc_validator

import (
	"regexp"
)

const scRegEx string = "^[0-9a-f]{76}$" // 76 hex chars

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(scRegEx, address)
	return match
}
