import "math"

// START OMIT
type Surface interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return float64(r.Width * r.Height)
}

func (c Circle) Area() float64 {
	return 2 * math.Pi * float64(c.Radius)
}

// END OMIT