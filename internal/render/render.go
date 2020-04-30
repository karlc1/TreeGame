package render

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	//rl "github.com/gen2brain/raylib-go/raylib"
)

// Renderer should by used to draw on screen
// Any position and scale transformations should
// be done by the camera and sent to the renderer
// for drawing only
type Renderer struct {
	window *pixelgl.Window
}

func NewRenderer(window *pixelgl.Window) *Renderer {
	return &Renderer{
		window: window,
	}
}

// DrawLine between point a and b
func (r *Renderer) DrawLine(aX, aY, bX, bY float64) {
	imd := imdraw.New(nil)
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(aX, aY))
	imd.Push(pixel.V(bX, bY))
	imd.Line(1)
	imd.Draw(r.window)
}

func (r *Renderer) DrawRect(x, y, w, h, a float64) {
	tlx, tly := x-w/2, y+h/2 // top left
	trx, try := x+w/2, y+h/2 // top right
	blx, bly := x-w/2, y-h/2 // bottom left
	brx, bry := x+w/2, y-h/2 // bottom right

	// if rect has angle, translate corner points
	if a != 0 {
		tlx, tly = rotate(x, y, tlx, tly, a)
		trx, try = rotate(x, y, trx, try, a)
		blx, bly = rotate(x, y, blx, bly, a)
		brx, bry = rotate(x, y, brx, bry, a)
	}

	imd := imdraw.New(nil)
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(tlx, tly))
	imd.Push(pixel.V(trx, try))
	imd.Push(pixel.V(brx, bry))
	imd.Push(pixel.V(blx, bly))
	imd.Push(pixel.V(tlx, tly))
	imd.Line(1)
	imd.Draw(r.window)
}

// rotate rotates a corner around a center point theta radians
func rotate(centerX, centerY, cornerX, cornerY, theta float64) (float64, float64) {
	tempX, tempY := cornerX-centerX, cornerY-centerY
	rotatedX := tempX*math.Cos(theta) - tempY*math.Sin(theta)
	rotatedY := tempX*math.Sin(theta) + tempY*math.Cos(theta)
	return rotatedX + centerX, rotatedY + centerY
}

// remove later, used to experiment with pixelgl easily
func (r *Renderer) Test() {
	//
}
