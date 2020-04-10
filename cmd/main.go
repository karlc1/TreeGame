package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "")
	rl.SetTargetFPS(60)

	//world := physics.NewWorld()

	//timeStep := 1.0 / 10.0
	//velocityIterations := 6
	//positionIterations := 2

	for !rl.WindowShouldClose() {
		//world.PhysWorld.Step(
		//timeStep,
		//velocityIterations,
		//positionIterations,
		//)

		//render.DrawWorld(world)
	}
}
