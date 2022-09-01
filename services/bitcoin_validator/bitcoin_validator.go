package bitcoin_validator

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func IsValidAddress(addr string) bool {
	_, err := btcutil.DecodeAddress(addr, &chaincfg.MainNetParams)
	if err != nil {
		return false
	}
	return true
}
