package render

import (
	"fmt"
	"karlc/treegame/internal/physics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var tempPos int32 = 0

func DrawBox(box *physics.Box) {
	pos := box.Body.GetPosition()
	width := box.Width
	height := box.Height

	angle64 := box.Body.GetAngle()
	rect := rl.Rectangle{
		Width:  float32(width),
		Height: float32(height),
		X:      float32(pos.X),
		Y:      float32(pos.Y),
	}

	origin := rl.Vector2{
		X: float32(width / 2),
		Y: float32(height / 2),
	}

	angle := float32(angle64)

	colors := []rl.Color{
		rl.White,
	}

	fmt.Printf("Width: %v \n Height %v \n X: %v \n Y: %v \n A: %v \n\n",
		width, height, pos.X, pos.Y, angle)

	rl.DrawRectanglePro(rect, origin, angle, colors)

	//rl.DrawRectangle(int32(pos.X), int32(pos.Y), int32(width), int32(height), rl.White)
}

func DrawWorld(world *physics.World) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for _, box := range world.Boxes {
		DrawBox(box)
	}

	rl.EndDrawing()
}
