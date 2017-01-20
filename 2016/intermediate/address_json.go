package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/addresses", addressHandler)
	http.ListenAndServe(":8080", nil)
}

// START OMIT
type Address struct {
	Postalcode  string `json:"postal-code,omitempty" xml:"PostalCode"`
	HouseNumber int    `json:"house-number"`
	Street      string `json:"street,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
}

func addressHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("POST expected, got %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	a := Address{}
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		log.Println("Address decode failed, got error %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("%#v\n", a)
	if err := json.NewEncoder(w).Encode(a); err != nil {
		log.Println("Address encode failed, got error %v", err)
		return
	}
}

// END OMIT
