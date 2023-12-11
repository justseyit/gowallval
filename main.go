package main

import (
	"log"
	"net/http"

	"github.com/seyitahmetgkc/gowallval/data"
	handlers "github.com/seyitahmetgkc/gowallval/handlers"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Wallet Validator is being started...")
	mux := mux.NewRouter().StrictSlash(true)
	data.LoadData()
	mux.HandleFunc("/service", handlers.AddressValidationHandler).Methods("POST")
	mux.HandleFunc("/", HomePageHandler).Methods("GET")

	log.Println("Starting the server on port 9000...")
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home Page")
	w.Write([]byte("Welcome to Wallet Validator!"))
}
