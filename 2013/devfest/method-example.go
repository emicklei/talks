type Point struct {
	X, Y int
}

func (p Point) Plus(a Point) Point {
	return Point{p.X + a.X, p.Y + a.Y}
}