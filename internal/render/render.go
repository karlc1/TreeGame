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

	rect := rl.Rectangle{
		Width:  float32(box.Width),
		Height: float32(box.Height),
		// invert x and y axis since raylib and box2d
		// uses mirrored coordinate systems
		X: float32(float64(r.screenWidth) - pos.X),
		Y: float32(float64(r.screenHeight) - pos.Y),
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
