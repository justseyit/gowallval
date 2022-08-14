package atom_validator

import (
	"regexp"

	"github.com/btcsuite/btcutil/bech32"
)

const allowed_chars = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"
const atomRegex = "^(cosmos)1([" + allowed_chars + "]+)$" // cosmos + bech32 separated by "1"

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(atomRegex, address)
	if match {
		return verifyChecksum(address)
	} else {
		return false
	}
}

func verifyChecksum(address string) bool {
	_, decoded, _ := bech32.Decode(address)

	if decoded != nil {
		return len(decoded) == 32
	} else {
		return false
	}
}
