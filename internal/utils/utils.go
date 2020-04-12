package utils

import (
	"math/rand"
	"time"
)

func RandFloat32(min, max float32) float32 {
	rand.Seed(time.Now().UnixNano())

	// horrible hack to get negative minimum range
	// up to 3 decimal precision
	i := rand.Intn(int(max*1000)-int(min*1000)) + int(min*1000)
	return float32(i) / 1000.0
}
