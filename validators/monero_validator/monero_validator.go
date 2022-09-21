package monero_validator


import (
	"github.com/MixinNetwork/mixin/domains/monero"
)

func IsValidAddress(address string) bool {
	return monero.VerifyAddress(address) == nil
}