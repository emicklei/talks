package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/spiro", spiroHandler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func spiroHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "spiro")
}
