package main

import (
	"github.com/ajstarks/svgo"
	"log"
	"math"
	"net/http"
	"strconv"
)

const DegToRad = math.Pi / 180

func main() {
	http.HandleFunc("/spiro", spiroHandler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func spiroHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	// START OMIT
	spiro := Spiro{
		revolutions: intParameter(r, "rev", 3),
		angleDelta:  float64Parameter(r, "delta", 2),
		innerRadius: float64Parameter(r, "inner", 30),
		outerRadius: float64Parameter(r, "outer", 200),
		innerOffset: float64Parameter(r, "offset", 45),
	}
	width := intParameter(r, "width", 1000)
	height := intParameter(r, "height", 1000)
	canvas := svg.New(w)
	canvas.Start(width, height)
	xcoords, ycoords := spiro.computeCoordinates(width/2, height/2)
	canvas.Polyline(xcoords, ycoords, `stroke="blue" stroke-width="1" fill="none"`)
	canvas.End()
}

func float64Parameter(r *http.Request, name string, missing float64) (value float64) {
	// END OMIT
	value = missing
	param := r.URL.Query().Get(name)
	if param != "" {
		conv, err := strconv.ParseFloat(param, 64)
		if err == nil { // invalid means missing
			value = conv
		}
	}
	return
}

func intParameter(r *http.Request, name string, missing int) (value int) {
	value = missing
	param := r.URL.Query().Get(name)
	if param != "" {
		conv, err := strconv.Atoi(param)
		if err == nil { // invalid means missing
			value = conv
		}
	}
	return
}

type Spiro struct {
	revolutions int
	angleDelta  float64
	innerRadius float64
	outerRadius float64
	innerOffset float64
}

func (s Spiro) computeCoordinates(xc int, yc int) (xcoords []int, ycoords []int) {
	xcoords = []int{}
	ycoords = []int{}
	for g := 0.0; g <= float64(s.revolutions)*360; g += s.angleDelta {
		x, y := computeSpiro(float64(g)*DegToRad, s.innerRadius, s.outerRadius, s.innerOffset)
		xcoords = append(xcoords, int(x)+xc)
		ycoords = append(ycoords, int(y)+yc)
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
