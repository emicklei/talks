package main

import (
	"log"
	"math"
	"net/http"

	"github.com/ajstarks/svgo"
)

const DegToRad = math.Pi / 180

func main() {
	http.HandleFunc("/spiro", spiroHandler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}

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

	xcoords, ycoords := spiro.computeCoordinates(width/2, height/2)

	canvas.Polyline(xcoords, ycoords, `stroke="blue" stroke-width="1" fill="none"`)
	canvas.End()
}

type Spiro struct {
	revolutions int
	angleDelta  float64
	innerRadius float64
	outerRadius float64
	innerOffset float64
}

// START OMIT
func (s Spiro) computeCoordinates(xc int, yc int) (xcoords []int, ycoords []int) {
	for g := 0.0; g <= float64(s.revolutions)*360; g += s.angleDelta {
		x, y := computeSpiro(float64(g)*DegToRad, s.innerRadius, s.outerRadius, s.innerOffset)
		xcoords = append(xcoords, int(x)+xc)
		ycoords = append(ycoords, int(y)+yc)
	}
	return
}

// END OMIT

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
