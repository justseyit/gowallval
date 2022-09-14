package aud_validator

import "regexp"

const audRegex = "/^[0-9]{6,10}$/g" // Must be numbers only for the first 1 - 20 charactors with a capital L at the end

func IsValidAddress(address string) bool {
	search, _ := regexp.Compile(audRegex)
	match := search.MatchString(address)
	if match {
		return true
	} else {
		return false
	}
}
