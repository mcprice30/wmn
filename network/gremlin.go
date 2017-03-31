package network

import (
	"math/rand"
)

// dropFixedRate will return truee iff the gremlin function decides to drop a
// packet, which occurs at the given rate.
func dropFixedRate(rate float64) bool {
	return rand.Float64() < rate
}

// dropDistance will return true iff the gremlin function decides to drop a
// packet when transmitting over the given distance.
func dropDistance(dist float64) bool {
	return rand.Float64() > 1.0 - (dist / 125.0) * (dist / 125.0)
}
