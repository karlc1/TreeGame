package main

import (
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

		renderer.DrawWorld(world)
	}
}
