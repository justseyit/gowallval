package nem_validator

import "regexp"

const allowed_chars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
const nemRegEx string = "^N[" + allowed_chars + "]{39}$"

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(nemRegEx, address)
	return match
}
