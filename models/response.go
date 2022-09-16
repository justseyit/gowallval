package models

type Response struct {
	Address        string   `json:"address"`
	Currency       Currency `json:"currency"`
	IsValidAddress bool     `json:"isValidAddress"`
	Message        string   `json:"message"`
}

func GetResponse(req Request) (Response, error) {
	var response Response
	curr, err := getCurrency(req.CurrencySymbol, req.NetworkSymbol)
	response.Address = req.Address
	response.Currency = curr
	if err != nil {
		response.Message = err.Error()
		response.IsValidAddress = false
		return response, err
	} else {
		response.IsValidAddress = response.Currency.Network.Validator(req.Address)
		response.Message = ""
		return response, nil
	}
}
