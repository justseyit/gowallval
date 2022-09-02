//NOT COMPLETED
//Gives false positives

package xtz_validator

import (
	base58 "github.com/btcsuite/btcutil/base58"
)

func IsValidAddress(address string) bool {
	decoded := base58.Decode(address)

	return decoded != nil && len(decoded) == 27
}
