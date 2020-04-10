package render

import (
	"fmt"
	"karlc/treegame/internal/physics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var tempPos int32 = 0

func DrawBox(box *physics.Box) {
	pos := box.Body.GetPosition()
	angle64 := box.Body.GetAngle()
	width := box.Width
	height := box.Height

	rect := rl.Rectangle{
		Width:  float32(width) * 10,
		Height: float32(height) * 10,
		X:      float32(pos.X),
		Y:      float32(pos.Y),
	}

	origin := rl.Vector2{
		X: float32(pos.X),
		Y: float32(pos.Y),
	}

	angle := float32(angle64)

	colors := []rl.Color{
		rl.White,
	}

	fmt.Printf("Width: %v \n Height %v \n X: %v \n Y: %v \n A: %v \n\n",
		width, height, pos.X, pos.Y, angle)

	rl.DrawRectanglePro(rect, origin, angle, colors)
}

func DrawWorld(world *physics.World) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for _, box := range world.Boxes {
		DrawBox(box)
	}

	rl.EndDrawing()
}
