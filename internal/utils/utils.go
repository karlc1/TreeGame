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

// rotate rotates a corner around a center point theta radians
func RotatePoint(centerX, centerY, pointX, pointY, theta float64) (float64, float64) {
	tempX, tempY := pointX-centerX, pointY-centerY
	rotatedX := tempX*math.Cos(theta) - tempY*math.Sin(theta)
	rotatedY := tempX*math.Sin(theta) + tempY*math.Cos(theta)
	return rotatedX + centerX, rotatedY + centerY
}

// DegToRad coverts degrees to radians
func DegToRad(deg float64) float64 {
	return deg * 0.0174532925199432957
}

// RadToDeg converts radians to degrees
func RadToDeg(rad float64) float64 {
	return rad * 57.295779513082320876
}
