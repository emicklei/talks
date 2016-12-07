package convert

import (
	"math"
	"testing"
)

func TestDegreesToRadians(t *testing.T) {
	r := DegreesToRadians(90)
	t.Log(90, "->", r)
	if r != math.Pi/2 {
		t.Fail()
	}
}
