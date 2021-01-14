package utils

import (
	"math/rand"
)

// Random .
func Random(a, b int) int {
	return rand.Intn(b-a) + a
}
