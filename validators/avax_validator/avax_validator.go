package avax_validator

import (
	"github.com/MixinNetwork/mixin/domains/avalanche"
)

func IsValidAddress(address string) bool {
	return avalanche.VerifyAddress(address) == nil
}