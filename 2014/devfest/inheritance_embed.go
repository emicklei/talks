package main

import (
	"fmt"
)

// START OMIT
type Positioned struct {
	X, Y int
}

func (t *Positioned) Move(dx, dy int) {
	t.X += dx
	t.Y += dy
}

type Box struct {
	Positioned
	W, H int
}

func main() {
	b := Box{Positioned{1, 2}, 3, 4}
	b.Move(1, 1)

	// END OMIT
	fmt.Println(b)
}

func (b Box) String() string {
	return fmt.Sprintf("(%d,%d) extent: (%d,%d)", b.X, b.Y, b.W, b.H)
}
