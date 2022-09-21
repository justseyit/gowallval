package dash_validator

import (
	"github.com/MixinNetwork/mixin/domains/dash"
)

//const dashRegex string = "^X[1-9A-HJ-NP-Za-km-z]{33}"

func IsValidAddress(address string) bool {
	return dash.VerifyAddress(address) == nil
}
