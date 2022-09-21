package sol_validator

import (
	"github.com/MixinNetwork/mixin/domains/solana"
)

func IsValidAddress(address string) bool {
	return solana.VerifyAddress(address) == nil
}