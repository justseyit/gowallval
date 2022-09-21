package zcash_validator

import (
	"github.com/MixinNetwork/mixin/domains/zcash"
)

func IsValidAddress(address string) bool {
	err := zcash.VerifyAddress(address)
	return err == nil
}