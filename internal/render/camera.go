package render

import (
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

	// Offset denotes how much to offset the camera
	// in relation to the attached actor. The value is
	// given in percentage of screen size
	OffsetY float32
	OffsetX float32

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

func (c *Camera) drawActor(actor models.Actor) {

	c.updateCameraPosition()

	// TODO: determine if the actor is within the viewport
	// if not, just return

	adjustedX, adjustedY := c.TranslatePosition(actor.GetPosition())
	adjustedW, adjustedH := c.TranslateSize(actor.GetSize())

	angle := actor.GetAngle()

	c.renderer.DrawRect(adjustedX, adjustedY, adjustedW, adjustedH, float32(angle))
}

// TranslatePosition takes a coordinate from the physics simulation and translates it to a
// point in the cameras viewport using the given camera position, viewport size and unitScale/zoom
func (c *Camera) TranslatePosition(x, y float32) (adjustedX, adjustedY float32) {
	adjustedOffsetY := float32(c.viewportHeight) * c.OffsetY / 100
	adjustedOffsetX := float32(c.viewportHeight) * c.OffsetX / 100
	adjustedX = x*float32(c.unitScale) - c.PosX + adjustedOffsetX + float32(c.viewportWidth/2)
	adjustedY = y*float32(c.unitScale) - c.PosY + adjustedOffsetY + float32(c.viewportHeight/2)
	return
}

func (c *Camera) TranslateSize(w, h float32) (adjustedW, adjustedH float32) {
	adjustedW = (w * float32(c.unitScale))
	adjustedH = (h * float32(c.unitScale))
	return
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
		c.drawActor(a)
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
