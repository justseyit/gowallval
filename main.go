package main

import (
	"net/http"

	"github.com/seyitahmetgkc/gowallval/data"
	. "github.com/seyitahmetgkc/gowallval/handlers"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter().StrictSlash(true)
	data.LoadFromDB()
	data.LoadFromDB()
	mux.HandleFunc("/", AddressValidationHandler).Methods("POST")

	http.ListenAndServe(":9000", mux)
}
