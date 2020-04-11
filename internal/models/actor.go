package models

type Actor interface {
	GetPosition() (float32, float32)
	GetSize() (float32, float32)
	GetAngle() float64
}
