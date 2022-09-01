//NOT COMPLETED

package iota_validator

import (
	"github.com/btcsuite/btcutil/bech32"
	iotago "github.com/iotaledger/iota.go/v2"
)

func IsValidAddress(address string) bool {
	enc, _ := bech32.Encode("", []byte(address))
	_, addr, err := iotago.ParseBech32(enc)
	if addr.String() != address || err != nil {
		return false
	}
	return true
}
