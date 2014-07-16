package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Address struct {
	Postalcode  string
	HouseNumber string
	Street      string
	City        string
	Country     string
}

// START OMIT
func main() {
	http.HandleFunc("/addresses", addressesHandler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func addressesHandler(w http.ResponseWriter, r *http.Request) {
	address := Address{}
	json.NewEncoder(w).Encode(address)
}

// END OMIT
