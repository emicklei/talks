package convert

import "math"

func DegreesToRadians(d int) float64 {
	return float64(d) * math.Pi / 180.0
}
