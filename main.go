package main

import (
	"log"
	"net/http"

	"github.com/seyitahmetgkc/gowallval/data"
	handlers "github.com/seyitahmetgkc/gowallval/handlers"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Wallet Validator Has Started")
	mux := mux.NewRouter().StrictSlash(true)
	data.LoadFromDB()
	mux.HandleFunc("/service", handlers.AddressValidationHandler).Methods("POST")
	mux.HandleFunc("/", HomePageHandler).Methods("GET")

	log.Println("Starting the server on port 9000")
	http.ListenAndServe(":9000", mux)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home Page")
	w.Write([]byte("Welcome to Wallet Validator!"))
}
