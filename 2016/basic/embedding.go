package main

import (
	"fmt"
	"math"
)

// START OMIT
type Point struct {
	X, Y int
}
type Shape struct {
	Center Point
}
type Circle struct {
	Shape
	Radius int
}

func main() {
	c := Circle{
		Shape: Shape{
			Center: Point{1, 2},
		}, Radius: 3}

	fmt.Printf("%#v\n", c)
	fmt.Printf("%v", c.Center)
}

// END OMIT

type Rectangle struct {
	Width, Height int
}

type Surface interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return float64(r.Width * r.Height)
}

func (c Circle) Area() float64 {
	return 2 * math.Pi * float64(c.Radius)
}
