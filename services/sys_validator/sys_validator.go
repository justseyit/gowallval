//NOT COMPLETED

package sys_validator

import (
	"regexp"
	//"gowallval/services/bitcoin_validator"
)

const sysRegEx string = "^sys1[qpzry9x8gf2tvdw0s3jn54khce6mua7l]{39}$"

func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(sysRegEx, address)
	return match //|| bitcoin_validator.IsValidAddress(address)
}
