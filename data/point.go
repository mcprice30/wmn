package data

import (
	"math"
)

// Point encapsulates an (X,Y) coordinate.
type Point struct {
	X int
	Y int
}

// Dist will calculate the distance from this point to another point.
func (p *Point) Dist(o *Point) float64 {
	dx := p.X - o.X
	dy := p.Y - o.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
}
