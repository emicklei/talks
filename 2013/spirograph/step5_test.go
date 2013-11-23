package main

import (
	"fmt"
	"math"
	"testing"
)

// go test -v step5_test.go

// START OMIT
const DegToRad = math.Pi / 180

func TestComputeSpiro(t *testing.T) {
	for g := 0.0; g <= 360.0; g += 10.0 {
		x, y := computeSpiro(g*DegToRad, 10, 100, 2)

		// this time no assertions, just dump
		//
		fmt.Printf("(%f,%f)\n", x, y)
	}
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
