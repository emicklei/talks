package main

import (
	"encoding/json"
	"os"
)

type Address struct {
	Postalcode  string
	HouseNumber string
	Street      string
	City        string
	Country     string
}

//START OMIT
func main() {
	a := Address{
		Postalcode:  "100AB",
		HouseNumber: "12",
		Street:      "Keulsekade",
		City:        "Utrecht",
		Country:     "Netherlands",
	}
	json.NewEncoder(os.Stdout).Encode(a)
}

//END OMIT
