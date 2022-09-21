package ltc_validator

import (
	"github.com/ltcsuite/ltcd/ltcutil"
	"github.com/ltcsuite/ltcd/chaincfg"
)

func IsValidAddress(address string) bool {
	addr, err := ltcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	addr.EncodeAddress()
	return err == nil
}