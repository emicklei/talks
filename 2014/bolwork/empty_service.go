package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hi", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "gopher says hi")
}
