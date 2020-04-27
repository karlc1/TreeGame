package render

import (
	"karlc/treegame/internal/models"

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

func (r *Renderer) DrawRect(x, y, w, h, a float64) {

	topLeftX := x - w/2
	topLeftY := y + h/2

	topRightX := x + w/2
	topRightY := y + h/2

	bottomLeftX := x - w/2
	bottomLeftY := y - h/2

	bottomRightX := x + w/2
	bottomRightY := y - h/2

	imd := imdraw.New(nil)
	imd.EndShape = imdraw.RoundEndShape
	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(pixel.V(topLeftX, topLeftY))
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(topRightX, topRightY))
	imd.Color = pixel.RGB(0, 0, 1)
	imd.Push(pixel.V(bottomRightX, bottomRightY))
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(bottomLeftX, bottomLeftY))
	imd.Color = pixel.RGB(0, 1, 1)
	imd.Push(pixel.V(topLeftX, topLeftY))
	imd.Polygon(0)
	imd.Draw(r.window)

	//rect := rl.Rectangle{
	//Width:  w,
	//Height: h,
	//// invert x and y axis since raylib and box2d
	//// uses mirrored coordinate systems
	//X: float32(r.screenWidth) - x,
	//Y: float32(r.screenHeight) - y,
	//}

	//origin := rl.Vector2{
	//X: w / 2,
	//Y: h / 2,
	//}

	//// box2d uses radians, raylib uses degrees
	//// this converts the angle to degrees
	//angle := float32(a * (180 / math.Pi))

	//colors := []rl.Color{
	//rl.White,
	////rl.Green,
	//}

	//rl.DrawRectanglePro(rect, origin, angle, colors)
}

func (r *Renderer) DrawLine(aX, aY, bX, bY, thickness float32) {
	//v1 := rl.NewVector2(float32(r.screenWidth)-aX, float32(r.screenHeight)-aY)
	//v2 := rl.NewVector2(float32(r.screenWidth)-bX, float32(r.screenHeight)-bY)
	//rl.DrawLineEx(v1, v2, 1.0, rl.White)
}

func (r *Renderer) DrawBox(box *models.Box) {
	//posX, posY := box.GetPosition()

	//rect := rl.Rectangle{
	//Width:  float32(box.Width),
	//Height: float32(box.Height),
	//// invert x and y axis since raylib and box2d
	//// uses mirrored coordinate systems
	//X: float32(r.screenWidth) - posX,
	//Y: float32(r.screenHeight) - posY,
	//}

	//origin := rl.Vector2{
	//X: float32(box.Width / 2),
	//Y: float32(box.Height / 2),
	//}

	//// box2d uses radians, raylib uses degrees
	//// this converts the angle to degrees
	//angle := float32(box.Body.GetAngle() * (180 / math.Pi))

	//colors := []rl.Color{
	//rl.White,
	//}

	//rl.DrawRectanglePro(rect, origin, angle, colors)

}

func (r *Renderer) Test() {
	imd := imdraw.New(nil)
	imd.EndShape = imdraw.RoundEndShape
	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(pixel.V(100, 100))
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(100, 200))
	imd.Color = pixel.RGB(0, 0, 1)
	imd.Push(pixel.V(200, 200))
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(200, 100))

	imd.Color = pixel.RGB(0, 1, 1)
	imd.Push(pixel.V(100, 100))
	imd.Polygon(0)
	imd.Draw(r.window)
}
