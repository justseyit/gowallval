package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/seyitahmetgkc/gowallval/models"
)

var AppCurrencies []Currency
var AppNetworks []Network
var AppCurrencyNetworkMappings []CurrencyNetworkMapping
var appDB *sql.DB
var UnsupportedNetworks = []string{"STX", "MIOTA", "NXS"}
var UnsupportedCurrencies = []string{"STX", "IOTA","NXS"}

func LoadFromDB() {
	appDB = openConnection()
	defer closeConnection(appDB)
	initializeValidators()
	AppNetworks = loadNetworks()
	AppCurrencies = loadCurrencies()
	AppCurrencyNetworkMappings = loadCurrencyNetworkMappings()
}

func openConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root.seyit122@/walletvalidator")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func closeConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func loadNetworks() []Network {
	rows, err := appDB.Query("SELECT * FROM network")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var networks []Network
	for rows.Next() {
		var network Network
		err := rows.Scan(&network.ID, &network.Symbol, &network.Name)
		if err != nil {
			log.Fatal(err)
		}
		networks = append(networks, network)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return networks
}

func loadCurrencies() []Currency {
	rows, err := appDB.Query("SELECT * FROM currency")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var currencies []Currency
	for rows.Next() {
		var currency Currency
		err := rows.Scan(&currency.ID, &currency.Symbol, &currency.Name)
		if err != nil {
			log.Fatal(err)
		}
		currencies = append(currencies, currency)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return currencies
}

func loadCurrencyNetworkMappings() []CurrencyNetworkMapping {
	rows, err := appDB.Query("SELECT * FROM currencynetwork")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var mappings []CurrencyNetworkMapping
	for rows.Next() {
		var mapping CurrencyNetworkMapping
		err := rows.Scan(&mapping.CurrencyID, &mapping.NetworkID)
		if err != nil {
			log.Fatal(err)
		}
		mappings = append(mappings, mapping)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return mappings
}
