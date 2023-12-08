package data

import (
	customerrors "github.com/seyitahmetgkc/gowallval/customerrors"
	models "github.com/seyitahmetgkc/gowallval/models"
)

func getCurrency(currSymbol string) (models.Currency, error) {

	//if unsupported currencies list contains the currency symbol, return error
	for _, currency := range UnsupportedCurrencies {
		if currency == currSymbol {
			return models.Currency{
					Symbol: currSymbol,
					Name:   "Unknown",
				},
				customerrors.ErrorCurrencyNotSupported
		}
	}

	for _, curr := range AppCurrencies {
		if curr.Symbol == currSymbol {
			return curr, nil
		}
	}

	return models.Currency{
		Symbol: currSymbol,
		Name:   "Unknown",
	}, customerrors.ErrorInvalidCurrencySymbol
}

func getNetwork(netwSymbol string) (models.Network, error) {

	//if unsupported networks list contains the network symbol, return error
	for _, network := range UnsupportedNetworks {
		if network == netwSymbol {
			return models.Network{
					Symbol: netwSymbol,
					Name:   "Unknown",
				},
				customerrors.ErrorNetworkNotSupported
		}
	}

	for _, netw := range AppNetworks {
		if netw.Symbol == netwSymbol {
			return netw, nil
		}
	}

	return models.Network{
		Symbol: netwSymbol,
		Name:   "Unknown",
	}, customerrors.ErrorInvalidNetworkSymbol
}

func checkCurrencyExistsInNetwork(currID int, netwID int) error {
	for _, mapping := range AppCurrencyNetworkMappings {
		if mapping.CurrencyID == currID && mapping.NetworkID == netwID {
			return nil
		}
	}
	return customerrors.ErrorCurrencyNotInThisNetwork
}

func GetResponse(req models.Request) (models.Response, error) {
	var response models.Response
	curr, err := getCurrency(req.CurrencySymbol)
	if err != nil {
		return response, err
	}
	netw, err := getNetwork(req.NetworkSymbol)
	if err != nil {
		return response, err
	}
	err = checkCurrencyExistsInNetwork(curr.ID, netw.ID)
	if err != nil {
		response.Address = req.Address
		response.Currency = curr
		response.Network = netw
		return response, err
	}
	response.Address = req.Address
	response.Currency = curr
	response.Network = netw
	response.IsValidAddress = Validators[req.NetworkSymbol](req.Address)
	return response, nil
}
