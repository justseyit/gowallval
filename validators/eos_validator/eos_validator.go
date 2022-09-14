package eos_validator

import "regexp"

const eosRegEx string = "^[a-z0-9]+$" // Must be numbers and lowercase letters only

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(eosRegEx, address)
	if len(address) == 12 && match {
		return true
	} else {
		return false
	}
}
