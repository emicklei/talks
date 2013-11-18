package main

import (
	"github.com/ajstarks/svgo"
	"log"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/spiro", spiroHandler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func spiroHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	canvas := svg.New(w)
	canvas.Start(500, 500)
	canvas.Circle(100, 50, 40, `stroke="black" stroke-width="2" fill="red"`)
	canvas.End()
}

// START OMIT

func computeSpiro(t float64, r float64, R float64, offset float64) (x, y float64) {
	k := r / R
	l := offset / r
	x1 := (1 - k) * math.Cos(t)
	x2 := l * k * math.Cos(((1-k)/k)*t)
	x = R * (x1 + x2)
	y1 := (1 - k) * math.Sin(t)
	y2 := l * k * math.Sin(((1-k)/k)*t)
	y = R * (y1 - y2)
	return
}

// END OMIT
