package models

type Actor interface {
	GetPosition() (float64, float64)
	GetSize() (float64, float64)
	GetAngle() float64
	GetZVal() int
}
