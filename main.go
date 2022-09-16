package main

import (
	"net/http"

	. "github.com/seyitahmetgkc/gowallval/handlers"
	"github.com/seyitahmetgkc/gowallval/models"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter().StrictSlash(true)
	models.InitializeNetworks()
	mux.HandleFunc("/", AddressValidationHandler).Methods("POST")

	http.ListenAndServe(":9000", mux)

}
