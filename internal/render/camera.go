package render

import (
	"fmt"
	"karlc/treegame/internal/game"
	"karlc/treegame/internal/models"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera struct {
	PosX           float32
	PosY           float32
	attachedTo     models.Actor
	viewportWidth  int
	viewportHeight int

	// unit scale is similar to a zoom level
	// 1.0 size unit is 1 meter in the physical world,
	// scale is how many pixels represent 1 meter
	unitScale int

	// OffsetY denotes how much to offset the camera
	// in relation to the attached player. The value is
	// given in percentage of screen size
	OffsetY float32

	renderer *Renderer
}

// AttachTo lets the camera attach to an actor and
// follow it. Probably attached to the player character
// most of the time
func (c *Camera) AttachTo(b models.Actor) {
	c.attachedTo = b
}

func (c *Camera) updateCameraPosition() {
	aX, aY := c.attachedTo.GetPosition()
	c.PosX, c.PosY = aX*float32(c.unitScale), aY*float32(c.unitScale)
}

func (c *Camera) drawActor(actor models.Actor, debug bool) {

	c.updateCameraPosition()

	// TODO: determine if the actor is within the viewport
	// if not, just return

	adjustedOffsetY := float32(c.viewportHeight) * c.OffsetY / 100

	aX, aY := actor.GetPosition()
	adjustedX := aX*float32(c.unitScale) - c.PosX + float32(c.viewportWidth/2)
	adjustedY := aY*float32(c.unitScale) - c.PosY + adjustedOffsetY + float32(c.viewportHeight/2)

	aW, aH := actor.GetSize()

	adjustedW := (aW * float32(c.unitScale)) // TODO: make less ugly
	adjustedH := (aH * float32(c.unitScale)) // TODO: make less ugly

	angle := actor.GetAngle()

	if debug {
		fmt.Printf("W: %v \n H: %v \n X: %v \n Y: %v \n \n", adjustedW, adjustedH, adjustedX, adjustedY)
	}

	c.renderer.DrawRect(adjustedX, adjustedY, adjustedW, adjustedH, float32(angle))
}

func (c *Camera) drawJoint(joint *models.Joint) {

	cA := joint.B2Joint.GetBodyA().GetLocalCenter()
	cB := joint.B2Joint.GetBodyB().GetLocalCenter()

	adjustedOffsetY := float32(c.viewportHeight) * c.OffsetY / 100

	adjustedAX := float32(cA.X)*float32(c.unitScale) - c.PosX + float32(c.viewportWidth/2)
	adjustedAY := float32(cA.Y)*float32(c.unitScale) - c.PosY + adjustedOffsetY + float32(c.viewportHeight/2)

	adjustedBX := float32(cB.X)*float32(c.unitScale) - c.PosX + float32(c.viewportWidth/2)
	adjustedBY := float32(cB.Y)*float32(c.unitScale) - c.PosY + adjustedOffsetY + float32(c.viewportHeight/2)

	rl.DrawLine(
		int32(c.renderer.screenWidth)-int32(adjustedAX),
		int32(c.renderer.screenHeight)-int32(adjustedAY),
		int32(c.renderer.screenWidth)-int32(adjustedBX),
		int32(c.renderer.screenHeight)-int32(adjustedBY),
		rl.White,
	)
}

// DrawGame draws all actors in the game
// TODO: should be complemented with DrawMenu() etc
func (c *Camera) DrawGame(g *game.Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	for _, a := range g.AllActors {
		c.drawActor(a, false)
	}

	for _, j := range g.AllJoints {
		c.drawJoint(j)
	}

	rl.EndDrawing()
}

func NewCamera(w, h int, scale int) *Camera {
	return &Camera{
		PosX:           0,
		PosY:           0,
		renderer:       NewRenderer(w, h),
		viewportWidth:  w,
		viewportHeight: h,
		unitScale:      scale,
	}
}
