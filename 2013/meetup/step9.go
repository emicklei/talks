package main

import (
	"github.com/ajstarks/svgo"
	"log"
	"math"
	"net/http"
)

const DegToRad = math.Pi / 180

func main() {
	http.HandleFunc("/spiro", spiroHandler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}

// START OMIT

func spiroHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	width := 1000
	height := 1000

	canvas := svg.New(w)
	canvas.Start(width, height)

	spiro := Spiro{
		revolutions: 4,
		angleDelta:  5,
		innerRadius: 30,
		outerRadius: 200,
		innerOffset: 20,
	}

	xcoords, ycoords := computeSpiroPoints(width/2, height/2, spiro)

	canvas.Polyline(xcoords, ycoords, `stroke="blue" stroke-width="1" fill="none"`)
	canvas.End()
}

// END OMIT

type Spiro struct {
	revolutions int
	angleDelta  int
	innerRadius float64
	outerRadius float64
	innerOffset float64
}

func computeSpiroPoints(xc int, yc int, spiro Spiro) (xcoords []int, ycoords []int) {
	xcoords = []int{}
	ycoords = []int{}
	for g := 0; g <= spiro.revolutions*360; g += spiro.angleDelta {
		x, y := computeSpiro(float64(g)*DegToRad, spiro.innerRadius, spiro.outerRadius, spiro.innerOffset)
		xcoords = append(xcoords, round(x)+xc)
		ycoords = append(ycoords, round(y)+yc)
	}
	return
}

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

func round(v float64) int {
	if v < 0 {
		return int(v - 0.5)
	} else {
		return int(v + 0.5)
	}
}
