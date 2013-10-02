package main

import (
	"fmt"
	"math"
)

// START OMIT
type Vector struct {
	X, Y int
}

func (v Vector) Length() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func main() {
	v := Vector{3, 4}
	fmt.Printf("%g\n", v.Length())
}

// END OMIT
