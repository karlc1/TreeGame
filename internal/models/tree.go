package models

import (
	"karlc/treegame/internal/utils"
	"math"
)

type Tree struct {
	Width          float64
	Height         float64
	PosX           float64
	Base           float64 // lower y
	Circumference  float64
	Decor          []*DecorBox
	RotationOffset float64 // not angle, but how many units player has moved?
}

func NewTree(w, h, x, b float64) *Tree {
	t := &Tree{
		Width:         w,
		Height:        h,
		PosX:          x,
		Base:          b,
		Circumference: 2 * math.Pi * (w / 2),
	}
	t.InitTreeDecor()
	return t
}

func (t *Tree) InitTreeDecor() {
	n := 50
	for i := 0; i < n; i++ {
		xSize := 0.4
		ySize := utils.RandFloat64(0.5, 1)
		posX := utils.RandFloat64(-(t.Circumference / 2), t.Circumference/2)
		posY := utils.RandFloat64(0, t.Height)

		t.Decor = append(t.Decor, &DecorBox{
			Height: ySize,
			Width:  xSize,
			PosX:   posX,
			PosY:   posY,
		})
	}
}
