package tron_validator

import (
	"github.com/MixinNetwork/mixin/domains/tron"
)

func IsValidAddress(address string) bool {
	return tron.VerifyAddress(address) == nil
}
