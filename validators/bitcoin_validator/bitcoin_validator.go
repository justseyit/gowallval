package bitcoin_validator

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func IsValidAddress(address string) bool {
	_, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	return err == nil
}
