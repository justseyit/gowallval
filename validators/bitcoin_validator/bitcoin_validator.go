package bitcoin_validator

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func IsValidAddress(address string) bool {
	addr, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	addr.EncodeAddress()
	return err == nil
}
