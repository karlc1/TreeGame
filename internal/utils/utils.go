package utils

import (
	"math"
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

func RandFloat64(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())

	// horrible hack to get negative minimum range
	// up to 3 decimal precision
	i := rand.Intn(int(max*1000)-int(min*1000)) + int(min*1000)
	return float64(i) / 1000.0
}

func RotateRectCorner(centerX, centerY, cornerX, cornerY, angle float64) (Rx, Ry float64) {
	Rx = centerX + (cornerX * math.Cos(angle)) - (cornerY * math.Sin(angle))
	Ry = centerY + (cornerX * math.Sin(angle)) + (cornerY * math.Cos(angle))
	return
}
