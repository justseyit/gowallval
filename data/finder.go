package data

import (
	. "github.com/seyitahmetgkc/gowallval/customerrors"
	. "github.com/seyitahmetgkc/gowallval/models"
)


func getCurrency(currSymbol string) (Currency, error) {

	//if unsupported currencies list contains the currency symbol, return error
	for _, currency := range UnsupportedCurrencies {
		if currency == currSymbol {
			return Currency{
				Symbol: currSymbol,
				Name:   "Unknown",
			}, 
			ErrorCurrencyNotSupported
		}
	}

	for _, curr := range AppCurrencies {
		if curr.Symbol == currSymbol {
			return curr, nil
		}
	}

	return Currency{
		Symbol: currSymbol,
		Name:   "Unknown",
	}, ErrorInvalidCurrencySymbol
}

func getNetwork(netwSymbol string) (Network, error) {
	
	//if unsupported networks list contains the network symbol, return error
	for _, network := range UnsupportedNetworks {
		if network == netwSymbol {
			return Network{
				Symbol: netwSymbol,
				Name:   "Unknown",
			}, 
			ErrorNetworkNotSupported
		}
	}

	for _, netw := range AppNetworks {
		if netw.Symbol == netwSymbol {
			return netw, nil
		}
	}

	return Network{
		Symbol: netwSymbol,
		Name:   "Unknown",
	}, ErrorInvalidNetworkSymbol
}

func GetResponse(req Request) (Response, error) {
	var response Response
	curr, err := getCurrency(req.CurrencySymbol)
	if err != nil {
		return response, err
	}
	netw, err := getNetwork(req.NetworkSymbol)
	if err != nil {
		return response, err
	}
	response.Address = req.Address
	response.Currency = curr
	response.Network = netw
	response.IsValidAddress = Validators[req.NetworkSymbol](req.Address)
	return response, nil
}