package hbar_validator

import (
	"regexp"
	"strings"
)

const hbarRegEx string = "^\\d+$"

func IsValidAddress(address string) bool {
	split := strings.Split(address, ".")
	if split[0] != "0" || split[1] != "0" {
		return false
	}

	match, _ := regexp.MatchString(hbarRegEx, split[2])

	if len(split[2]) <= 6 && match {
		return true
	}

	return false
}
