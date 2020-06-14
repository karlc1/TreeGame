package render

import (
	"fmt"
	"karlc/treegame/internal/game"
	"karlc/treegame/internal/models"
	"karlc/treegame/internal/utils"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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

	TreeView bool
}

// AttachTo lets the camera attach to an actor and
// follow it. Probably attached to the player character
// most of the time
func (c *Camera) AttachTo(b models.PhysicsActor) {
	c.attachedTo = b
}

func (c *Camera) updateCameraPosition() {

	if c.TreeView {
		_, aY := c.attachedTo.GetPosition()
		c.PosY = aY * c.zoom
		return
	}

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

func (c *Camera) drawPlayerNormal() {
	c.drawActor(c.attachedTo)
}

func (c *Camera) drawPlayerTree() {
	player := c.attachedTo
	adjustedW, adjustedH := c.TranslateSize(player.GetSize())
	_, adjustedY := c.TranslatePosition(player.GetPosition())

	//adjustedOffsetX := float64(c.viewportHeight) * c.OffsetX / 100
	adjustedX := float64(c.viewportWidth / 2)

	angle := player.GetAngle()
	c.renderer.DrawRect(adjustedX, adjustedY, adjustedW, adjustedH, float64(angle))
}

func (c *Camera) drawTreeDecor(tree *models.Tree) {

	//c.renderer.DrawText(500, 500, fmt.Sprintf("%v", int(0-(tree.Width/2))))
	for _, d := range tree.Decor {

		relPosX := d.PosX + tree.RotationOffset
		// check if decor is placed behind currently visible half of tree
		if relPosX < 0-tree.Width/2 || relPosX > tree.Width/2 {
			continue
		}

		// adjusted for tree position
		tpx := tree.PosX + d.PosX + tree.RotationOffset
		tpy := tree.Base + d.PosY

		adjustedX, adjustedY := c.TranslatePosition(tpx, tpy)
		adjustedW, adjustedH := c.TranslateSize(d.GetSize())

		c.renderer.DrawRectRed(adjustedX, adjustedY, adjustedW, adjustedH, 0)
		//c.renderer.DrawText(adjustedX, adjustedY, fmt.Sprintf("%v", int(relPosX)))

		//var xx float64
		//if d.PosX >= tree.RotationOffset {
		//xx = tree.RotationOffset - d.PosX
		//} else {
		//xx = tree.RotationOffset + (tree.RotationOffset - d.PosX)
		//}

		//adjustedX, adjustedY := c.TranslatePosition(xx, tree.Base+d.Height)
		//adjustedW, adjustedH := c.TranslateSize(d.Width, d.Height)
		//c.renderer.DrawRect(adjustedX, adjustedY, adjustedW, adjustedH, 0)
	}
}

func (c *Camera) updateTreeRotation(tree *models.Tree, player models.PhysicsActor) {
	px, _ := player.GetPosition()
	tree.RotationOffset = tree.PosX - px
}

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

	if c.TreeView {
		c.updateTreeRotation(g.Tree, g.Player)
		c.drawTreeDecor(g.Tree)
		c.drawPlayerTree()
	} else {
		c.drawPlayerNormal()
	}

	for _, a := range g.AllActors {
		c.drawActor(a)
	}

	if g.Rope != nil {
		c.drawJoint(g.Rope)
	}

	if g.Tree != nil {
		c.DrawTree(g.Tree)
	}

	c.renderer.DrawTreeTile(g.Tree)

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

func (c *Camera) DrawTree(tree *models.Tree) {

	//blx, bly := tree.PosX-tree.Width/2, tree.Base
	//brx, bry := tree.PosX+tree.Width/2, tree.Base
	//tlx, tly := tree.PosX-tree.Width/2, tree.Base+tree.Height
	//trx, try := tree.PosX+tree.Width/2, tree.Base+tree.Height

	//tblx, tbly := c.TranslatePosition(blx, bly)
	//tbrx, tbry := c.TranslatePosition(brx, bry)
	//ttlx, ttly := c.TranslatePosition(tlx, tly)
	//ttrx, ttry := c.TranslatePosition(trx, try)

	//c.renderer.DrawLine(tblx, tbly, ttlx, ttly)
	//c.renderer.DrawLine(tbrx, tbry, ttrx, ttry)
	//c.renderer.DrawLine(ttlx, ttly, ttrx, ttry)

	//////

	treeCenterX, treeCenterY := c.TranslatePosition(
		tree.PosX,
		(tree.Height+(tree.Base*2))/2,
	)
	minX, minY := c.TranslatePosition(
		tree.PosX-tree.Width/2,
		tree.Base,
	)
	c.renderer.DrawPoint(minX, minY)
	c.renderer.DrawPoint(treeCenterX, treeCenterY)
	maxX, maxY := c.TranslatePosition(
		tree.PosX+tree.Width/2,
		tree.Base+tree.Height,
	)
	c.renderer.DrawPoint(maxX, maxY)
	tree.Canvas.SetBounds(pixel.R(minX, minY, maxX, maxY))
	tree.Canvas.Clear(colornames.Brown)

	////////////// Sprite /////////

	spriteWidth, spriteHeight := c.TranslateSize(
		tree.Sprite.Frame().W(),
		tree.Sprite.Frame().H(),
	)
	_ = spriteWidth
	_ = spriteHeight

	scale := c.zoom / c.unitScale

	spriteWidth = tree.Sprite.Frame().W() * scale
	spriteHeight = tree.Sprite.Frame().H() * scale

	fmt.Println(spriteWidth)

	tree.Canvas.SetSmooth(true)
	mat := pixel.IM
	//mat = mat.Scaled(tree.Sprite.Frame().Center(), scale)
	//mat = mat.Moved(tree.Canvas.Bounds().Center())
	mat = mat.Scaled(tree.Sprite.Frame().Min, scale)

	mat = mat.Moved(pixel.V(minX+spriteWidth/2, minY+spriteHeight/2))
	tree.Sprite.Draw(tree.Canvas, mat)

	////////////

	tree.Canvas.Draw(c.renderer.window, pixel.IM.Moved(pixel.Vec{X: treeCenterX, Y: treeCenterY}))
	c.renderer.DrawCanvas(tree)

}

func (c *Camera) DrawTreeSprites() {

}
