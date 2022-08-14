package customerrors

import (
	"errors"
)

var InvalidNetworkSymbol error = errors.New("error: invalid network symbol")
var InvalidWalletAddress error = errors.New("error: invalid wallet address")

const Succeeded = "succeeded"
