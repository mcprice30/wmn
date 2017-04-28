package network

import (
	"math/rand"
)

// dropChance indicates the probability that a given packet is dropped by
// the ethernet connection. Note that the same probability is applied to both
// the data packet and the acknowledgement packet. For example, if this is 50%,
// then the probability of a completely successful transmission is only 25%.
var dropChance = 0.00
var manetDropChance = 0.00

func SetDropChance(dc float64) {
	dropChance = dc
}

func SetManetDropChance(dc float64) {
	manetDropChance = dc
}

// dropFixedRate will return truee iff the gremlin function decides to drop a
// packet, which occurs at the given rate.
func dropFixedRate() bool {
	return rand.Float64() < dropChance
}

func dropManetChance() bool {
	return rand.Float64() < manetDropChance
}
