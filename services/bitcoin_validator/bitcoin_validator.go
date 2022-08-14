//NOT COMPLETED

package bitcoin_validator

import (
	"encoding/hex"
	"gowallval/utils/base58"

	blake256 "github.com/dchest/blake256"
	keccak256 "github.com/wealdtech/go-merkletree/keccak256"
	"golang.org/x/crypto/blake2b"
)

const default_network_type = "prod"

func getDecoded(address string) []int {
	arr, err := base58.Decode(address)
	if err != nil {
		return nil
	} else {
		return arr
	}
}

func getChecksum(hashFunction, payload string) string {

	switch hashFunction {
	case "blake256keccak256":
		b256, _ := blake2b.New(32, nil)
		b256.Write([]byte(payload))
		return string(keccak256.New().Hash([]byte(byteArrayToHexString(b256.Sum(nil)))))

	case "blake256":
		b256 := blake256.Sum256([]byte(payload))
		return byteArrayToString(b256)

	case "keccak256":
		return string(keccak256.New().Hash([]byte(payload)))

	case "sha256":
		return byteArrayToString(blake256.Sum256([]byte(payload)))

	default:
		return byteArrayToString(blake256.Sum256([]byte(payload)))
	}
}

func getAddressType(address, currency string) string {
	decoded := getDecoded(address)
	if decoded == nil {
		return ""
	}
	if decoded[0] == 0 {
		return "mainnet"
	} else if decoded[0] == 111 {
		return "testnet"
	} else {
		return ""
	}
}

func byteArrayToHexString(byteArray []byte) string {
	return hex.EncodeToString(byteArray)
}

func byteArrayToString(byteArray [32]byte) string {
	return string(byteArray[:])
}
