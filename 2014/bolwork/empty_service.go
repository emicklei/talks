package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/addresses", addressesHandler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func addressesHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello moon")
}
