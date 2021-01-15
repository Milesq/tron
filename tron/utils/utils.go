package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random .
func Random(a, b int) int {
	return rand.Intn(b-a) + a
}
