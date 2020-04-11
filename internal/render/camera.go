package render

import (
	"karlc/treegame/internal/models"
)

type Camera struct {
	PosX           float32
	PosY           float32
	attachedTo     *models.Actor
	viewportWidth  int
	viewportHeight int

	// unit scale is similar to a zoom level
	// 1.0 size unit is 1 meter in the physical world,
	// scale is how many pixels represent 1 meter
	unitScale int

	renderer Renderer
}

// AttachTo lets the camera attach to an actor and
// follow it. Probably attached to the player character
// most of the time
func (c *Camera) AttachTo(b *models.Actor) {
	c.attachedTo = b
}

func (c *Camera) DrawActor(actor models.Actor) {

	// TODO: determine if the actor is within the viewport
	// if not, just return

	aX, aY := actor.GetPosition()
	adjustedX := aX - c.PosX
	adjustedY := aY - c.PosY

	aW, aH := actor.GetSize()

	adjustedW := (aW * float32(c.unitScale)) // TODO: make less ugly
	adjustedH := (aH * float32(c.unitScale)) // TODO: make less ugly

	angle := actor.GetAngle()

	c.renderer.DrawRect(adjustedX, adjustedY, adjustedW, adjustedH, float32(angle))
}

func NewCamera() *Camera {
	return &Camera{PosX: 0, PosY: 0}
}
