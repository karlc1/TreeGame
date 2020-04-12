package main

import (
	"fmt"
	"karlc/treegame/internal/models"
	"karlc/treegame/internal/render"

	"github.com/ByteArena/box2d"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WITH   = 1200
	SCREEN_HEIGHT = 600
	TARGET_FPS    = 60
)

func main() {

	rl.InitWindow(SCREEN_WITH, SCREEN_HEIGHT, "")
	rl.SetTargetFPS(TARGET_FPS)

	world := models.NewPhysicalWorld()
	renderer := render.NewRenderer(SCREEN_WITH, SCREEN_HEIGHT)

	timeStep := 1.0 / TARGET_FPS
	velocityIterations := 2
	positionIterations := 2

	player := world.NewBox(true, 0, 0, 2, 2)
	player.SetDensity(10)
	player.SetFriction(4)
	world.Player = player

	ground := world.NewBox(false, 0, -50, 10, 1)

	test := world.NewBox(true, 0, 30, 4, 1)
	test.SetDensity(10)
	test.SetFriction(4)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyUp) {
			impulse := world.Player.Body.GetMass() * 50
			world.Player.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(0, impulse), world.Player.Body.GetWorldCenter(), true)
		}

		if rl.IsKeyDown(rl.KeyDown) {
			impulse := world.Player.Body.GetMass() * 50
			world.Player.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(0, -impulse), world.Player.Body.GetWorldCenter(), true)
		}

		if rl.IsKeyDown(rl.KeyRight) {
			impulse := world.Player.Body.GetMass() * 50
			world.Player.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(-impulse, 0), world.Player.Body.GetWorldCenter(), true)
		}

		if rl.IsKeyDown(rl.KeyLeft) {
			impulse := world.Player.Body.GetMass() * 50
			world.Player.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(impulse, 0), world.Player.Body.GetWorldCenter(), true)
		}

		world.PhysWorld.Step(
			timeStep,
			velocityIterations,
			positionIterations,
		)

		//renderer.DrawWorld(world)

		////////////////////////////////////
		camera := render.NewCamera(
			SCREEN_WITH,
			SCREEN_HEIGHT,
			*renderer,
			10,
		)

		// remove
		ppX, ppY := player.GetPosition()
		fmt.Printf("%v   %v \n", ppX, ppY)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		camera.DrawActor(player, false)
		camera.DrawActor(ground, false)
		//camera.DrawActor(test, false)
		rl.EndDrawing()
	}
}
