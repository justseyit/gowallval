//NOT COMPLETED

package dash_validator

import (
	"github.com/dashpay/godash/chaincfg"
	"github.com/dashpay/godashutil"
)

func IsValidAddress(address string) bool {
	_, err := godashutil.DecodeAddress(address, &chaincfg.MainNetParams)
	return err == nil
}
