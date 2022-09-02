package nano_validator

import (
	"encoding/hex"
	"regexp"
	"strings"

	basex "github.com/hailongz/golang/basex"
	blake2b "golang.org/x/crypto/blake2b"
)

const allowedChars string = "13456789abcdefghijkmnopqrstuwxyz"

const regEx string = "^(xrb|nano)_([" + allowedChars + "]{60})$"

func IsValidAddress(address string) bool {
	return verifyChecksum(address)
}

func verifyChecksum(address string) bool {
	match, err := regexp.MatchString(regEx, address)
	if err != nil || !match {
		return false
	}
	codec, _ := basex.NewEncoding(allowedChars)
	addrstr := strings.Split(address, "_")[1]
	addrstr = addrstr[len(addrstr)-37:]
	bytes, _ := codec.Decode(addrstr)

	hexx, _ := hex.DecodeString(string(bytes[len(bytes)-5:]))
	computedChecksum, _ := blake2b.New(5, hexx)

	checksum, _ := hex.DecodeString(reverse(string(bytes[len(bytes)-5:])))

	same := true
	for j := range checksum {
		if checksum[j] != computedChecksum.Sum(nil)[j] {
			same = false
		}
	}

	return same
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
