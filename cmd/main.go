package main

import (
	"karlc/treegame/internal/models"
	"karlc/treegame/internal/render"
	"karlc/treegame/internal/utils"

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

	player := world.NewBox(true, -29, 10, 1, 1.5)
	player.SetDensity(10)
	player.SetFriction(4)
	world.Player = player

	ground := world.NewBox(false, 0, -25, 100, 10)
	_ = ground

	decor := make([]*models.DecorBox, 1000, 1000)
	for i := range decor {
		size := utils.RandFloat32(0, 0.5)
		posX := utils.RandFloat32(-200, 200)
		posY := utils.RandFloat32(-50, 50)

		// for parallax test
		z := int(size * 10)

		decor[i] = &models.DecorBox{
			Height: size,
			Width:  size,
			PosX:   posX,
			PosY:   posY,
			Zval:   z,
		}
	}

	//test := world.NewBox(true, 0, 30, 4, 1)
	//test.SetDensity(10)
	//test.SetFriction(4)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyUp) {
			impulse := world.Player.Body.GetMass() * 2
			world.Player.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(0, impulse), world.Player.Body.GetWorldCenter(), true)
		}

		if rl.IsKeyDown(rl.KeyDown) {
			impulse := world.Player.Body.GetMass() * 2
			world.Player.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(0, -impulse), world.Player.Body.GetWorldCenter(), true)
		}

		if rl.IsKeyDown(rl.KeyRight) {
			impulse := world.Player.Body.GetMass() * 2
			world.Player.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(-impulse, 0), world.Player.Body.GetWorldCenter(), true)
		}

		if rl.IsKeyDown(rl.KeyLeft) {
			impulse := world.Player.Body.GetMass() * 2
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

		camera.AttachTo(player)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		camera.DrawActor(player, false)
		camera.DrawActor(ground, false)
		//camera.DrawActor(test, false)

		for _, e := range decor {
			camera.DrawActor(e, false)
		}

		rl.EndDrawing()
	}
}
