package lisk_validator

import "regexp"

const regex string = "^[0-9]{1,20}L$" // Must be numbers only for the first 1 - 20 charactors with a capital L at the end

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(regex, address)
	return match
}
