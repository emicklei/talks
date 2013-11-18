package main

import (
	"github.com/ajstarks/svgo"
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

	canvas := svg.New(w)
	canvas.Start(500, 500)
	canvas.Circle(100, 50, 40, `stroke="black" stroke-width="2" fill="red"`)
	canvas.End()
}

// END OMIT
