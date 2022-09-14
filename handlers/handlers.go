package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/seyitahmetgkc/gowallval/models"

	"github.com/gorilla/mux"
)

func AddressValidationHandler(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
	currencySymbol := vars["currencySymbol"]
	networkSymbol := vars["networkSymbol"]
	address := vars["address"]

	var request models.Request
	request.Address = address
	request.CurrencySymbol = currencySymbol
	request.NetworkSymbol = networkSymbol

	var response models.Response

	curr, err := models.GetCurrency(currencySymbol, networkSymbol)
	response.Address = address
	response.Currency = curr
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		response.Message = err.Error()
		response.IsValidAddress = false
	} else {
		response.IsValidAddress = response.Currency.Network.Validator(address)
		response.Message = "Address validation completed."
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, _ := json.Marshal(response)

	w.Write([]byte(data))
	//json.NewEncoder(w).Encode(response)
}
