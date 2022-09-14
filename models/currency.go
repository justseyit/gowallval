package models

import (
	"gowallval/customerrors"
	//"gowallval/models/data"
)

type Currency struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Network   Network `json:"network"`
}

func GetCurrency(currSymbol string, netwSymbol string) (Currency, error) {

	var currency Currency
	for _, network := range UnsupportedNetworks {
		if network == netwSymbol {
			return Currency{}, customerrors.ErrorNetworkNotSupported
		}
	}

	for _, currency := range UnsupportedCurrencies {
		if currency == currSymbol {
			return Currency{}, customerrors.ErrorCurrencyNotSupported
		}
	}

	for _, netw := range Networks {
		if netw.Symbol == netwSymbol {
			currency.Network = netw
		}
	}

	for _, curr := range Currencies {
		if curr.Symbol == currSymbol {
			currency.Symbol = curr.Symbol
			currency.Name = curr.Name
		}
	}

	if currency.Symbol == "" && currency.Name == "" {
		return Currency{
			Symbol: currSymbol,
			Name:   "Unknown",
			Network: Network{
				Name:      currency.Network.Name,
				Symbol:    currency.Network.Symbol,
			},
		}, customerrors.ErrorInvalidCurrencySymbol
	}else if currency.Network.Symbol == "" && currency.Network.Name == "" {
		return Currency{
			Symbol: currSymbol,
			Name:   currency.Name,
			Network: Network{
				Name:      "Unknown",
				Symbol:    netwSymbol,
			},
		}, customerrors.ErrorInvalidNetworkSymbol
	}


	return currency, nil
}
