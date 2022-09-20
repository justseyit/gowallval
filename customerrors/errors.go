package customerrors

import (
	"errors"
)

var ErrorInvalidNetworkSymbol error = errors.New("error: invalid network symbol")
var ErrorInvalidWalletAddress error = errors.New("error: invalid wallet address")
var ErrorInvalidCurrencySymbol error = errors.New("error: invalid currency symbol")
var ErrorCurrencyNotSupported error = errors.New("error: the currency not supported currently")
var ErrorNetworkNotSupported error = errors.New("error: the network not supported currently")
var ErrorCurrencyNotInThisNetwork error = errors.New("error: the currency does not exist on this network")

