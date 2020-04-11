package models

type Game struct {
	PhysicalWorld PhysicalWorld
	WorldSizeX    int
	WorldSizeY    int
}

type Camera struct {
	PosX       int
	PosY       int
	attachedTo *Box
}

func (c *Camera) AttachTo(b *Box) {
	c.attachedTo = b
}

func NewCamera() *Camera {
	return &Camera{PosX: 0, PosY: 0}
}
