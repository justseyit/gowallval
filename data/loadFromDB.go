package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	models "github.com/seyitahmetgkc/gowallval/models"
)

var AppCurrencies []models.Currency
var AppNetworks []models.Network
var AppCurrencyNetworkMappings []models.CurrencyNetworkMapping
var appDB *sql.DB
var UnsupportedNetworks = []string{"STX", "MIOTA", "NXS"}
var UnsupportedCurrencies = []string{"STX", "IOTA", "NXS"}

func LoadFromDB() {
	appDB = openConnection()
	defer closeConnection(appDB)
	initializeValidators()
	AppNetworks = loadNetworks()
	AppCurrencies = loadCurrencies()
	AppCurrencyNetworkMappings = loadCurrencyNetworkMappings()
}

func openConnection() *sql.DB {
	db, err := sql.Open("mysql", DBUser+":"+DBPass+"@/"+DBName)
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

func loadNetworks() []models.Network {
	rows, err := appDB.Query("SELECT * FROM network")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var networks []models.Network
	for rows.Next() {
		var network models.Network
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

func loadCurrencies() []models.Currency {
	rows, err := appDB.Query("SELECT * FROM currency")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var currencies []models.Currency
	for rows.Next() {
		var currency models.Currency
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

func loadCurrencyNetworkMappings() []models.CurrencyNetworkMapping {
	rows, err := appDB.Query("SELECT * FROM currencynetwork")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var mappings []models.CurrencyNetworkMapping
	for rows.Next() {
		var mapping models.CurrencyNetworkMapping
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
