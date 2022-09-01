package ethereum_validator

import (
	"fmt"
	"regexp"

	goeth "github.com/ethereum/go-ethereum/common"
)

const ethRegEx string = "^0x[0-9a-fA-F]{40}$"
const ethLowerRegEx string = "^0x[0-9a-f]{40}$"
const ethUpperRegEx string = "^0x?[0-9A-F]{40}$"

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(ethRegEx, address)
	// Check if it has the basic requirements of an address
	if !match {
		fmt.Println("Has no basic requirements of an address")
		return false
	}

	smallMatch, _ := regexp.MatchString(ethLowerRegEx, address)
	upperMatch, _ := regexp.MatchString(ethUpperRegEx, address)
	// If it's all small caps or all upper caps, return true
	if smallMatch || upperMatch {
		return true
	}

	// Otherwise check each case
	return goeth.IsHexAddress(address)
	//return verifyChecksum(address)
}

/*
func verifyChecksum(address string) bool {
	// Check each case
	address = strings.Replace(address, "0x", "", 1)

	//addressHash := keccak256.New().Hash([]byte(strings.ToLower(address)))
	addressHash := keccak256([]byte(strings.ToLower(address)))
	fmt.Println("addressHash: ", addressHash)

	for i := 0; i < 40; i++ {
		// The nth letter should be uppercase if the nth digit of casemap is 1
		if (parseInt(string(addressHash[i]), 16) > 7 && toUpperCase(address[i]) != address[i]) || (parseInt(string(addressHash[i]), 16) <= 7 && toLowerCase(address[i]) != address[i]) {
			return false
		}
	}
	return true
}

// The parseInt() function parses a string argument and returns an integer of the specified radix (the base in mathematical numeral systems).
func parseInt(str string, radix int) int {
	result, _ := strconv.ParseInt(str, radix, 0)
	return int(result)
}

// The toUpperCase() function returns the uppercase version of a byte.
func toUpperCase(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 32
	}
	return b
}

// The toLowerCase() function returns the lowercase version of a byte.
func toLowerCase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}

// The keccak256() function returns the Keccak-256 hash of the input data.
func keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}
*/
