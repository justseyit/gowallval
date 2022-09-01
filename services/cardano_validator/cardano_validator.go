package cardano_validator

import (
	gocardano "github.com/fivebinaries/go-cardano-serialization/address"
)

func IsValidAddress(address string) bool {
	_, err := gocardano.NewAddress(address)
	if err != nil {
		return false
	}
	return true
}
