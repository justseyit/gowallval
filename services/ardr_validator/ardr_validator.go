package ardr_validator

import "regexp"

const ardorRegex = "^ARDOR(-[A-Z0-9]{4}){3}(-[A-Z0-9]{5})$"

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(ardorRegex, address)

	return match
}
