package zil_validator

import (
	"regexp"

	bech32 "github.com/btcsuite/btcutil/bech32"
)

const allowedChars string = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

const regExp string = "^(zil)1([" + allowedChars + "]+)$" // zil + bech32 separated by '1'

func IsValidAddress(address string) bool {
	match, err := regexp.MatchString(regExp, address)
	if err != nil {
		return false
	}
	return verifyChecksum(address) && match
}

func verifyChecksum(address string) bool {
	_, decodedbyte, _ := bech32.Decode(address)

	return len(decodedbyte) == 32
}
