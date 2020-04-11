package render

import (
	"karlc/treegame/internal/physics"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func (r *Renderer) DrawBox(box *physics.Box) {
	pos := box.Body.GetPosition()
	width := box.Width
	height := box.Height

	angle64 := box.Body.GetAngle()

	rect := rl.Rectangle{
		Width:  float32(width),
		Height: float32(height),
		X:      float32(pos.X),
		// Invert Y axis since box2d and
		// raylib has vertically opposed Y axis
		Y: float32(float64(r.screenHeight) - pos.Y),
	}

	origin := rl.Vector2{
		X: float32(width / 2),
		Y: float32(height / 2),
	}

	// box2d uses radians, raylib uses degrees
	// this converts the angle to degrees
	angle := float32(angle64 * (180 / math.Pi))

	colors := []rl.Color{
		rl.White,
	}

	//fmt.Printf("Width: %v \n Height %v \n X: %v \n Y: %v \n A: %v \n\n",
	//width, height, pos.X, pos.Y, angle)

	rl.DrawRectanglePro(rect, origin, angle, colors)

	//rl.DrawRectangle(int32(pos.X), int32(pos.Y), int32(width), int32(height), rl.White)
}

func (r *Renderer) DrawWorld(world *physics.World) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for _, box := range world.Boxes {
		r.DrawBox(box)
	}

	rl.EndDrawing()
}
