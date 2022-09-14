package bnb_validator

import "regexp"

const bnbRegex string = "^(bnb)([a-z0-9]{39})$" // bnb + 39 a-z0-9

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(bnbRegex, address)
	return match
}
