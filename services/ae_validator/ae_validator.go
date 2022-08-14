package ae_validator

import (
	"log"
	"regexp"
	"strings"

	base58 "github.com/btcsuite/btcutil/base58"
)

const ALLOWED_CHARS = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

const RegExp = "^(ak_)([" + ALLOWED_CHARS + "]+)$"

func verifyChecksum(address string) bool {
	addr := strings.TrimPrefix(address, "ak_")
	decoded := base58.Decode(addr)
	//remove last 4 bytes
	decoded = decoded[:len(decoded)-4]

	return len(decoded) == 32
}

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(RegExp, address)
	if !match {
		log.Println("IsValidAddress - Address is not valid")
		return false
	} else {
		return verifyChecksum(address)
	}
}
