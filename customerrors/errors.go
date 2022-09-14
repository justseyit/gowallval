package customerrors

import (
	"errors"
)

var ErrorInvalidNetworkSymbol error = errors.New("error: invalid network symbol")
var ErrorInvalidWalletAddress error = errors.New("error: invalid wallet address")
var ErrorInvalidCurrencySymbol error = errors.New("error: invalid currency symbol")
var ErrorCurrencyNotSupported error = errors.New("error: currency not supported yet")
var ErrorNetworkNotSupported error = errors.New("error: network not supported yet")

