package render

import (
	"karlc/treegame/internal/game"
	"karlc/treegame/internal/models"

	"github.com/faiface/pixel/pixelgl"
	//"karlc/treegame/internal/models"
	//rl "github.com/gen2brain/raylib-go/raylib"
)

// Camera
type Camera struct {
	PosX           float64
	PosY           float64
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
	OffsetY float64
	OffsetX float64

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
	c.PosX, c.PosY = aX*float64(c.unitScale), aY*float64(c.unitScale)
}

func (c *Camera) drawActor(actor models.Actor) {

	c.updateCameraPosition()

	// TODO: determine if the actor is within the viewport
	// if not, just return

	adjustedX, adjustedY := c.TranslatePosition(actor.GetPosition())
	adjustedW, adjustedH := c.TranslateSize(actor.GetSize())

	angle := actor.GetAngle()

	c.renderer.DrawRect(adjustedX, adjustedY, adjustedW, adjustedH, float64(angle))
}

// TranslatePosition takes a coordinate from the physics simulation and translates it to a
// point in the cameras viewport using the given camera position, viewport size and unitScale/zoom
func (c *Camera) TranslatePosition(x, y float64) (adjustedX, adjustedY float64) {
	adjustedOffsetY := float64(c.viewportHeight) * c.OffsetY / 100
	adjustedOffsetX := float64(c.viewportHeight) * c.OffsetX / 100
	adjustedX = x*float64(c.unitScale) - float64(c.PosX) + adjustedOffsetX + float64(c.viewportWidth/2)
	adjustedY = y*float64(c.unitScale) - float64(c.PosY) + adjustedOffsetY + float64(c.viewportHeight/2)
	return
}

func (c *Camera) TranslateSize(w, h float64) (adjustedW, adjustedH float64) {
	adjustedW = (w * float64(c.unitScale))
	adjustedH = (h * float64(c.unitScale))
	return
}

//func (c *Camera) drawJoint(joint *models.Joint) {

//cA := joint.B2Joint.GetBodyA().GetPosition()
//cB := joint.B2Joint.GetBodyB().GetPosition()

//ax, ay := c.TranslatePosition(float32(cA.X), float32(cA.Y))
//bx, by := c.TranslatePosition(float32(cB.X), float32(cB.Y))

//c.renderer.DrawLine(ax, ay, bx, by, 5)
//}

// DrawGame draws all actors in the game
// TODO: should be complemented with DrawMenu() etc
func (c *Camera) DrawGame(g *game.Game) {
	//rl.BeginDrawing()
	//rl.ClearBackground(rl.Black)
	for _, a := range g.AllActors {
		c.drawActor(a)
	}

	//for _, j := range g.AllJoints {
	//c.drawJoint(j)
	//}

	//rl.EndDrawing()
}

func NewCamera(w, h, scale int, win *pixelgl.Window) *Camera {
	return &Camera{
		PosX:           0,
		PosY:           0,
		viewportWidth:  w,
		viewportHeight: h,
		unitScale:      scale,
		renderer:       NewRenderer(win),
	}
}

func (c *Camera) TestDraw() {
	c.renderer.Test()
}
