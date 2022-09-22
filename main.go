package main

import (
	"net/http"

	"github.com/seyitahmetgkc/gowallval/data"
	. "github.com/seyitahmetgkc/gowallval/handlers"

	"github.com/gorilla/mux"
	"github.com/seyitahmetgkc/gowallval/validators/band_validator"
)

func main() {
	mux := mux.NewRouter().StrictSlash(true)
	models.InitializeNetworks()
	data.LoadFromDB()
	mux.HandleFunc("/", AddressValidationHandler).Methods("POST")

	http.ListenAndServe(":9000", mux)
}
