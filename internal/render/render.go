package render

import (
	"karlc/treegame/internal/models"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func InitWindow(w, h, fps int32) {
	rl.InitWindow(w, h, "")
	rl.SetTargetFPS(fps)
}

type Renderer struct {
	screenWidth  int
	screenHeight int
}

func NewRenderer(w, h int) *Renderer {
	return &Renderer{
		screenHeight: h,
		screenWidth:  w,
	}
}

func (r *Renderer) DrawRect(x, y, w, h, a float32) {
	rect := rl.Rectangle{
		Width:  w,
		Height: h,
		// invert x and y axis since raylib and box2d
		// uses mirrored coordinate systems
		X: float32(r.screenWidth) - x,
		Y: float32(r.screenHeight) - y,
	}

	origin := rl.Vector2{
		X: w / 2,
		Y: h / 2,
	}

	// box2d uses radians, raylib uses degrees
	// this converts the angle to degrees
	angle := float32(a * (180 / math.Pi))

	colors := []rl.Color{
		rl.White,
		//rl.Green,
	}

	rl.DrawRectanglePro(rect, origin, angle, colors)
}

func (r *Renderer) DrawBox(box *models.Box) {
	posX, posY := box.GetPosition()

	rect := rl.Rectangle{
		Width:  float32(box.Width),
		Height: float32(box.Height),
		// invert x and y axis since raylib and box2d
		// uses mirrored coordinate systems
		X: float32(r.screenWidth) - posX,
		Y: float32(r.screenHeight) - posY,
	}

	origin := rl.Vector2{
		X: float32(box.Width / 2),
		Y: float32(box.Height / 2),
	}

	// box2d uses radians, raylib uses degrees
	// this converts the angle to degrees
	angle := float32(box.Body.GetAngle() * (180 / math.Pi))

	colors := []rl.Color{
		rl.White,
	}

	rl.DrawRectanglePro(rect, origin, angle, colors)

}
