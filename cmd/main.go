package main

import (
	"karlc/treegame/internal/physics"
	"karlc/treegame/internal/render"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "")
	rl.SetTargetFPS(60)

	world := physics.NewWorld()

	timeStep := 1.0 / 60.0
	velocityIterations := 5
	positionIterations := 5

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeySpace) {
			world.Player.Body.ApplyAngularImpulse(1000.0, true)
		}

		world.PhysWorld.Step(
			timeStep,
			velocityIterations,
			positionIterations,
		)

		render.DrawWorld(world)
	}
}
