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

	xcoords := []int{50, 100, 150}
	ycoords := []int{150, 50, 100}

	canvas.Polyline(xcoords, ycoords, `stroke="blue" stroke-width="1" fill="none"`)
	canvas.End()
}

// END OMIT
