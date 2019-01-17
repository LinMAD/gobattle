package generator

import (
	"math/rand"
	"time"
)

// RandomBool return random bool
func RandomBool() bool {
	return rand.Intn(2) == 0
}

// RandomNum return random num between min and max
func RandomNum(min, max int) int8 {
	rand.Seed(time.Now().UnixNano())

	return int8(rand.Intn(max-min) + min)
}
