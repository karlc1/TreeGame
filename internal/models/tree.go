package models

type Tree struct {
	Width  float64
	Height float64
	PosX   float64
	Base   float64 // lower y
}

func NewTree(w, h, x, b float64) *Tree {
	return &Tree{
		Width:  w,
		Height: h,
		PosX:   x,
		Base:   b,
	}
}
