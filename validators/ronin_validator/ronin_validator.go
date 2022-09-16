package ronin_validator

import (
	ethereum_validator "github.com/seyitahmetgkc/gowallval/validators/ethereum_validator"
)

func IsValidAddress(address string) bool {
	addr := roninTo0x(address)
	return ethereum_validator.IsValidAddress(addr)
}

func roninTo0x(address string) string {
	addr := address[6:]
	return "0x" + addr
}