package calculation

import "math/rand"

func RandomNumber(size int) int {
	value := rand.Intn(size)
	return value
}
