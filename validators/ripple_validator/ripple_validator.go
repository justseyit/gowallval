package ripple_validator

import (
	//"crypto/sha256"

	"regexp"

	goripple "github.com/rubblelabs/ripple/data"
)

const allowed_chars string = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
const regEx string = "^r[" + allowed_chars + "]{24,34}$"

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(regEx, address)
	_, err := goripple.NewAccountFromAddress(address)
	return match && err == nil
}

/*
func verifyChecksum(address string) bool {
	codec, _ := basex.NewEncoding(allowed_chars)
	bytes, err := codec.Decode(address)
	if err != nil {
		fmt.Println(err)
		return false
	}

	computedChecksum := sha256.Sum256([]byte(toHex(bytes[:len(bytes)-4])))
	checksum := toHex(bytes[len(bytes)-4:])

	fmt.Println("computedChecksum: " + byteArrayToString(computedChecksum))
	fmt.Println("checksum: " + checksum)

	return checksum == byteArrayToString(computedChecksum)
}

func toHex(bytes []byte) string {
	hex := ""

	for _, b := range bytes {
		hex += fmt.Sprintf("%02x", b)
	}

	return hex
}

// converts byte array to string
func byteArrayToString(byteArray [32]byte) string {
	var str string
	for _, b := range byteArray {
		str += fmt.Sprintf("%d", b)
	}
	return str
}

func isEquals(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sha256(data []byte) []byte {
	hashh := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}
*/
