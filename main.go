package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/seyitahmetgkc/gowallval/data"
	. "github.com/seyitahmetgkc/gowallval/handlers"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting Wallet Validator...")
	mux := mux.NewRouter().StrictSlash(true)
	data.LoadFromDB()
	data.LoadFromDB()
	mux.HandleFunc("/service", AddressValidationHandler).Methods("POST")
	mux.HandleFunc("/", HomePageHandler).Methods("GET")

	log.Println("Starting the server on port 9000")
	http.ListenAndServe(":9000", mux)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home Page")
	w.Write([]byte("Welcome to Wallet Validator!"))
}
