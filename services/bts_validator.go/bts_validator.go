package bts_validator

import "regexp"

const btsRegex string = "/^[a-z0-9-.]+$/g" // Must be numbers and lowercase letters only

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(btsRegex, address)
	return match
}
