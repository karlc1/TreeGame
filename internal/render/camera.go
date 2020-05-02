package render

import (
	"fmt"
	"karlc/treegame/internal/game"
	"karlc/treegame/internal/models"
	"karlc/treegame/internal/utils"
	"math"

	"github.com/faiface/pixel/pixelgl"
	//"karlc/treegame/internal/models"
	//rl "github.com/gen2brain/raylib-go/raylib"
)

// Camera
type Camera struct {
	PosX           float64
	PosY           float64
	attachedTo     models.PhysicsActor
	viewportWidth  int
	viewportHeight int

	// unit scale is similar to a zoom level
	// 1.0 size unit is 1 meter in the physical world,
	// scale is how many pixels represent 1 meter
	unitScale float64

	zoom      float64
	zoomList  []float64 // used to average out zoom values for smoother zoom
	zoomIndex int       // used for zoomList

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
func (c *Camera) AttachTo(b models.PhysicsActor) {
	c.attachedTo = b
}

func (c *Camera) updateCameraPosition() {

	aX, aY := c.attachedTo.GetPosition()
	c.PosX, c.PosY = aX*c.zoom, aY*c.zoom

	c.updateZoom()
}

func (c *Camera) updateZoom() {
	if c.zoomIndex == 399 {
		c.zoomIndex = 0
	}
	velX, velY := c.attachedTo.GetLinearVelocity()
	adjustedVel := math.Pow(math.Abs(velX)+math.Abs(velY), 0.6)
	cZoom := c.unitScale - adjustedVel
	c.zoomList[c.zoomIndex] = cZoom
	tot := 0.0
	for _, e := range c.zoomList {
		tot += e
	}
	c.zoom = math.Max(tot/400, 2)
	fmt.Println(c.zoom)
	c.zoomIndex++
}

// isWithinView determines if an actor is visible on screen
// the arguments should have been translated according to
// camera position and unit size before calling this
func (c *Camera) isWithinView(x, y, w, h float64) bool {
	return x > 0-w/2 &&
		x < float64(c.viewportWidth)+w/2 &&
		y > 0-h/2 &&
		y < float64(c.viewportHeight)+w/2
}

var drawnActors = 1

func (c *Camera) drawActor(actor models.Actor) {

	adjustedX, adjustedY := c.TranslatePosition(actor.GetPosition())
	adjustedW, adjustedH := c.TranslateSize(actor.GetSize())

	//TODO: does not account for rotation, problem later?
	if !c.isWithinView(adjustedX, adjustedY, adjustedW, adjustedH) {
		return
	}

	drawnActors++

	angle := actor.GetAngle()

	c.renderer.DrawRect(adjustedX, adjustedY, adjustedW, adjustedH, float64(angle))
}

// TranslatePosition takes a coordinate from the physics simulation and translates it to a
// point in the cameras viewport using the given camera position, viewport size and unitScale/zoom
func (c *Camera) TranslatePosition(x, y float64) (adjustedX, adjustedY float64) {
	adjustedOffsetY := float64(c.viewportHeight) * c.OffsetY / 100
	adjustedOffsetX := float64(c.viewportHeight) * c.OffsetX / 100
	adjustedX = x*c.zoom - float64(c.PosX) + adjustedOffsetX + float64(c.viewportWidth/2)
	adjustedY = y*c.zoom - float64(c.PosY) + adjustedOffsetY + float64(c.viewportHeight/2)
	return
}

func (c *Camera) TranslateSize(w, h float64) (adjustedW, adjustedH float64) {
	adjustedW = (w * c.zoom)
	adjustedH = (h * c.zoom)
	return
}

func (c *Camera) drawJoint(joint *models.Joint) {
	cA := joint.B2Joint.GetBodyA().GetPosition()
	cB := joint.B2Joint.GetBodyB().GetPosition()

	aax := cA.X + joint.AnchorAX
	aay := cA.Y + joint.AnchorAY
	abx := cB.X + joint.AnchorBX
	aby := cB.Y + joint.AnchorBY

	tcax, tcay := c.TranslatePosition(cA.X, cA.Y)
	tcbx, tcby := c.TranslatePosition(cB.X, cB.Y)
	tax, tay := c.TranslatePosition(aax, aay)
	tbx, tby := c.TranslatePosition(abx, aby)

	ra := joint.B2Joint.GetBodyA().GetAngle()
	rb := joint.B2Joint.GetBodyB().GetAngle()

	rax, ray := utils.RotatePoint(tcax, tcay, tax, tay, ra)
	rbx, rby := utils.RotatePoint(tcbx, tcby, tbx, tby, rb)

	c.renderer.DrawLine(rax, ray, rbx, rby)
}

// DrawGame draws all actors in the game
// TODO: should be complemented with DrawMenu() etc
func (c *Camera) DrawGame(g *game.Game) {
	//fmt.Printf("DrawnActors: %d \n", drawnActors)
	drawnActors = 0

	c.updateCameraPosition()

	for _, a := range g.AllActors {
		c.drawActor(a)
	}

	if g.Rope != nil {
		c.drawJoint(g.Rope)
	}

	//for _, j := range g.AllJoints {
	//c.drawJoint(j)
	//}
}

func NewCamera(w, h int, scale float64, win *pixelgl.Window) *Camera {

	// TODO: adjust number of avg vals
	avgVals := 400
	zoomList := make([]float64, avgVals, avgVals)
	for i := 0; i < avgVals; i++ {
		zoomList[i] = scale
	}

	return &Camera{
		PosX:           0,
		PosY:           0,
		viewportWidth:  w,
		viewportHeight: h,
		unitScale:      scale,
		renderer:       NewRenderer(win),
		zoomList:       zoomList,
		zoom:           scale,
	}
}

// remove later, used to experiment with pixelgl easily
func (c *Camera) TestDraw() {
	c.renderer.Test()
}
