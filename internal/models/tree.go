package models

import (
	"karlc/treegame/internal/utils"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Tree struct {
	Width          float64
	Height         float64
	PosX           float64
	Base           float64 // lower y
	Circumference  float64
	Decor          []*DecorBox
	RotationOffset float64 // not angle, but how many units player has moved?
	SpriteVecs     []pixel.Vec
	Sprite         *pixel.Sprite
	SpritePicture  *pixel.PictureData
	Canvas         *pixelgl.Canvas
}

func NewTree(w, h, x, b float64) *Tree {
	t := &Tree{
		Width:         w,
		Height:        h,
		PosX:          x,
		Base:          b,
		Circumference: 2 * math.Pi * (w / 2),
		Canvas:        pixelgl.NewCanvas(pixel.ZR),
	}
	t.InitTreeDecor()
	t.InitSprite()
	return t
}

func (t *Tree) InitSprite() {
	t.SpritePicture = utils.LoadPicture("./assets/sprites/treetile.jpg")
	t.Sprite = utils.SpriteFromPic(t.SpritePicture)

}

func (t *Tree) InitSprites() {
	t.InitSprite()
	spriteVecs := make([]pixel.Vec, 0, 0)
	for x := 0.0; x < t.Width; x += t.Width / t.Sprite.Frame().W() {
		for y := 0.0; y < t.Height; y += t.Height / t.Sprite.Frame().H() {
			spriteVecs = append(spriteVecs, pixel.Vec{X: x, Y: y})
		}
	}
	t.SpriteVecs = spriteVecs
}

func (t *Tree) InitTreeDecor() {
	n := 50
	for i := 0; i < n; i++ {
		xSize := 0.4
		ySize := utils.RandFloat64(0.5, 1)
		posX := utils.RandFloat64(0, t.Circumference)
		posY := utils.RandFloat64(0, t.Height)

		t.Decor = append(t.Decor, &DecorBox{
			Height: ySize,
			Width:  xSize,
			PosX:   posX,
			PosY:   posY,
		})
	}
}
