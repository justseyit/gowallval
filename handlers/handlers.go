package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	data "github.com/seyitahmetgkc/gowallval/data"
	models "github.com/seyitahmetgkc/gowallval/models"
)

func AddressValidationHandler(w http.ResponseWriter, r *http.Request) {

	var request models.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	logRequest(request)

	response, err := data.GetResponse(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Message = err.Error()
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, _ := json.Marshal(response)

	w.Write([]byte(data))

	logResponse(response)
}

func logRequest(request models.Request) {
	log.Println("Request: ", request.String())
	val, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(val))
	fmt.Println("----------------------------------------")
}

func logResponse(response models.Response) {
	log.Println("Response: ", response.String())
	val, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(val))
	fmt.Println("----------------------------------------")
}
