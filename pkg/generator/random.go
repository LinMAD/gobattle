package generator

import (
	"math/rand"
	"time"
)

// RandomBool return random bool
func RandomBool() bool {
	c := make(chan bool)
	close(c)
	select {
	case <-c:
		return true
	case <-c:
		return false
	}
}

// RandomNum return random num between min and max
func RandomNum(min, max int) int8 {
	rand.Seed(time.Now().UnixNano())
	return int8(rand.Intn(max-min) + min)
}
