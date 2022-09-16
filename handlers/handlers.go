package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/seyitahmetgkc/gowallval/models"
)

func AddressValidationHandler(w http.ResponseWriter, r *http.Request) {

	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	response, _ := GetResponse(request)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, _ := json.Marshal(response)

	w.Write([]byte(data))
}
