package steem_validator

import (
	"regexp"
	"strings"
)

const accountRegex string = "^[a-z0-9-.]{3,}$"
const segmentRegex string = "^[a-z][a-z0-9-]+[a-z0-9]$"
const doubleDashRegex string = "--"

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(accountRegex, address)
	if !match {
		return false
	}
	segments := strings.Split(address, ".")

	for _, segment := range segments {
		if len(segment) < 3 {
			return false
		}

		match, _ := regexp.MatchString(segmentRegex, segment)
		if !match {
			return false
		}

		match, _ = regexp.MatchString(doubleDashRegex, segment)
		if match {
			return false
		}
	}
	return true
}
