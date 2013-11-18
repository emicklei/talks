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

// START OMIT
func spiroHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	io.WriteString(w, svg)
}

var svg = `<?xml version="1.0" standalone="no"?>
<svg xmlns="http://www.w3.org/2000/svg" version="1.1">
  <circle cx="100" cy="50" r="40" stroke="black" stroke-width="2" fill="red" />
</svg>`

// END OMIT
